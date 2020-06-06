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
)

func (s *Server) CloneAttributesData(ctx context.Context, in *pbsemantic.CloneAttributesDataRequest) (*pbsemantic.CloneAttributesDataResponse, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	_, err = server.AuthorizeThroughEntity(ctx, in.GetAccessToken(), in.GetDestEntityId(), in.GetEntityTable())
	if err != nil {
		return nil, err
	}

	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return s.cloneAttributesData(ctx, db, in)
}

func (s *Server) cloneAttributesData(ctx context.Context, db *sqlx.DB, in *pbsemantic.CloneAttributesDataRequest) (*pbsemantic.CloneAttributesDataResponse, error) {
	atts, err := data.FindProductAttributesDataByMatch(ctx, db, in.GetEntityTable(), in.GetFromEntityId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindProductAttributesDataByMatch: %v", err)
	}

	var respData []*pbsemantic.AttributeData
	var clones []storeql.Storable
	for _, att := range atts {
		clone, err := att.Clone()
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "Clone: %v", err)
		}
		clones = append(clones, clone)
		respData = append(respData, server.DataAttributeToPbAttributeData(clone))
	}
	err = storeql.InsertIntoDB(ctx, db, clones...)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}
	resp := &pbsemantic.CloneAttributesDataResponse{}

	for i, data := range respData {
		att := &pbsemantic.SetAttributesDataResponse{AttributeDataId: clones[i].GetId(), Data: data}
		resp.Attributes = append(resp.Attributes, att)
	}
	return resp, nil
}
