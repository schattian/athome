package server

import (
	"context"
	"sync"

	"github.com/athomecomar/athome/backend/images/store"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/xerrors"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

func (s *Server) CloneImages(ctx context.Context, in *pbimages.CloneImagesRequest) (*pbimages.CloneImagesResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	_, err := AuthorizeThroughEntity(ctx, in.GetAccessToken(), in.GetDestEntityId(), in.GetEntityTable())
	if err != nil {
		return nil, err
	}
	_, err = AuthorizeThroughEntity(ctx, in.GetAccessToken(), in.GetFromEntityId(), in.GetEntityTable())
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, 1)
	done := make(chan struct{})

	dds, err := s.Store.RetrieveMany(ctx, in.GetFromEntityId(), in.GetEntityTable())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "RetrieveMany: %v", err)
	}
	resp := &pbimages.CloneImagesResponse{}
	resp.Images = make(map[string]*pbimages.Image)
	var lock sync.RWMutex
	for _, dd := range dds {
		wg.Add(1)
		dd := dd
		go s.cloneImages(ctx, dd, in.GetDestEntityId(), in.GetEntityTable(), resp, &lock, &wg, errCh)
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
			return resp, nil
		}
	}
}

func (s *Server) cloneImages(
	ctx context.Context,
	dd store.Data,
	destEntityId uint64,
	destEntityTable string,
	resp *pbimages.CloneImagesResponse,
	lock *sync.RWMutex,
	wg *sync.WaitGroup,
	errCh chan<- error,
) {
	destDd, meta, err := s.copyImages(ctx, dd, destEntityId, destEntityTable)
	if err != nil {
		errCh <- errors.Wrap(err, "copyImages")
	}
	go s.deleteImage(ctx, dd, errCh, wg)
	lock.Lock()
	defer lock.Unlock()
	resp.Images[destDd.Id()] = &pbimages.Image{Uri: destDd.URI(), EntityId: meta.EntityId, EntityTable: meta.EntityTable}
}
