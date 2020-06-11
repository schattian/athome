package srvproducts

import (
	"context"

	"github.com/athomecomar/athome/backend/semantic/data"
	"github.com/athomecomar/athome/backend/semantic/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/semantic/server"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveAttributeDatas(ctx context.Context, in *pbsemantic.RetrieveAttributeDatasRequest) (*pbsemantic.RetrieveAttributeDatasResponse, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s.retrieveAttributeDatas(ctx, db, in)
}

func (s *Server) retrieveAttributeDatas(ctx context.Context, db *sqlx.DB, in *pbsemantic.RetrieveAttributeDatasRequest) (*pbsemantic.RetrieveAttributeDatasResponse, error) {
	atts, err := data.FindProductAttributeDatasByMatch(ctx, db, in.GetEntityTable(), in.GetEntityId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindProductAttributeDatasByMatch: %v", err)
	}
	resp := &pbsemantic.RetrieveAttributeDatasResponse{}
	resp.Attributes = make(map[uint64]*pbsemantic.AttributeData)
	for _, att := range atts {
		resp.Attributes[att.Id] = data.AttributeToPb(att)
	}

	return resp, nil
}
