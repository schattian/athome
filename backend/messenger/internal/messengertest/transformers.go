package messengertest

import (
	"testing"

	"github.com/athomecomar/athome/backend/messenger/ent"
	"github.com/athomecomar/athome/pb/pbmessenger"
)

// func AddToId(t *testing.T, s *ent.messenger, qt int) *ent.messenger {
// 	svc := Copymessenger(t, s)
// 	svc.Id = uint64(int(svc.Id) + qt)
// 	return svc
// }

func MessageToPb(t *testing.T, s *ent.Message) *pbmessenger.Message {
	pbn, err := s.ToPb()
	if err != nil {
		t.Fatalf("MessageToPb: %v", err)
	}
	return pbn
}
