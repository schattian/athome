package ent

import (
	"context"
	"database/sql"
	"time"

	"github.com/athomecomar/athome/backend/notifier/ent/prior"
	"github.com/athomecomar/athome/pb/pbnotifier"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/storeql"
	"github.com/golang/protobuf/ptypes"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Notification struct {
	Id          uint64         `json:"id,omitempty"`
	Priority    prior.Priority `json:"priority,omitempty"`
	UserId      uint64         `json:"user_id,omitempty"`
	EntityTable string         `json:"entity_table,omitempty"`
	EntityId    uint64         `json:"entity_id,omitempty"`
	Body        string

	CreatedAt  Time `json:"created_at,omitempty"`
	ReceivedAt Time `json:"received_at,omitempty"`
	SeenAt     Time `json:"seen_at,omitempty"`
}

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

func NewNotification(uid uint64, p prior.Priority) *Notification {
	now := time.Now()
	return &Notification{
		Priority:  p,
		UserId:    uid,
		CreatedAt: Time{sql.NullTime{Time: now}},
	}
}

func StatusFromPb(in *pbnotifier.Status) (cAt, rAt, sAt Time, err error) {
	cAt.Time, err = ptypes.Timestamp(in.GetCreatedAt())
	if err != nil {
		err = errors.Wrap(err, "createdAt Timestamp")
		return
	}
	rAt.Time, err = ptypes.Timestamp(in.GetReceivedAt())
	if err != nil {
		err = errors.Wrap(err, "receivedAt Timestamp")
		return
	}
	sAt.Time, err = ptypes.Timestamp(in.GetSeenAt())
	if err != nil {
		err = errors.Wrap(err, "seenAt Timestamp")
		return
	}
	return
}

func NotificationFromPb(in *pbnotifier.Notification) (*Notification, error) {
	notif := &Notification{
		UserId:      in.GetUserId(),
		Body:        in.GetBody(),
		EntityId:    in.GetEntity().GetEntityId(),
		EntityTable: in.GetEntity().GetEntityTable(),
	}

	var err error
	notif.Priority, err = prior.FromString(in.GetPriority())
	if err != nil {
		return nil, errors.Wrap(err, "prior.FromString")
	}
	notif.CreatedAt, notif.ReceivedAt, notif.SeenAt, err = StatusFromPb(in.GetStatus())
	if err != nil {
		return nil, errors.Wrap(err, "StatusFromPb")
	}
	return notif, nil
}

func (n *Notification) ToPbStatus() (*pbnotifier.Status, error) {
	status := &pbnotifier.Status{}
	var err error
	status.CreatedAt, err = ptypes.TimestampProto(n.CreatedAt.Time)
	if err != nil {
		return nil, errors.Wrap(err, "createdAt TimestampProto")
	}

	status.ReceivedAt, err = ptypes.TimestampProto(n.ReceivedAt.Time)
	if err != nil {
		return nil, errors.Wrap(err, "receivedAt TimestampProto")
	}

	status.SeenAt, err = ptypes.TimestampProto(n.SeenAt.Time)
	if err != nil {
		return nil, errors.Wrap(err, "seenAt TimestampProto")
	}
	return status, nil
}

func (n *Notification) ToPb() (*pbnotifier.Notification, error) {
	status, err := n.ToPbStatus()
	if err != nil {
		return nil, errors.Wrap(err, "ToPbStatus")
	}
	return &pbnotifier.Notification{
		UserId: n.UserId,
		Body:   n.Body,
		Entity: pbutil.ToPbNotifierEntity(n),
		Status: status,
	}, nil
}

func FindNotificationsByUser(ctx context.Context, db *sqlx.DB, uid uint64) ([]*Notification, error) {
	rows, err := storeql.WhereMany(ctx, db, &Notification{}, `user_id=$1 ORDER BY created_at`, uid)
	if err != nil {
		return nil, errors.Wrap(err, "storeql.WhereMany")
	}
	defer rows.Close()
	var notifs []*Notification
	for rows.Next() {
		notif := &Notification{}
		err := rows.StructScan(notif)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		notifs = append(notifs, notif)
	}
	return notifs, nil
}

func FindNotificationsByUserWithOffset(ctx context.Context, db *sqlx.DB, uid uint64, offset time.Time) ([]*Notification, error) {
	rows, err := storeql.WhereMany(ctx, db, &Notification{}, `user_id=$1 AND created_at>=$2 ORDER BY created_at`, uid, offset)
	if err != nil {
		return nil, errors.Wrap(err, "storeql.WhereMany")
	}
	defer rows.Close()
	var notifs []*Notification
	for rows.Next() {
		notif := &Notification{}
		err := rows.StructScan(notif)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		notifs = append(notifs, notif)
	}
	return notifs, nil
}

func FindNotification(ctx context.Context, db *sqlx.DB, id, uid uint64) (*Notification, error) {
	notif := &Notification{}
	row := storeql.Where(ctx, db, notif, "id=$1 AND user_id=$2", id, uid)
	err := row.StructScan(notif)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return notif, nil
}
