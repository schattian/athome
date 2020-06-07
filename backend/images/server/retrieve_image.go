package server

import (
	"context"

	"github.com/athomecomar/athome/backend/images/pb/pbimages"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveImage(ctx context.Context, in *pbimages.RetrieveImageRequest) (*pbimages.RetrieveImageResponse, error) {
	return s.retrieveImage(ctx, in)
}

func (s *Server) retrieveImage(ctx context.Context, in *pbimages.RetrieveImageRequest) (*pbimages.RetrieveImageResponse, error) {
	dd, err := s.Store.Retrieve(ctx, in.GetId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "store.Find: %v", err)
	}
	meta, err := dd.Metadata()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "data.Metadata: %v", err)
	}
	return &pbimages.RetrieveImageResponse{Uri: dd.URI(), UserId: meta.UserId}, nil
}
