package servicestest

import (
	"testing"

	"github.com/athomecomar/athome/backend/services/ent"
)

func AddToId(t *testing.T, s *ent.Service, qt int) *ent.Service {
	svc := CopyService(t, s)
	svc.Id = uint64(int(svc.Id) + qt)
	return svc
}
