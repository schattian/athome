package srvproducts

import (
	"context"

	"github.com/athomecomar/athome/backend/semantic/data"
	"github.com/athomecomar/athome/backend/semantic/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/semantic/server"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteAttributesData(ctx context.Context, in *pbsemantic.DeleteAttributesDataRequest) (*emptypb.Empty, error) {
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

	return s.deleteAttributesData(ctx, db, in)
}

func (s *Server) deleteAttributesData(ctx context.Context, db *sqlx.DB, in *pbsemantic.DeleteAttributesDataRequest) (*emptypb.Empty, error) {
	atts, err := data.FindProductAttributesDataByMatch(ctx, db, in.GetEntityTable(), in.GetEntityId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindProductAttributesDataByMatch: %v", err)
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
