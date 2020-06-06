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

func (s *Server) GetAttributesData(ctx context.Context, in *pbsemantic.GetAttributesDataRequest) (*pbsemantic.GetAttributesDataResponse, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return s.getAttributesData(ctx, db, in)
}

func (s *Server) getAttributesData(ctx context.Context, db *sqlx.DB, in *pbsemantic.GetAttributesDataRequest) (*pbsemantic.GetAttributesDataResponse, error) {
	atts, err := data.FindProductAttributesDataByMatch(ctx, db, in.GetEntityTable(), in.GetEntityId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindProductAttributesDataByMatch: %v", err)
	}
	resp := &pbsemantic.GetAttributesDataResponse{}
	for _, att := range atts {
		pbAtt := &pbsemantic.SetAttributesDataResponse{Data: server.DataAttributeToPbAttributeData(att), AttributeDataId: att.Id}
		resp.Attributes = append(resp.Attributes, pbAtt)
	}

	return resp, nil
}
