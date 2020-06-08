package ent

import (
	"context"
	"database/sql"
	"sync"

	"github.com/athomecomar/athome/backend/products/ent/stage"
	"github.com/athomecomar/athome/backend/products/pb/pbsemantic"
	"github.com/athomecomar/storeql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Draft struct {
	Id     uint64      `json:"id,omitempty"`
	Stage  stage.Stage `json:"stage,omitempty"`
	UserId uint64      `json:"user_id,omitempty"`
}

func (d *Draft) ValidateLineByStage(l *DraftLine) error {
	switch d.Stage {
	case stage.First:
		if l.Title == "" {
			return errors.New("line's title is nil")
		}
		if l.CategoryId == 0 {
			return errors.New("line's categoryId is nil")
		}
	case stage.Second:
		if l.Price <= 0 {
			return errors.New("line's price is <= 0")
		}
		if l.Stock == 0 {
			return errors.New("line's stock is nil")
		}
	case stage.Third:
		if l.ImageIds == nil {
			return errors.New("line's stock is nil")
		}
	}
	return nil
}

func (d *Draft) Prev(ctx context.Context, db *sqlx.DB) error {
	d.Stage = d.Stage.Prev()
	err := storeql.UpdateIntoDB(ctx, db, d)
	if err != nil {
		return errors.Wrap(err, "UpdateIntoDB")
	}

	return nil
}

func (d *Draft) finish(ctx context.Context, db *sqlx.DB, sem pbsemantic.ProductsClient, lns []*DraftLine, access string) (prods []*Product, err error) {
	var wg sync.WaitGroup

	prodCh := make(chan *Product)
	errCh := make(chan error, 1)
	done := make(chan struct{})
	for _, ln := range lns {
		ln := ln

		wg.Add(1)
		go func() {
			defer wg.Done()
			prod, err := ln.finish(ctx, db, sem, d.UserId, access)
			if err != nil {
				errCh <- err
			}
			prodCh <- prod
		}()
	}

	go func() {
		wg.Wait()
		close(done)
	}()
	for {
		select {
		case err = <-errCh:
			err = errors.Wrap(err, "select")
			return
		case prod := <-prodCh:
			prods = append(prods, prod)
		case <-done:
			err = storeql.DeleteFromDB(ctx, db, d)
			if err != nil {
				return nil, errors.Wrap(err, "DeleteFromDB")
			}
			return
		}
	}
}

func (d *Draft) Next(ctx context.Context, db *sqlx.DB, sem pbsemantic.ProductsClient, accessToken string) (int, error) {
	d.Stage = d.Stage.Next()

	lns, err := d.Lines(ctx, db)
	if err != nil {
		return 0, errors.Wrap(err, "Lines")
	}

	var upstreamLines []*DraftLine
	for _, ln := range lns {
		if d.ValidateLineByStage(ln) == nil {
			upstreamLines = append(upstreamLines, ln)
		}
	}

	if len(upstreamLines) == 0 {
		return 0, errors.New("no upstream lines")
	}
	qt := len(upstreamLines)

	switch d.Stage {
	case stage.Fourth:
		prods, err := d.finish(ctx, db, sem, upstreamLines, accessToken)
		if err != nil {
			return 0, errors.Wrap(err, "draft.finish")
		}
		qt = len(prods)
	default:
		err = storeql.UpdateIntoDB(ctx, db, d)
		if err != nil {
			return 0, errors.Wrap(err, "UpdateIntoDB")
		}
	}

	return qt, nil
}

func (d *Draft) LineByTitle(ctx context.Context, db *sqlx.DB, title string) (*DraftLine, error) {
	rows := db.QueryRowxContext(ctx, `SELECT * FROM draft_lines WHERE draft_id=$1 AND title=$2`, d.Id, title)
	ln := &DraftLine{}
	err := rows.StructScan(ln)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return ln, nil
}

func (d *Draft) LineById(ctx context.Context, db *sqlx.DB, id uint64) (*DraftLine, error) {
	rows := db.QueryRowxContext(ctx, `SELECT * FROM draft_lines WHERE draft_id=$1 AND id=$2`, d.Id, id)
	ln := &DraftLine{}
	err := rows.StructScan(ln)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return ln, nil
}

func (d *Draft) Lines(ctx context.Context, db *sqlx.DB) (lns []*DraftLine, err error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM draft_lines WHERE draft_id=$1`, d.Id)
	if err != nil {
		return nil, errors.Wrap(err, "QueryxContext")
	}
	for rows.Next() {
		ln := &DraftLine{}
		err = rows.StructScan(ln)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		lns = append(lns, ln)
	}
	return
}

func FindOrCreateDraft(ctx context.Context, db *sqlx.DB, userId uint64) (*Draft, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM drafts WHERE user_id=$1`, userId)

	d := &Draft{UserId: userId, Stage: stage.First}

	err := row.StructScan(d)
	if errors.Is(err, sql.ErrNoRows) {
		err = storeql.InsertIntoDB(ctx, db, d)
		if err != nil {
			return nil, errors.Wrap(err, "storeql.InsertIntoDB")
		}
	}
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}

	return d, nil
}
