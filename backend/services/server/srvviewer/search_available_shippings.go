package srvviewer

import (
	"context"
	"sync"

	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/server"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/athome/pb/pbsemantic"
	"github.com/athomecomar/athome/pb/pbservices"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/status"
)

func (s *Server) SearchAvailableShippings(ctx context.Context, in *pbservices.SearchAvailableShippingsRequest) (*pbservices.SearchAvailableShippingsResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}

	sem, semCloser, err := pbutil.ConnSemanticServiceProviders(ctx)
	defer semCloser()

	img, imgCloser, err := pbutil.ConnImages(ctx)
	if err != nil {
		return nil, err
	}
	defer imgCloser()

	users, usersCloser, err := pbutil.ConnUsersViewer(ctx)
	if err != nil {
		return nil, err
	}
	defer usersCloser()

	return s.searchAvailableShippings(ctx, db, sem, users, img, in)
}

func (s *Server) searchAvailableShippings(ctx context.Context, db *sqlx.DB,
	sem pbsemantic.ServiceProvidersClient, users pbusers.ViewerClient, img pbimages.ImagesClient,
	in *pbservices.SearchAvailableShippingsRequest) (*pbservices.SearchAvailableShippingsResponse, error) {
	dow, err := ent.DayOfWeekByName(in.GetDow())
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "DayOfWeekByName: %v", err)
	}
	svcs, err := ent.AvailableServicesByCategory(ctx, db, dow, in.GetStart(), in.GetEnd(), 12)
	if err != nil {
		return nil, status.Errorf(xerrors.InvalidArgument, "DayOfWeekByName: %v", err)
	}

	resp := &pbservices.SearchAvailableShippingsResponse{}
	var wg sync.WaitGroup
	errCh := make(chan error, 1)
	done := make(chan struct{})
	resp.Services = make(map[uint64]*pbservices.ServiceSearchResult)
	var lock sync.RWMutex
	for _, svc := range svcs {
		wg.Add(1)
		svc := svc
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			res, err := svc.ToPbSearchResult(ctx, users)
			if err != nil {
				errCh <- status.Errorf(xerrors.Internal, "ToPbSearchResult: %v", err)
			}
			lock.Lock()
			defer lock.Unlock()
			resp.Services[svc.Id] = res
		}(&wg)
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
