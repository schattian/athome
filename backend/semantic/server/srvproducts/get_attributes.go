package srvproducts

import (
	"context"

	"github.com/athomecomar/athome/backend/semantic/ent"
	"github.com/athomecomar/athome/backend/semantic/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/semantic/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAttributes(ctx context.Context, in *pbsemantic.GetAttributesRequest) (*pbsemantic.GetAttributesResponse, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	return s.getAttributes(ctx, db, in)
}

func (s *Server) getAttributes(ctx context.Context, db *sqlx.DB, in *pbsemantic.GetAttributesRequest) (*pbsemantic.GetAttributesResponse, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM product_attributes WHERE category_id = $1`, in.GetCategoryId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "QueryxContext: %v", err)
	}

	atts := &pbsemantic.GetAttributesResponse{}
	for rows.Next() {
		att := &ent.ProductAttribute{}
		err = rows.StructScan(att)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "StructScan: %v", err)
		}
		atts.Attributes = append(atts.Attributes, server.AttributeToPbsemanticAttribute(att))
	}

	return atts, nil
}
