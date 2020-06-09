package srvregister

import (
	"context"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/ent/stage"
	"github.com/athomecomar/athome/backend/services/pb/pbservices"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/currency"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Second(ctx context.Context, in *pbservices.SecondRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	auth, authCloser, err := server.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()

	reg, err := retrieveRegistryByUser(ctx, db, auth, in.GetAccessToken(), server.GetUserFromAccessToken)
	if err != nil {
		return nil, err
	}
	authCloser()

	return s.second(ctx, db, in, reg)
}

func (s *Server) second(ctx context.Context, db *sqlx.DB, in *pbservices.SecondRequest, reg *ent.Registry) (*emptypb.Empty, error) {
	err := mustStage(reg.Stage, stage.Second)
	if err != nil {
		return nil, err
	}
	reg = applySecondRequestToRegistry(in.GetBody(), reg)
	reg.Stage = reg.Stage.Next()
	err = storeql.UpdateIntoDB(ctx, db, reg)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "UpdateIntoDB: %v", err)
	}
	return &emptypb.Empty{}, nil
}

func applySecondRequestToRegistry(f *pbservices.SecondRequest_Body, r *ent.Registry) *ent.Registry {
	r.DurationInMinutes = f.GetDurationInMinutes()
	r.Name = f.GetName()
	r.PriceMax = currency.ToARS(f.GetPrice().GetMax())
	r.PriceMin = currency.ToARS(f.GetPrice().GetMin())
	return r
}
