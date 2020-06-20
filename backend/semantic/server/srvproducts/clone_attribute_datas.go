package srvproducts

import (
	"context"

	"github.com/athomecomar/athome/backend/semantic/data"
	"github.com/athomecomar/athome/backend/semantic/server"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbshared"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) CloneAttributeDatas(ctx context.Context, in *pbsemantic.CloneAttributeDatasRequest) (*pbsemantic.CloneAttributeDatasResponse, error) {
	err := in.Validate()
	if err != nil {
		return nil, err
	}
	_, err = server.AuthorizeThroughEntity(ctx, in.GetAccessToken(), &pbshared.Entity{EntityId: in.GetDestEntityId(), EntityTable: in.GetEntityTable()})
	if err != nil {
		return nil, err
	}

	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	return s.cloneAttributeDatas(ctx, db, in)
}

func (s *Server) cloneAttributeDatas(ctx context.Context, db *sqlx.DB, in *pbsemantic.CloneAttributeDatasRequest) (*pbsemantic.CloneAttributeDatasResponse, error) {
	atts, err := data.FindProductAttributeDatasByMatch(ctx, db, &pbshared.Entity{EntityTable: in.GetEntityTable(), EntityId: in.GetFromEntityId()})
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "FindProductAttributeDatasByMatch: %v", err)
	}

	var respData []*pbsemantic.AttributeData
	var clones []storeql.Storable
	for _, att := range atts {
		clone, err := att.Clone()
		if err != nil {
			return nil, status.Errorf(xerrors.Internal, "Clone: %v", err)
		}
		clones = append(clones, clone)
		respData = append(respData, data.AttributeToPb(clone))
	}
	err = storeql.InsertIntoDB(ctx, db, clones...)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "storeql.InsertIntoDB: %v", err)
	}
	resp := &pbsemantic.CloneAttributeDatasResponse{}
	resp.Attributes = make(map[uint64]*pbsemantic.AttributeData)
	for i, data := range respData {
		resp.Attributes[clones[i].GetId()] = data
	}
	return resp, nil
}
