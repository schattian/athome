package server

import (
	"bytes"
	"context"
	"io"

	"github.com/athomecomar/athome/backend/images/imageconf"
	"github.com/athomecomar/athome/backend/images/img"
	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateImage(srv pbimages.Images_CreateImageServer) error {
	ctx := srv.Context()
	resp := &pbimages.CreateImageResponse{}
	auth, authCloser, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return err
	}
	defer authCloser()

	var meta *img.Metadata
	buffer := &bytes.Buffer{}
	for {
		in, err := srv.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		err = in.Validate()
		if err != nil {
			return err
		}
		if meta == nil {
			meta, err = s.createImageMetadata(ctx, auth, in.GetMetadata())
			if err != nil {
				return err
			}
			continue
		}

		resp, err = s.createImage(ctx, in, buffer, resp.GetImage().GetSize(), imageconf.GetMAX_IMAGE_SIZE())
		if err != nil {
			return err
		}
	}
	err = img.MustExt(buffer, meta.Ext)
	if err != nil {
		return status.Errorf(xerrors.InvalidArgument, "MustExt: %v", err)
	}

	data, err := s.Store.Create(ctx, meta, buffer)
	if err != nil {
		return status.Errorf(xerrors.Internal, "store.Save: %v", err)
	}
	resp.ImageId, resp.Image.Uri = data.Id(), data.URI()
	err = srv.SendAndClose(resp)
	if err != nil {
		return err
	}
	return ctx.Err()
}

func (s *Server) createImageMetadata(ctx context.Context, c pbauth.AuthClient, meta *pbimages.CreateImageRequest_Metadata) (*img.Metadata, error) {
	id, table := meta.GetEntity().GetEntityId(), meta.GetEntity().GetEntityTable()
	_, err := AuthorizeThroughEntity(ctx, meta.GetAccessToken(), id, table)
	if err != nil {
		return nil, err
	}
	return &img.Metadata{Ext: img.Ext(meta.GetExt()), Entity: img.Entity{Id: id, Table: table}}, nil
}

func (s *Server) createImage(
	ctx context.Context,
	in *pbimages.CreateImageRequest,
	buffer *bytes.Buffer,
	sz int64,
	maxSize int64,
) (*pbimages.CreateImageResponse, error) {
	image := &pbimages.Image{}
	chunk := in.GetChunk()
	writedSz, err := buffer.Write(chunk)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "buffer.Write: %v", err)
	}
	image.Size = sz + int64(writedSz)
	if image.Size > maxSize {
		return nil, status.Errorf(xerrors.InvalidArgument, "image size overflow")
	}
	resp := &pbimages.CreateImageResponse{Image: image}

	return resp, nil
}
