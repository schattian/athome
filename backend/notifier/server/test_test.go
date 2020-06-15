package server

import "github.com/athomecomar/athome/backend/notifier/internal/notifiertest"

var (
	gPbUsers       notifiertest.GoldenPbUsers
	gNotifications notifiertest.GoldenNotifications
)

func init() {
	notifiertest.Init(&gPbUsers, &gNotifications)
}
