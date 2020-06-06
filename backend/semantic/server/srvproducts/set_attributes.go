package srvproducts

import (
	"context"
	"io"
	"time"

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

	var authorized bool
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

		if !authorized {
			err = server.AuthorizeThroughEntity(ctx, in.GetAccessToken(), in.GetEntityId(), in.GetEntityTable())
			if err != nil {
				return err
			}
			authorized = true
		}

		resp, err := s.setAttributesData(ctx, db, in.GetData())
		if err != nil {
			return err
		}

		err = srv.Send(resp)
		if err != nil {
			return err
		}
	}
}

func (s *Server) setAttributesData(ctx context.Context, db *sqlx.DB, in *pbsemantic.AttributeData) (*pbsemantic.SetAttributesDataResponse, error) {
	attSchema, err := schema.FindProductAttributeSchema(ctx, db, in.GetSchemaId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindProductAttributesSchema: %v", err)
	}
	val, err := value.Parse(attSchema.GetValueType(), in.GetValues()...)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "value.Parse: %v", err)
	}
	data, err := attSchema.NewData()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "schema.NewData: %v", err)
	}
	err = data.SetValue(val)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "data.SetValue: %v", err)
	}
	err = storeql.InsertIntoDB(ctx, db, data)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}

	return &pbsemantic.SetAttributesDataResponse{
		AttributeDataId: data.GetId(),
		Data:            server.DataAttributeToPbAttributeData(data),
	}, nil
}
