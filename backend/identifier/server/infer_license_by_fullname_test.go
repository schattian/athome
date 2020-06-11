package server

import (
	"context"
	"testing"

	"github.com/athomecomar/athome/backend/identifier/identifierconf"
	"github.com/athomecomar/athome/backend/identifier/infer"
	"github.com/athomecomar/athome/pb/pbidentifier"
	"github.com/athomecomar/semantic/semprov"
	"github.com/athomecomar/xerrors"
	"github.com/athomecomar/xtest"
	"github.com/google/go-cmp/cmp"
	"github.com/spf13/afero"
	"google.golang.org/grpc/status"
)

func TestServer_inferLicenseByFullname(t *testing.T) {
	type args struct {
		ctx      context.Context
		category *semprov.Category
		in       *pbidentifier.InferByFullnameRequest
	}
	tests := []struct {
		name        string
		args        args
		want        *pbidentifier.InferLicenseResponse
		fileContent []byte
		wantStatus  xerrors.Code
	}{
		{
			name: "medic inference by fullname strict",
			args: args{
				ctx:      context.Background(),
				category: semprov.Medic,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Foo",
					Surname: "Bar",
				},
			},
			fileContent: []byte(`{"bar foo": 123}`),
			wantStatus:  xerrors.OK,
			want:        &pbidentifier.InferLicenseResponse{License: 123},
		},
		{
			name: "medic inference by fullname strict - multiple names",
			args: args{
				ctx:      context.Background(),
				category: semprov.Medic,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Foo Baz",
					Surname: "Bar",
				},
			},
			fileContent: []byte(`{"bar foo baz": 123}`),
			wantStatus:  xerrors.OK,
			want:        &pbidentifier.InferLicenseResponse{License: 123},
		},
		{
			name: "medic inference by fullname strict - multiple names diff order",
			args: args{
				ctx:      context.Background(),
				category: semprov.Medic,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Baz Foo",
					Surname: "Bar",
				},
			},
			fileContent: []byte(`{"bar foo baz": 123}`),
			wantStatus:  xerrors.OK,
			want:        &pbidentifier.InferLicenseResponse{License: 123},
		},
		{
			name: "medic inference by fullname strict - more names on file",
			args: args{
				ctx:      context.Background(),
				category: semprov.Medic,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Baz",
					Surname: "Bar",
				},
			},
			fileContent: []byte(`{"bar foo baz": 123}`),
			wantStatus:  xerrors.OK,
			want:        &pbidentifier.InferLicenseResponse{License: 123},
		},
		{
			name: "medic inference by fullname strict - more names given",
			args: args{
				ctx:      context.Background(),
				category: semprov.Medic,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Baz Foo",
					Surname: "Bar",
				},
			},
			fileContent: []byte(`{"bar baz": 123}`),
			wantStatus:  xerrors.OK,
			want:        &pbidentifier.InferLicenseResponse{License: 123},
		},
		{
			name: "medic inference by fullname strict - multiple surnames",
			args: args{
				ctx:      context.Background(),
				category: semprov.Medic,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Foo",
					Surname: "Bar Baz",
				},
			},
			fileContent: []byte(`{"bar baz foo": 123}`),
			wantStatus:  xerrors.OK,
			want:        &pbidentifier.InferLicenseResponse{License: 123},
		},
		{
			name: "medic inference by fullname strict - multiple surnames multiple names",
			args: args{
				ctx:      context.Background(),
				category: semprov.Medic,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Foo Qux",
					Surname: "Bar Baz",
				},
			},
			fileContent: []byte(`{"bar baz foo qux": 123}`),
			wantStatus:  xerrors.OK,
			want:        &pbidentifier.InferLicenseResponse{License: 123},
		},
		{
			name: "medic inference by fullname - non alphabetics and accents",
			args: args{
				ctx:      context.Background(),
				category: semprov.Medic,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Fó'o",
					Surname: "Bar",
				},
			},
			fileContent: []byte(`{"bar foó": 123}`),
			wantStatus:  xerrors.OK,
			want:        &pbidentifier.InferLicenseResponse{License: 123},
		},
		{
			name: "medic inference by fullname - not found",
			args: args{
				ctx:      context.Background(),
				category: semprov.Medic,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Foo",
					Surname: "Bar",
				},
			},
			fileContent: []byte(`{"baz qux": 123}`),
			wantStatus:  xerrors.NotFound,
			want:        nil,
		},
		{
			name: "file read err",
			args: args{
				ctx:      context.Background(),
				category: semprov.Medic,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Foo",
					Surname: "Bar",
				},
			},
			fileContent: []byte("xd"),
			wantStatus:  xerrors.Internal,
			want:        nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := afero.NewMemMapFs()
			xtest.AddFileToFs(t, identifierconf.GetDATA_DIR()+"/"+infer.ByFullnameFilenames[tt.args.category], tt.fileContent, fs)
			s := &Server{}
			got, err := s.inferLicenseByFullname(tt.args.ctx, fs, tt.args.category, tt.args.in)
			if status.Code(err) != tt.wantStatus {
				t.Fatalf("Server.inferLicenseByFullname() error = %v, status: %v;  wantStatus %v", err, status.Code(err), tt.wantStatus)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Server.inferLicenseByFullname()errored mismatch (-want +got): %s", diff)
			}
		})
	}
}
