package ent

import (
	"context"

	"github.com/athomecomar/athome/pb/pbproducts"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/currency"
	"github.com/athomecomar/storeql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type DraftLine struct {
	Id      uint64 `json:"id,omitempty"`
	DraftId uint64 `json:"draft_id,omitempty"`

	// First
	Title      string `json:"title,omitempty"`
	CategoryId uint64 `json:"category_id,omitempty"`

	// Second
	Price currency.ARS `json:"price,omitempty"`
	Stock uint64       `json:"stock,omitempty"`

	// Third
	ImageIds []string `json:"image_ids,omitempty"`
}

func (ln *DraftLine) toProduct() *Product {
	return &Product{
		CategoryId: ln.CategoryId,
		Title:      ln.Title,
		Price:      ln.Price,
		Stock:      ln.Stock,
		ImageIds:   ln.ImageIds,
	}
}

func (ln *DraftLine) finish(ctx context.Context, db *sqlx.DB, sem pbsemantic.ProductsClient, userId uint64, access string) (*Product, error) {
	prod := ln.toProduct()
	prod.UserId = userId

	err := storeql.InsertIntoDB(ctx, db, prod)
	if err != nil {
		return nil, errors.Wrap(err, "InsertIntoDB")
	}

	_, err = sem.ChangeEntityAttributeDatas(ctx, &pbsemantic.ChangeEntityAttributeDatasRequest{
		AccessToken:     access,
		FromEntityTable: ln.SQLTable(),
		FromEntityId:    ln.Id,
		DestEntityTable: prod.SQLTable(),
		DestEntityId:    prod.Id,
	})
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func (ln *DraftLine) ToPb(atts []*pbproducts.AttributeData) *pbproducts.DraftLine {
	return &pbproducts.DraftLine{
		DraftLineId: ln.Id,

		First: &pbproducts.DraftLineFirst{
			Title:      ln.Title,
			CategoryId: ln.CategoryId,
		},

		Second: &pbproducts.DraftLineSecond{
			Price:      ln.Price.Float64(),
			Stock:      ln.Stock,
			Attributes: atts,
		},

		Third: &pbproducts.DraftLineThird{
			ImageIds: ln.ImageIds,
		},
	}
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
