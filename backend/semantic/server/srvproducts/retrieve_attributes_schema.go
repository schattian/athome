package srvproducts

import (
	"context"

	"github.com/athomecomar/athome/backend/semantic/schema"
	"github.com/athomecomar/athome/backend/semantic/server"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveAttributeSchemas(ctx context.Context, in *pbsemantic.RetrieveAttributeSchemasRequest) (*pbsemantic.RetrieveAttributeSchemasResponse, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s.getAttributeSchemas(ctx, db, in)
}

func (s *Server) getAttributeSchemas(ctx context.Context, db *sqlx.DB, in *pbsemantic.RetrieveAttributeSchemasRequest) (*pbsemantic.RetrieveAttributeSchemasResponse, error) {
	rows, err := db.QueryxContext(ctx, `SELECT * FROM product_attribute_schemas WHERE category_id = $1`, in.GetCategoryId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "QueryxContext: %v", err)
	}
	defer rows.Close()
	atts := &pbsemantic.RetrieveAttributeSchemasResponse{}
	atts.Attributes = make(map[uint64]*pbsemantic.AttributeSchema)
	for rows.Next() {
		att := &schema.ProductAttributeSchema{}
		err = rows.StructScan(att)
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "StructScan: %v", err)
		}
		atts.Attributes[att.Id] = schema.AttributeToPb(att)
	}

	return atts, nil
}
