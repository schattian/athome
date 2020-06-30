package server

import "github.com/athomecomar/athome/backend/messenger/internal/messengertest"

var (
	gPbUsers  messengertest.GoldenPbUsers
	gMessages messengertest.GoldenMessages
)

func init() {
	messengertest.Init(&gPbUsers, &gMessages)
}
