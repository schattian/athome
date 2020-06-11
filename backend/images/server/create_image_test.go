package server

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/athomecomar/athome/pb/pbimages"
	"github.com/athomecomar/xerrors"
	"github.com/athomecomar/xtest"
	"google.golang.org/grpc/status"
)

func TestServer_Iterator_createImage(t *testing.T) {
	type args struct {
		ctx     context.Context
		meta    *pbimages.CreateImageRequest_Metadata
		buffer  *bytes.Buffer
		sz      int64
		maxSize int64
	}
	tests := []struct {
		name           string
		args           args
		goldenFilename string
		wantStatus     xerrors.Code
	}{
		{
			name: "png",
			args: args{
				ctx:     context.Background(),
				buffer:  &bytes.Buffer{},
				maxSize: 10 * 1e6,
			},
			goldenFilename: "wp.png",
		},
		{
			name: "jpg",
			args: args{
				ctx:     context.Background(),
				buffer:  &bytes.Buffer{},
				maxSize: 10 * 1e6,
			},
			goldenFilename: "test.jpg",
		},
		{
			name: "jpg overflow",
			args: args{
				ctx:     context.Background(),
				buffer:  &bytes.Buffer{},
				maxSize: 1,
			},
			goldenFilename: "test.jpg",
			wantStatus:     xerrors.InvalidArgument,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := &Server{}
			reader, closer := xtest.ReaderFromGoldenFile(t, tt.goldenFilename)
			defer closer()
			buffer := make([]byte, 1024)
			var err error
			var code xerrors.Code
			var n int
			for {
				n, err = reader.Read(buffer)
				if err == io.EOF {
					break
				}
				in := &pbimages.CreateImageRequest{Corpus: &pbimages.CreateImageRequest_Chunk{Chunk: buffer[:n]}}
				_, err = s.createImage(tt.args.ctx, in, tt.args.buffer, tt.args.sz, tt.args.maxSize)
				code = status.Code(err)
				if code != xerrors.OK {
					break
				}
			}
			if code != tt.wantStatus {
				t.Errorf("Server.createImage() error = %v, wantStatus %v", err, tt.wantStatus)
				return
			}
			if tt.wantStatus == xerrors.OK {
				xtest.CmpWithGoldenFile(t, tt.args.buffer.Bytes(), tt.goldenFilename, "iterator server.createImage()")
			}
		})
	}
}
