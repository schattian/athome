package server

import "github.com/athomecomar/athome/backend/messager/internal/messagertest"

var (
	gPbUsers  messagertest.GoldenPbUsers
	gMessages messagertest.GoldenMessages
)

func init() {
	messagertest.Init(&gPbUsers, &gMessages)
}
