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

func (s *Server) RetrieveAttributesData(ctx context.Context, in *pbsemantic.RetrieveAttributesDataRequest) (*pbsemantic.RetrieveAttributesDataResponse, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s.retrieveAttributesData(ctx, db, in)
}

func (s *Server) retrieveAttributesData(ctx context.Context, db *sqlx.DB, in *pbsemantic.RetrieveAttributesDataRequest) (*pbsemantic.RetrieveAttributesDataResponse, error) {
	atts, err := data.FindProductAttributesDataByMatch(ctx, db, in.GetEntityTable(), in.GetEntityId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindProductAttributesDataByMatch: %v", err)
	}
	resp := &pbsemantic.RetrieveAttributesDataResponse{}
	for _, att := range atts {
		pbAtt := &pbsemantic.SetAttributesDataResponse{Data: server.DataAttributeToPbAttributeData(att), AttributeDataId: att.Id}
		resp.Attributes = append(resp.Attributes, pbAtt)
	}

	return resp, nil
}
