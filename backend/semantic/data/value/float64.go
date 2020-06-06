package value

import (
	"database/sql"
	"fmt"
)

type Float64 sql.NullFloat64

func (f *Float64) Type() Type {
	return TypeFloat64
}

func (f *Float64) SetValue(v interface{}) error {
	val, ok := v.(float64)
	if !ok {
		return errInvalidValueType(v, f.Type())
	}
	f.Float64, f.Valid = val, true
	return nil
}

func (f *Float64) GetValue() interface{} {
	return f.Float64
}

func (f *Float64) IsNil() bool {
	if f == nil {
		return true
	}

	return !f.Valid
}

func (f *Float64) Strings() (strs []string) {
	return []string{fmt.Sprintf("%f", f.Float64)}
}

type SlFloat64 []*Float64

func (slf *SlFloat64) Type() Type {
	return TypeSlFloat64
}

func (slf *SlFloat64) IsNil() bool {
	if slf == nil {
		return true
	}

	for _, f := range *slf {
		if !f.IsNil() {
			return false
		}
	}
	return true
}

func (sli *SlFloat64) Strings() (strs []string) {
	for _, s := range *sli {
		strs = append(strs, fmt.Sprintf("%f", s.Float64))
	}
	return
}

func (sli *SlFloat64) GetValue() interface{} {
	var vals []interface{}
	for _, value := range *sli {
		vals = append(vals, value.GetValue())
	}
	return vals
}

func (sli *SlFloat64) SetValue(v interface{}) error {
	val, ok := v.([]float64)
	if !ok {
		return errInvalidValueType(v, sli.Type())
	}
	for _, value := range val {
		s := &Float64{Float64: value, Valid: true}
		*sli = append(*sli, s)
	}
	return nil
}
