package field

type Stage int64

const (
	Nil            Stage = 0
	Start          Stage = 1
	Shared         Stage = 2
	SelectCategory Stage = 3

	End Stage = -1
)

type nextFunc func(Stage) Stage

var StateMachine = map[Role]nextFunc{
	Consumer:        nextConsumer,
	Merchant:        nextMerchant,
	ServiceProvider: nextServiceProvider,
}

func (s Stage) Next(r Role) Stage {
	return StateMachine[r](s)
}

func nextConsumer(actual Stage) (next Stage) {
	switch actual {
	case Nil:
		next = Start
	case Start:
		next = Shared
	case Shared:
		next = End
	case End:
		next = Nil
	}
	return
}

func nextServiceProvider(actual Stage) (next Stage) {
	switch actual {
	case Nil:
		next = Start
	case Start:
		next = Shared
	case Shared:
		next = SelectCategory
	case SelectCategory:
		next = End
	case End:
		next = Nil
	}
	return
}

func nextMerchant(actual Stage) (next Stage) {
	switch actual {
	case Nil:
		next = Start
	case Start:
		next = Shared
	case Shared:
		next = SelectCategory
	case SelectCategory:
		next = End
	case End:
		next = Nil
	}
	return
}
