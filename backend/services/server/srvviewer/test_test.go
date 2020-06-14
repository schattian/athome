package srvviewer

import "github.com/athomecomar/athome/backend/services/internal/servicestest"

var (
	gCalendars      servicestest.GoldenCalendars
	gServices       servicestest.GoldenServices
	gAvailabilities servicestest.GoldenAvailabilities
	gEvents         servicestest.GoldenEvents
)

func init() {
	servicestest.Init(&gServices, &gCalendars, &gAvailabilities, &gEvents)
}
