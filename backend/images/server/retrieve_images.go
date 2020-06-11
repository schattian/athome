package server

import (
	"context"
	"sync"

	"github.com/athomecomar/athome/backend/images/pb/pbimages"
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
	var response *pbimages.RetrieveImagesResponse
	var wg sync.WaitGroup
	idsQt := len(in.GetIds())
	respCh := make(chan *pbimages.Image, idsQt)
	errCh := make(chan error, 1)
	done := make(chan struct{})
	wg.Add(idsQt)
	for _, id := range in.GetIds() {
		go s.retrieveImage(ctx, id, respCh, errCh, &wg)
	}
	go func() {
		wg.Wait()
		close(done)
	}()

	for {
		select {
		case err := <-errCh:
			return nil, err
		case resp := <-respCh:
			response.Images = append(response.Images, resp)
		case <-done:
			return response, nil
		}
	}
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

	respCh <- &pbimages.Image{Uri: dd.URI(), UserId: meta.UserId}
}
