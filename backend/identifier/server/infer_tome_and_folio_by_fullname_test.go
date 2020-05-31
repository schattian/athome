package server

import (
	"context"
	"testing"

	"github.com/athomecomar/athome/backend/identifier/identifierconf"
	"github.com/athomecomar/athome/backend/identifier/infer"
	"github.com/athomecomar/athome/backend/identifier/pb/pbidentifier"
	"github.com/athomecomar/semantic/semprov"
	"github.com/athomecomar/xerrors"
	"github.com/athomecomar/xtest"
	"github.com/google/go-cmp/cmp"
	"github.com/spf13/afero"
	"google.golang.org/grpc/status"
)

func TestServer_inferTomeAndFolioByFullname(t *testing.T) {
	type args struct {
		ctx      context.Context
		category *semprov.Category
		in       *pbidentifier.InferByFullnameRequest
	}
	tests := []struct {
		name        string
		args        args
		want        *pbidentifier.InferTomeAndFolioResponse
		fileContent []byte
		wantStatus  xerrors.Code
	}{
		{
			name: "lawyer inference by fullname strict",
			args: args{
				ctx:      context.Background(),
				category: semprov.Lawyer,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Maria",
					Surname: "NOBLE",
				},
			},
			fileContent: []byte(`[{"tome":123,"folio":456,"surname":"NOBLE","name":"MARIA","id":576237}]`),
			wantStatus:  xerrors.OK,
			want:        &pbidentifier.InferTomeAndFolioResponse{Tome: 123, Folio: 456},
		},
		{
			name: "lawyer inference by fullname strict - multiple names",
			args: args{
				ctx:      context.Background(),
				category: semprov.Lawyer,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Maria Soledad",
					Surname: "Noble",
				},
			},
			fileContent: []byte(`[{"tome":123,"folio":456,"surname":"NOBLE","name":"MARIA SOLEDAD","id":576237}]`),
			want:        &pbidentifier.InferTomeAndFolioResponse{Tome: 123, Folio: 456},
			wantStatus:  xerrors.OK,
		},
		{
			name: "lawyer inference by fullname strict - multiple names diff order",
			args: args{
				ctx:      context.Background(),
				category: semprov.Lawyer,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Soledad Maria",
					Surname: "Noble",
				},
			},
			fileContent: []byte(`[{"tome":123,"folio":456,"surname":"NOBLE","name":"MARIA SOLEDAD","id":576237}]`),
			want:        &pbidentifier.InferTomeAndFolioResponse{Tome: 123, Folio: 456},
			wantStatus:  xerrors.OK,
		},
		{
			name: "lawyer inference by fullname strict - more names on file",
			args: args{
				ctx:      context.Background(),
				category: semprov.Lawyer,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Maria",
					Surname: "Noble",
				},
			},
			fileContent: []byte(`[{"tome":123,"folio":456,"surname":"NOBLE","name":"MARIA SOLEDAD","id":576237}]`),
			want:        &pbidentifier.InferTomeAndFolioResponse{Tome: 123, Folio: 456},
			wantStatus:  xerrors.OK,
		},
		{
			name: "lawyer inference by fullname strict - more names given",
			args: args{
				ctx:      context.Background(),
				category: semprov.Lawyer,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Maria Soledad",
					Surname: "Noble",
				},
			},
			fileContent: []byte(`[{"tome":123,"folio":456,"surname":"NOBLE","name":"MARIA","id":576237}]`),
			want:        &pbidentifier.InferTomeAndFolioResponse{Tome: 123, Folio: 456},
			wantStatus:  xerrors.OK,
		},
		{
			name: "lawyer inference by fullname strict - multiple surnames",
			args: args{
				ctx:      context.Background(),
				category: semprov.Lawyer,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Maria",
					Surname: "NOBLE ASD",
				},
			},
			fileContent: []byte(`[{"tome":123,"folio":456,"surname":"NOBLE ASD","name":"MARIA","id":576237}]`),
			want:        &pbidentifier.InferTomeAndFolioResponse{Tome: 123, Folio: 456},
			wantStatus:  xerrors.OK,
		},
		{
			name: "lawyer inference by fullname strict - multiple surnames multiple names",
			args: args{
				ctx:      context.Background(),
				category: semprov.Lawyer,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "MARIA SOLEDAD",
					Surname: "NOBLE Asd",
				},
			},
			fileContent: []byte(`[{"tome":123,"folio":456,"surname":"NOBLE ASD","name":"MARIA SOLEDAD","id":576237}]`),
			want:        &pbidentifier.InferTomeAndFolioResponse{Tome: 123, Folio: 456},
			wantStatus:  xerrors.OK,
		},
		{
			name: "lawyer inference by fullname - non alphabetics and accents",
			args: args{
				ctx:      context.Background(),
				category: semprov.Lawyer,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "MARíá' SOLEDAD",
					Surname: "NOBLE Asd",
				},
			},
			fileContent: []byte(`[{"tome":123,"folio":456,"surname":"NOBLé ASD","name":"MARIA SOLEDAD","id":576237}]`),
			want:        &pbidentifier.InferTomeAndFolioResponse{Tome: 123, Folio: 456},
			wantStatus:  xerrors.OK,
		},
		{
			name: "lawyer inference by fullname - not found",
			args: args{
				ctx:      context.Background(),
				category: semprov.Lawyer,
				in: &pbidentifier.InferByFullnameRequest{
					Name:    "Foo",
					Surname: "Bar",
				},
			},
			fileContent: []byte(`[{"tome":123,"folio":456,"surname":"NOBLé ASD","name":"MARIA SOLEDAD","id":576237}]`),
			wantStatus:  xerrors.NotFound,
			want:        nil,
		},
		{
			name: "file read err",
			args: args{
				ctx:      context.Background(),
				category: semprov.Lawyer,
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
			got, err := s.inferTomeAndFolioByFullname(tt.args.ctx, fs, tt.args.category, tt.args.in)
			if status.Code(err) != tt.wantStatus {
				t.Fatalf("Server.inferTomeAndFolioByFullname() error = %v, status: %v;  wantStatus %v", err, status.Code(err), tt.wantStatus)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Server.inferTomeAndFolioByFullname()errored mismatch (-want +got): %s", diff)
			}
		})
	}
}
