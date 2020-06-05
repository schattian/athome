package stage

type Stage uint64

const (
	Nil Stage = iota
	First
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
	if s == Nil {
		return Nil
	}
	return s - 1
}
