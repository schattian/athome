package value

import "database/sql"

type Float64 sql.NullFloat64

func (f Float64) Type() Type {
	return TypeFloat64
}

func (f Float64) SetValue(v interface{}) error {
	val, ok := v.(float64)
	if !ok {
		return errInvalidValueType(v, f.Type())
	}
	f.Float64 = val
	return nil
}

func (f Float64) GetValue() interface{} {
	return f.Float64
}

func (f Float64) IsNil() bool {
	return !f.Valid
}

type SlFloat64 []Float64

func (slf SlFloat64) Type() Type {
	return TypeSlFloat64
}

func (slf SlFloat64) IsNil() bool {
	for _, f := range slf {
		if !f.IsNil() {
			return false
		}
	}
	return true
}

func (sli SlFloat64) GetValue() interface{} {
	var vals []interface{}
	for _, value := range sli {
		vals = append(vals, value.GetValue())
	}
	return vals
}

func (sli SlFloat64) SetValue(v interface{}) error {
	val, ok := v.([]float64)
	if !ok {
		return errInvalidValueType(v, sli.Type())
	}
	var xsli SlFloat64
	for _, value := range val {
		s := sql.NullFloat64{Float64: value, Valid: true}
		xsli = append(xsli, Float64(s))
	}
	sli = xsli
	return nil
}
