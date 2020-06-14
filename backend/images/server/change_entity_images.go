package server

import (
	"bytes"
	"context"
	"io"
	"sync"

	"github.com/athomecomar/athome/backend/images/img"
	"github.com/athomecomar/athome/backend/images/store"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ChangeEntityImages(ctx context.Context, in *pbimages.ChangeEntityImagesRequest) (*emptypb.Empty, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	_, err := AuthorizeThroughEntity(ctx, in.GetAccessToken(), in.GetDestEntityId(), in.GetDestEntityTable())
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, 1)
	done := make(chan struct{})

	dds, err := s.Store.RetrieveMany(ctx, in.GetDestEntityId(), in.GetFromEntityTable())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "RetrieveMany: %v", err)
	}

	for _, dd := range dds {
		wg.Add(1)
		dd := dd
		go func() {
			defer wg.Done()
			_, _, err := s.copyImages(ctx, dd, in.GetDestEntityId(), in.GetDestEntityTable())
			if err != nil {
				errCh <- err
			}
		}()
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

func (s *Server) copyImages(ctx context.Context,
	dd store.Data,
	destEntityId uint64,
	destEntityTable string,
) (store.Data, *img.Metadata, error) {
	reader, err := s.Store.Read(dd)
	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "store.Read: %v", err)
	}
	buffer := &bytes.Buffer{}
	_, err = io.Copy(buffer, reader)
	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "io.Copy: %v", err)
	}
	meta, err := dd.Metadata()
	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "dd.Metadata: %v", err)
	}
	meta.EntityId, meta.EntityTable = destEntityId, destEntityTable
	destDd, err := s.Store.Create(ctx, meta, buffer)
	if err != nil {
		return nil, nil, status.Errorf(xerrors.Internal, "store.Create: %v", err)
	}
	return destDd, meta, nil
}
