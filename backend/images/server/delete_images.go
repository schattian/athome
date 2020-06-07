package server

import (
	"context"
	"sync"

	"github.com/athomecomar/athome/backend/images/pb/pbimages"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteImages(ctx context.Context, in *pbimages.DeleteImagesRequest) (*emptypb.Empty, error) {
	auth, authCloser, err := ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()
	userId, err := GetUserFromAccessToken(ctx, auth, in.GetAccessToken())
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, 1)
	done := make(chan struct{})
	wg.Add(len(in.GetIds()))
	for _, id := range in.GetIds() {
		go s.deleteImage(ctx, userId, id, errCh, &wg)
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

func (s *Server) deleteImage(ctx context.Context, userId uint64, id string, errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	dd, err := s.Store.Retrieve(ctx, id)
	if err != nil {
		errCh <- status.Errorf(xerrors.Internal, "store.Retrieve: %v", err)
		return
	}
	meta, err := dd.Metadata()
	if err != nil {
		errCh <- status.Errorf(xerrors.Internal, "store.Metadata: %v", err)
		return
	}
	if meta.UserId != userId {
		errCh <- status.Error(xerrors.PermissionDenied, "you are unauthorized to delete another user images")
		return
	}
	err = s.Store.Delete(ctx, id)
	if err != nil {
		errCh <- status.Errorf(xerrors.Internal, "store.Delete: %v", err)
		return
	}
}
