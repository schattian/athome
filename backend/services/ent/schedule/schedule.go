package schedule

import "time"

type Scheduleable interface {
	GetDayOfWeek() time.Weekday

	GetStartHour() int64
	GetStartMinute() int64

	GetEndHour() int64
	GetEndMinute() int64
}

func absTs(h, m int64) int64 {
	return h*60 + m
}

func start(a Scheduleable) int64 {
	return absTs(a.GetStartHour(), a.GetStartMinute())
}

func end(a Scheduleable) int64 {
	return absTs(a.GetEndHour(), a.GetEndMinute())
}

type Comparator func(a, b Scheduleable) bool

func CompareWithSlice(cmp Comparator, a Scheduleable, as ...Scheduleable) bool {
	for _, ai := range as {
		if cmp(a, ai) {
			return true
		}
	}
	return false
}

func ComparePairwise(cmp Comparator, ts ...Scheduleable) bool {
	for i, ti := range ts {
		for j, tj := range ts {
			if j == i {
				continue
			}
			if cmp(ti, tj) {
				return true
			}
		}
	}
	return false
}

func NotNullIntersection(ti, tj Scheduleable) bool {
	if ti.GetDayOfWeek() != tj.GetDayOfWeek() {
		return false
	}

	if start(ti) >= end(tj) {
		return false
	}
	if end(ti) <= start(tj) {
		return false
	}

	return true
}

func IsContained(ti, tj Scheduleable) bool {
	if ti.GetDayOfWeek() != tj.GetDayOfWeek() {
		return false
	}
	if start(ti) <= start(tj) {
		return false
	}
	if end(ti) >= end(tj) {
		return false
	}

	return true
}
