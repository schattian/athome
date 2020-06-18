package ent

import (
	"database/sql"
	"time"

	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
)

type Time struct {
	sql.NullTime
}

func (t Time) UnmarshalJSON(b []byte) error {
	// "2006-01-02T15:04:05Z"
	var err error
	t.Time, err = time.Parse(`"`+time.RFC3339+`"`, string(b))
	if err != nil {
		return errors.Wrap(err, "parse")
	}
	return nil
}

type timestampable interface {
	GetCreatedAt() time.Time
	SetCreatedAt(time.Time)

	GetUpdatedAt() time.Time
	SetUpdatedAt(time.Time)
}

func GetTimestamp(ts timestampable) (*pbcheckout.Timestamp, error) {
	cAt, err := ptypes.TimestampProto(ts.GetCreatedAt())
	if err != nil {
		return nil, errors.Wrap(err, "createdAt TimestampProto")
	}
	uAt, err := ptypes.TimestampProto(ts.GetUpdatedAt())
	if err != nil {
		return nil, errors.Wrap(err, "createdAt TimestampProto")
	}
	return &pbcheckout.Timestamp{
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}, nil
}
