package server

import (
	"context"

	"github.com/athomecomar/athome/backend/images/pb/pbauth"
	"github.com/athomecomar/athome/backend/images/pb/pbimages"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteImage(ctx context.Context, in *pbimages.DeleteImageRequest) (*emptypb.Empty, error) {
	auth, authCloser, err := ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()

	return s.deleteImage(ctx, auth, in)
}

func (s *Server) deleteImage(ctx context.Context, auth pbauth.AuthClient, in *pbimages.DeleteImageRequest) (*emptypb.Empty, error) {
	dd, err := s.Store.Retrieve(ctx, in.GetId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "store.Retrieve: %v", err)
	}
	meta, err := dd.Metadata()
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "store.Metadata: %v", err)
	}

	userId, err := GetUserFromAccessToken(ctx, auth, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	if meta.UserId != userId {
		return nil, status.Error(xerrors.PermissionDenied, "you are unauthorized to delete another user images")
	}
	err = s.Store.Delete(ctx, in.GetId())
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "store.Delete: %v", err)
	}
	return nil, nil
}
