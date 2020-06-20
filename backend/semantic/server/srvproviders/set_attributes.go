package srvproviders

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"time"

	"github.com/athomecomar/athome/backend/semantic/data"
	"github.com/athomecomar/athome/backend/semantic/data/value"
	"github.com/athomecomar/athome/backend/semantic/schema"
	"github.com/athomecomar/athome/backend/semantic/server"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbshared"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) SetAttributeDatas(srv pbsemantic.ServiceProviders_SetAttributeDatasServer) error {
	ctx := srv.Context()
	context.WithTimeout(ctx, 500*time.Second)

	db, err := server.ConnDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var userId uint64
	var entity *pbshared.Entity
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default: // no-op
		}

		in, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = in.Validate()
		if err != nil {
			return err
		}

		if entity == nil {
			auth := in.GetAuthorization()
			userId, err = server.AuthorizeThroughEntity(ctx, auth.GetAccessToken(), auth.GetEntity())
			if err != nil {
				return err
			}
			entity = auth.GetEntity()
			continue
		}

		resp, err := s.setAttributeDatas(ctx, db, in.GetData(), userId, entity)
		if err != nil {
			return err
		}

		err = srv.Send(resp)
		if err != nil {
			return err
		}
	}
}

func (s *Server) setAttributeDatas(
	ctx context.Context,
	db *sqlx.DB,
	in *pbsemantic.AttributeData,
	userId uint64,
	entity *pbshared.Entity,
) (*pbsemantic.SetAttributeDatasResponse, error) {
	attSchema, err := schema.FindServiceProviderAttributeSchema(ctx, db, in.GetSchemaId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindServiceProviderAttributeSchemas: %v", err)
	}

	var d data.Attribute
	d, err = data.FindServiceProviderAttributeDataByMatch(ctx, db, in.GetSchemaId(), entity) // yes, it can store multi attrs in one match on wrapper, but thats safer
	if errors.Is(err, sql.ErrNoRows) {
		d, err = attSchema.NewData()
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "schema.NewData: %v", err)
		}
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindServiceProviderAttributeDataByMatch: %v", err)
	}

	err = data.MustUserId(d, userId)
	if err != nil {
		return nil, status.Errorf(xerrors.PermissionDenied, "MustUserId: %v", err)
	}

	val, err := value.Parse(attSchema.GetValueType(), in.GetValues()...)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "value.Parse: %v", err)
	}
	err = d.SetValue(val)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "data.SetValue: %v", err)
	}
	err = storeql.InsertIntoDB(ctx, db, d)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}

	return &pbsemantic.SetAttributeDatasResponse{
		AttributeDataId: d.GetId(),
		Data:            data.AttributeToPb(d),
	}, nil
}
