package ent

import (
	"context"

	"github.com/athomecomar/currency"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type DraftLine struct {
	Id      uint64
	DraftId uint64

	// First
	Title      string
	CategoryId uint64

	// Second
	Price currency.ARS
	Stock uint64

	// Third
}

func (ln *DraftLine) Draft(ctx context.Context, db *sqlx.DB) (*Draft, error) {
	row := db.QueryRowxContext(ctx, `SELECT * FROM drafts WHERE id=$1`, ln.DraftId)
	d := &Draft{}
	err := row.StructScan(ln)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return d, nil
}

func (ln *DraftLine) Clone() (*DraftLine, error) {
	if ln == nil {
		return nil, errors.New("nil is not clonable")
	}
	cp := DraftLine{}
	cp = *ln
	cp.Stock = 0
	cp.Id = 0
	return &cp, nil
}

func LineById(ctx context.Context, db *sqlx.DB, id uint64) (*DraftLine, error) {
	rows := db.QueryRowxContext(ctx, `SELECT * FROM draft_lines WHERE id=$1`, id)
	ln := &DraftLine{}
	err := rows.StructScan(ln)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return ln, nil
}
