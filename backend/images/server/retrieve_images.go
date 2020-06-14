package server

import (
	"context"
	"sync"

	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func (s *Server) RetrieveImages(ctx context.Context, in *pbimages.RetrieveImagesRequest) (*pbimages.RetrieveImagesResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	return s.retrieveImages(ctx, in)
}

func (s *Server) retrieveImages(ctx context.Context, in *pbimages.RetrieveImagesRequest) (*pbimages.RetrieveImagesResponse, error) {
	dds, err := s.Store.RetrieveMany(ctx, in.GetEntityId(), in.GetEntityTable())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "store.RetrieveMany: %v", err)
	}
	imgs := make(map[string]*pbimages.Image)
	for _, dd := range dds {
		imgs[dd.Id()] = &pbimages.Image{Uri: dd.URI(), EntityId: in.GetEntityId(), EntityTable: in.GetEntityTable()}
	}

	return &pbimages.RetrieveImagesResponse{Images: imgs}, nil
}

func (s *Server) retrieveImage(ctx context.Context, id string, respCh chan<- *pbimages.Image, errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	dd, err := s.Store.Retrieve(ctx, id)
	if err != nil {
		errCh <- status.Errorf(xerrors.Internal, "store.Find: %v", err)
		return
	}
	meta, err := dd.Metadata()
	if err != nil {
		errCh <- status.Errorf(xerrors.Internal, "data.Metadata: %v", err)
		return
	}

	respCh <- &pbimages.Image{Uri: dd.URI(), EntityId: meta.EntityId, EntityTable: meta.EntityTable}
}
