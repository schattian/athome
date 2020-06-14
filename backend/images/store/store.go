package store

import (
	"bytes"
	"context"
	"io"

	"github.com/athomecomar/athome/backend/images/img"
)

type Store interface {
	Create(ctx context.Context, meta *img.Metadata, data *bytes.Buffer) (Data, error)
	RetrieveMany(ctx context.Context, entityId uint64, entityTable string) ([]Data, error)
	Retrieve(ctx context.Context, id string) (Data, error)
	Delete(ctx context.Context, id string) error

	Read(Data) (io.Reader, error)
}

type Data interface {
	URI() string
	Metadata() (*img.Metadata, error)
	Id() string
}
