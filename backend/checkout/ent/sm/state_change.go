package sm

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/athomecomar/athome/backend/checkout/ent"
	"github.com/athomecomar/athome/pb/pbcheckout"
	"github.com/athomecomar/storeql"
	"github.com/golang/protobuf/ptypes"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type StateChange struct {
	Id          uint64    `json:"id,omitempty"`
	Name        StateName `json:"name,omitempty"`
	EntityId    uint64
	EntityTable string
	Stage       int64
	CreatedAt   ent.Time `json:"created_at,omitempty"`
}

func (o *StateChange) GetName() StateName      { return o.Name }
func (o *StateChange) GetStage() int64         { return o.Stage }
func (o *StateChange) GetCreatedAt() time.Time { return o.CreatedAt.Time }
func (o *StateChange) SetCreatedAt(t time.Time) {
	o.CreatedAt = ent.Time{NullTime: sql.NullTime{Time: t}}
}

func (s *StateChange) ToPb() (*pbcheckout.StateChange, error) {
	ts, err := ptypes.TimestampProto(s.CreatedAt.Time)
	if err != nil {
		return nil, errors.Wrap(err, "TimestampProto")
	}
	return &pbcheckout.StateChange{
		Stage:     s.Stage,
		Name:      string(s.Name),
		CreatedAt: ts,
	}, nil
}

func (s *StateChange) GetState(sm *StateMachine) *State {
	return sm.StateByStage(s.Stage)
}

func StateChanges(ctx context.Context, db *sqlx.DB, s Stateful) (scs []*StateChange, err error) {
	rows, err := storeql.WhereMany(ctx, db, &StateChange{}, "entity_id=$1 AND entity_table=$2", s.GetId(), s.SQLTable())
	if err != nil {
		err = errors.Wrap(err, "storeql.WhereMany")
		return
	}
	defer rows.Close()
	for rows.Next() {
		sc := &StateChange{}
		err = rows.StructScan(sc)
		if err != nil {
			err = errors.Wrap(err, "StructScan")
			return
		}
		scs = append(scs, sc)
	}
	return
}

func LatestStateChangeByRef(ctx context.Context, db *sqlx.DB, id uint64, table string) (*StateChange, error) {
	sc := &StateChange{}
	row := storeql.Where(ctx, db, sc, "entity_id=$1 AND entity_table=$2 ORDER BY created_at", id, table)
	err := row.StructScan(sc)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return sc, nil
}

func LatestStateChange(ctx context.Context, db *sqlx.DB, s Stateful) (*StateChange, error) {
	sc := &StateChange{}
	row := storeql.Where(ctx, db, sc, "entity_id=$1 AND entity_table=$2 ORDER BY created_at", s.GetId(), s.SQLTable())
	err := row.StructScan(sc)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return sc, nil
}

func NewStateChange(ctx context.Context, sId uint64, stateName StateName, s Stateful) (*StateChange, error) {
	state := s.StateMachine().StateByName(stateName)
	if state == nil {
		return nil, fmt.Errorf("state named %s doesn't exists", stateName)
	}
	stage := s.StateMachine().StageByName(stateName)
	p := &StateChange{
		EntityTable: s.SQLTable(),
		EntityId:    sId,
		Stage:       stage,
		Name:        state.Name,
	}
	p.SetCreatedAt(time.Now())
	return p, nil
}
