package notifiertest

import (
	"testing"

	"github.com/athomecomar/athome/backend/notifier/ent"
	"github.com/athomecomar/athome/pb/pbnotifier"
)

// func AddToId(t *testing.T, s *ent.notifier, qt int) *ent.notifier {
// 	svc := Copynotifier(t, s)
// 	svc.Id = uint64(int(svc.Id) + qt)
// 	return svc
// }

func NotificationToPb(t *testing.T, s *ent.Notification) *pbnotifier.Notification {
	pbn, err := s.ToPb()
	if err != nil {
		t.Fatalf("NotificationToPb: %v", err)
	}
	return pbn
}
