package messagertest

var (
	// gCalendars      GoldenCalendars
	gMessages GoldenMessages
	gPbUsers  GoldenPbUsers
	// gAvailabilities GoldenAvailabilities
	// gEvents         GoldenEvents
)

func init() {
	Init(&gPbUsers, &gMessages)
}
