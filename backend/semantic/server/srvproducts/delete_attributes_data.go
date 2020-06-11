package srvproducts

import (
	"context"

	"github.com/athomecomar/athome/backend/semantic/data"
	"github.com/athomecomar/athome/backend/semantic/server"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteAttributeDatas(ctx context.Context, in *pbsemantic.DeleteAttributeDatasRequest) (*emptypb.Empty, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	_, err = server.AuthorizeThroughEntity(ctx, in.GetAccessToken(), in.GetEntityId(), in.GetEntityTable())
	if err != nil {
		return nil, err
	}

	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return s.deleteAttributeDatas(ctx, db, in)
}

func (s *Server) deleteAttributeDatas(ctx context.Context, db *sqlx.DB, in *pbsemantic.DeleteAttributeDatasRequest) (*emptypb.Empty, error) {
	atts, err := data.FindProductAttributeDatasByMatch(ctx, db, in.GetEntityTable(), in.GetEntityId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindProductAttributeDatasByMatch: %v", err)
	}
	var sts []storeql.Storable
	for _, att := range atts {
		sts = append(sts, att)
	}
	err = storeql.DeleteFromDB(ctx, db, sts...)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.DeleteFromDB: %v", err)
	}
	return &emptypb.Empty{}, nil
}
