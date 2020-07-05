package sm

import (
	"context"
	"time"

	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/storeql"
	"github.com/golang/protobuf/ptypes"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type StateChange interface {
	GetName() string
	GetCreatedAt() time.Time
	GetStage() int64
	GetState() *State
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

func StateChanges(ctx context.Context, db *sqlx.DB, s Stateful) (scs []StateChange, err error) {
	rows, err := storeql.WhereMany(ctx, db, s.StateChange(), "entity_id=$1", s.GetId())
	if err != nil {
		err = errors.Wrap(err, "storeql.WhereMany")
		return
	}
	defer rows.Close()
	for rows.Next() {
		sc := s.StateChange()
		err = rows.StructScan(sc)
		if err != nil {
			err = errors.Wrap(err, "StructScan")
			return
		}
		scs = append(scs, sc)
	}
	return
}

func LatestStateChange(ctx context.Context, db *sqlx.DB, s Stateful) (sc StateChange, err error) {
	sc = s.StateChange()
	row := storeql.Where(ctx, db, sc, "order_id=$1 ORDER BY created_at", s.GetId())
	err = row.StructScan(sc)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return
}
