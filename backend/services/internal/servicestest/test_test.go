package servicestest

var (
	gCalendars      GoldenCalendars
	gServices       GoldenServices
	gAvailabilities GoldenAvailabilities
	gEvents         GoldenEvents
)

func init() {
	Init(&gServices, &gCalendars, &gAvailabilities, &gEvents)
}
