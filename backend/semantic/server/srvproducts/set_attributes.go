package srvproducts

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"time"

	"github.com/athomecomar/athome/backend/semantic/data"
	"github.com/athomecomar/athome/backend/semantic/data/value"
	"github.com/athomecomar/athome/backend/semantic/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/semantic/schema"
	"github.com/athomecomar/athome/backend/semantic/server"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) SetAttributesData(srv pbsemantic.Products_SetAttributesDataServer) error {
	ctx := srv.Context()
	context.WithTimeout(ctx, 500*time.Second)

	db, err := server.ConnDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var userId, entityId uint64
	var entityTable string
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

		if entityTable == "" || entityId == 0 {
			auth := in.GetAuthorization()
			entityTable, entityId = auth.GetEntityTable(), auth.GetEntityId()
			userId, err = server.AuthorizeThroughEntity(ctx, auth.GetAccessToken(), entityId, entityTable)
			if err != nil {
				return err
			}
			continue
		}

		resp, err := s.setAttributesData(ctx, db, in.GetData(), userId, entityTable, entityId)
		if err != nil {
			return err
		}

		err = srv.Send(resp)
		if err != nil {
			return err
		}
	}
}

func (s *Server) setAttributesData(
	ctx context.Context,
	db *sqlx.DB,
	in *pbsemantic.AttributeData,
	userId uint64,
	entityTable string,
	entityId uint64,
) (*pbsemantic.SetAttributesDataResponse, error) {
	attSchema, err := schema.FindProductAttributeSchema(ctx, db, in.GetSchemaId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindProductAttributesSchema: %v", err)
	}

	var d data.Attribute
	d, err = data.FindProductAttributeDataByMatch(ctx, db, in.GetSchemaId(), entityTable, entityId) // yes, it can store multi attrs in one match on wrapper, but thats safer
	if errors.Is(err, sql.ErrNoRows) {
		d, err = attSchema.NewData()
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "schema.NewData: %v", err)
		}
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindProductAttributeDataByMatch: %v", err)
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

	return &pbsemantic.SetAttributesDataResponse{
		AttributeDataId: d.GetId(),
		Data:            server.DataAttributeToPbAttributeData(d),
	}, nil
}