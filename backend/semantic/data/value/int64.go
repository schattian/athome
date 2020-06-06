package value

import (
	"database/sql"
	"strconv"
)

type Int64 sql.NullInt64

func (i *Int64) Type() Type {
	return TypeInt64
}

func (s *Int64) SetValue(v interface{}) error {
	var val int64
	val, ok := v.(int64)
	if !ok {
		return errInvalidValueType(v, s.Type())
	}
	s.Int64, s.Valid = val, true
	return nil
}

func (i *Int64) GetValue() interface{} {
	return i.Int64
}

func (i *Int64) IsNil() bool {
	if i == nil {
		return true
	}

	return !i.Valid
}

func (s *Int64) Strings() (strs []string) {
	return []string{strconv.Itoa(int(s.Int64))}
}

type SlInt64 []*Int64

func (sli *SlInt64) Type() Type {
	return TypeSlInt64
}

func (sli *SlInt64) Strings() (strs []string) {
	for _, s := range *sli {
		strs = append(strs, strconv.Itoa(int(s.Int64)))
	}
	return
}

func (sli *SlInt64) IsNil() bool {
	if sli == nil {
		return true
	}

	for _, f := range *sli {
		if !f.IsNil() {
			return false
		}
	}
	return true
}

func (sli *SlInt64) GetValue() interface{} {
	var vals []interface{}
	for _, value := range *sli {
		vals = append(vals, value.GetValue())
	}
	return vals
}

func (sli *SlInt64) SetValue(v interface{}) error {
	var val []int64

	val, ok := v.([]int64)
	if !ok {
		return errInvalidValueType(v, sli.Type())
	}

	for _, value := range val {
		s := &Int64{Int64: value, Valid: true}
		*sli = append(*sli, s)
	}
	return nil
}
