package srvproducts

import (
	"context"

	"github.com/athomecomar/athome/backend/semantic/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/semantic/schema"
	"github.com/athomecomar/athome/backend/semantic/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAttributesSchema(ctx context.Context, in *pbsemantic.GetAttributesSchemaRequest) (*pbsemantic.GetAttributesSchemaResponse, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s.getAttributesSchema(ctx, db, in)
}

func (s *Server) getAttributesSchema(ctx context.Context, db *sqlx.DB, in *pbsemantic.GetAttributesSchemaRequest) (*pbsemantic.GetAttributesSchemaResponse, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM product_attributes_schema WHERE category_id = $1`, in.GetCategoryId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "QueryxContext: %v", err)
	}

	atts := &pbsemantic.GetAttributesSchemaResponse{}
	for rows.Next() {
		att := &schema.ProductAttributeSchema{}
		err = rows.StructScan(att)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "StructScan: %v", err)
		}
		atts.Attributes = append(atts.Attributes, server.AttributeToPbsemanticAttribute(att))
	}

	return atts, nil
}
