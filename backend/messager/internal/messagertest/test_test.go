package messagertest

var (
	gMessages GoldenMessages
	gPbUsers  GoldenPbUsers
)

func init() {
	Init(&gPbUsers, &gMessages)
}
