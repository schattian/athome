package field

type Stage int64

const (
	Nil    Stage = 0
	Start  Stage = 1
	Shared Stage = 2
	End    Stage = -1
)

func (s Stage) Next(r Role) Stage {
	switch s {
	case Nil:
		s = Start
	case Start:
		s = Shared
	case Shared:
		s = End
	case End:
		s = Nil
	}
	return s
}
