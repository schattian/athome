package server

import (
	"bytes"
	"context"
	"io"
	"sync"

	"github.com/athomecomar/athome/backend/images/store"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/xerrors"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteImages(ctx context.Context, in *pbimages.DeleteImagesRequest) (*emptypb.Empty, error) {
	_, err := AuthorizeThroughEntity(ctx, in.GetAccessToken(), in.GetEntityId(), in.GetEntityTable())
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, 1)
	done := make(chan struct{})

	dds, err := s.Store.RetrieveMany(ctx, in.GetEntityId(), in.GetEntityTable())
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

func (s *Server) changeEntityImages(ctx context.Context, in *pbimages.ChangeEntityImagesRequest,
	dd store.Data,
	wg *sync.WaitGroup,
	errCh chan<- error,
) (*emptypb.Empty, error) {
	defer wg.Done()
	reader, err := s.Store.Read(dd)
	if err != nil {
		errCh <- errors.Wrap(err, "store.Read")
	}
	buffer := &bytes.Buffer{}
	_, err = io.Copy(buffer, reader)
	if err != nil {
		errCh <- errors.Wrap(err, "io.Copy")
	}
	meta, err := dd.Metadata()
	if err != nil {
		errCh <- errors.Wrap(err, "dd.Metadata")
	}
	meta.EntityId, meta.EntityTable = in.GetDestEntityId(), in.GetDestEntityTable()
	_, err = s.Store.Create(ctx, meta, buffer)
	if err != nil {
		errCh <- errors.Wrap(err, "Create")
	}
	return &emptypb.Empty{}, nil
}
