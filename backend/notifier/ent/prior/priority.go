package prior

import "fmt"

type Priority uint64

const (
	Low Priority = iota + 1
	Mid
	High
	Max
)

func (p Priority) String() string {
	switch p {
	case Low:
		return "low"
	case Mid:
		return "mid"
	case High:
		return "high"
	case Max:
		return "max"
	}
	return ""
}

func FromString(s string) (Priority, error) {
	var p Priority
	for _, constant := range []Priority{Low, Mid, High, Max} {
		if constant.String() == s {
			p = constant
			break
		}
	}
	if p == 0 {
		return 0, fmt.Errorf("invalid priority: %s", s)
	}
	return p, nil
}
