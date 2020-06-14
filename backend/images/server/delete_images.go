package server

import (
	"context"
	"sync"

	"github.com/athomecomar/athome/backend/images/store"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteImages(ctx context.Context, in *pbimages.DeleteImagesRequest) (*emptypb.Empty, error) {
	ent := in.GetEntity()
	_, err := AuthorizeThroughEntity(ctx, in.GetAccessToken(), ent.GetEntityId(), ent.GetEntityTable())
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, 1)
	done := make(chan struct{})

	dds, err := s.Store.RetrieveMany(ctx, ent.GetEntityId(), ent.GetEntityTable())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "RetrieveMany: %v", err)
	}
	for _, dd := range dds {
		wg.Add(1)
		go s.deleteImage(ctx, dd, errCh, &wg)
	}
	go func() {
		wg.Wait()
		close(done)
	}()

	for {
		select {
		case err := <-errCh:
			return nil, err
		case <-done:
			return &emptypb.Empty{}, nil
		}
	}
}

func (s *Server) deleteImage(ctx context.Context, dd store.Data, errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	err := s.Store.Delete(ctx, dd.Id())
	if err != nil {
		errCh <- status.Errorf(xerrors.Internal, "store.Delete: %v", err)
		return
	}
}
