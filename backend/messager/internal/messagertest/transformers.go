package messagertest

import (
	"testing"

	"github.com/athomecomar/athome/backend/messager/ent"
	"github.com/athomecomar/athome/pb/pbmessager"
)

// func AddToId(t *testing.T, s *ent.messager, qt int) *ent.messager {
// 	svc := Copymessager(t, s)
// 	svc.Id = uint64(int(svc.Id) + qt)
// 	return svc
// }

func MessageToPb(t *testing.T, s *ent.Message) *pbmessager.Message {
	pbn, err := s.ToPb()
	if err != nil {
		t.Fatalf("MessageToPb: %v", err)
	}
	return pbn
}
