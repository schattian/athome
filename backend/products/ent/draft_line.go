package ent

import (
	"context"

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
