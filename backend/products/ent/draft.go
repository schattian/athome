package ent

import (
	"context"
	"database/sql"

	"github.com/athomecomar/athome/backend/products/ent/stage"
	"github.com/athomecomar/storeql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Draft struct {
	Id     uint64
	Stage  stage.Stage
	UserId uint64
	// First

	// Second

	// Third
}

func (d *Draft) ValidateLineByStage(l *DraftLine) error {
	switch d.Stage {
	case stage.First:
	case stage.Second:
	case stage.Fourth:
	case stage.Nil:
	}
	return nil
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
