package ent

import (
	"context"
	"database/sql"
	"time"

	"github.com/athomecomar/athome/pb/pbmessenger"
	"github.com/athomecomar/storeql"
	"github.com/golang/protobuf/ptypes"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Message struct {
	Id uint64 `json:"id,omitempty"`

	SenderId   uint64
	ReceiverId uint64

	Body string

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

func NewMessage(sender, receiver uint64, body string) *Message {
	now := time.Now()
	return &Message{
		SenderId:   sender,
		ReceiverId: receiver,

		Body:      body,
		CreatedAt: Time{sql.NullTime{Time: now}},
	}
}

func StatusFromPb(in *pbmessenger.Status) (cAt, rAt, sAt Time, err error) {
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

func MessageFromPb(in *pbmessenger.Message) (*Message, error) {
	msg := &Message{
		SenderId:   in.GetSenderId(),
		ReceiverId: in.GetReceiverId(),
		Body:       in.GetBody(),
	}

	var err error
	msg.CreatedAt, msg.ReceivedAt, msg.SeenAt, err = StatusFromPb(in.GetStatus())
	if err != nil {
		return nil, errors.Wrap(err, "StatusFromPb")
	}
	return msg, nil
}

func (n *Message) ToPbStatus() (*pbmessenger.Status, error) {
	status := &pbmessenger.Status{}
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

func (n *Message) ToPb() (*pbmessenger.Message, error) {
	status, err := n.ToPbStatus()
	if err != nil {
		return nil, errors.Wrap(err, "ToPbStatus")
	}
	return &pbmessenger.Message{
		ReceiverId: n.ReceiverId,
		SenderId:   n.SenderId,
		Body:       n.Body,
		Status:     status,
	}, nil
}

func FindMessagesByUser(ctx context.Context, db *sqlx.DB, uid uint64) ([]*Message, error) {
	rows, err := storeql.WhereMany(ctx, db, &Message{}, `sender_id=$1 OR receiver_id=$1 ORDER BY created_at`, uid)
	if err != nil {
		return nil, errors.Wrap(err, "storeql.WhereMany")
	}
	defer rows.Close()
	var notifs []*Message
	for rows.Next() {
		notif := &Message{}
		err := rows.StructScan(notif)
		if err != nil {
			return nil, errors.Wrap(err, "StructScan")
		}
		notifs = append(notifs, notif)
	}
	return notifs, nil
}

// func FindMessagesByUserWithOffset(ctx context.Context, db *sqlx.DB, uid uint64, offset time.Time) ([]*Message, error) {
// 	rows, err := storeql.WhereMany(ctx, db, &Message{}, `user_id=$1 AND created_at>=$2 ORDER BY created_at`, uid, offset)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "storeql.WhereMany")
// 	}
// 	defer rows.Close()
// 	var notifs []*Message
// 	for rows.Next() {
// 		notif := &Message{}
// 		err := rows.StructScan(notif)
// 		if err != nil {
// 			return nil, errors.Wrap(err, "StructScan")
// 		}
// 		notifs = append(notifs, notif)
// 	}
// 	return notifs, nil
// }

func FindMessage(ctx context.Context, db *sqlx.DB, id, uid uint64) (*Message, error) {
	notif := &Message{}
	row := storeql.Where(ctx, db, notif, "id=$1 AND (sender_id=$2 OR receiver_id=$2)", id, uid)
	err := row.StructScan(notif)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return notif, nil
}

func FindMessageByReceiver(ctx context.Context, db *sqlx.DB, id, rid uint64) (*Message, error) {
	notif := &Message{}
	row := storeql.Where(ctx, db, notif, "id=$1 AND receiver_id=$2", id, rid)
	err := row.StructScan(notif)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return notif, nil
}

func FindMessageBySender(ctx context.Context, db *sqlx.DB, id, sid uint64) (*Message, error) {
	notif := &Message{}
	row := storeql.Where(ctx, db, notif, "id=$1 AND sender_id=$2", id, sid)
	err := row.StructScan(notif)
	if err != nil {
		return nil, errors.Wrap(err, "StructScan")
	}
	return notif, nil
}
