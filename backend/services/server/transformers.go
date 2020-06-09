package server

import (
	"github.com/athomecomar/athome/backend/services/ent"
	"github.com/athomecomar/athome/backend/services/pb/pbservices"
)

func ServiceToPbServiceData(s *ent.Service) *pbservices.ServiceData {
	return &pbservices.ServiceData{
		Name:       s.Name,
		UserId:     s.UserId,
		AddressId:  s.AddressId,
		CalendarId: s.CalendarId,

		DurationInMinutes: s.DurationInMinutes,
		Price:             &pbservices.Price{Min: s.PriceMin.Float64(), Max: s.PriceMax.Float64()},
	}
}
