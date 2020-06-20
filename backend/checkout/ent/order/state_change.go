package order

import (
	"time"

	"github.com/athomecomar/athome/backend/checkout/ent/sm"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/storeql"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
)

type StateChange interface {
	GetName() string
	GetCreatedAt() time.Time
	GetStage() int64
	GetState() *sm.State
	storeql.Storable
}

func StateChangeToPb(s StateChange) (*pbcheckout.StateChange, error) {
	ts, err := ptypes.TimestampProto(s.GetCreatedAt())
	if err != nil {
		return nil, errors.Wrap(err, "TimestampProto")
	}
	return &pbcheckout.StateChange{
		Stage:     s.GetStage(),
		Name:      s.GetName(),
		CreatedAt: ts,
	}, nil
}
