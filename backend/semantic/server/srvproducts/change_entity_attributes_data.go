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

func (s *Server) ChangeEntityAttributesData(ctx context.Context, in *pbsemantic.ChangeEntityAttributesDataRequest) (*emptypb.Empty, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}

	_, err = server.AuthorizeThroughEntity(ctx, in.GetAccessToken(), in.GetDestEntityId(), in.GetDestEntityTable())
	if err != nil {
		return nil, err
	}

	_, err = server.AuthorizeThroughEntity(ctx, in.GetAccessToken(), in.GetFromEntityId(), in.GetFromEntityTable())
	if err != nil {
		return nil, err
	}

	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return s.changeEntityAttributesData(ctx, db, in)
}

func (s *Server) changeEntityAttributesData(ctx context.Context, db *sqlx.DB, in *pbsemantic.ChangeEntityAttributesDataRequest) (*emptypb.Empty, error) {
	atts, err := data.FindProductAttributesDataByMatch(ctx, db, in.GetFromEntityTable(), in.GetFromEntityId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindProductAttributesDataByMatch: %v", err)
	}

	var storables []storeql.Storable
	for _, att := range atts {
		att.EntityId, att.EntityTable = in.GetDestEntityId(), in.GetDestEntityTable()
		storables = append(storables, att)
	}

	err = storeql.UpdateIntoDB(ctx, db, storables...)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.UpdateIntoDB: %v", err)
	}
	return &emptypb.Empty{}, nil
}
