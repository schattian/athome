package stage

type Stage uint64

const (
	First Stage = iota + 1
	Second
	Third
	Fourth
)

func (s Stage) Next() Stage {
	if s == Fourth {
		return Fourth
	}
	return s + 1
}

func (s Stage) Prev() Stage {
	if s == First {
		return First
	}
	return s - 1
}
