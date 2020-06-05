package value

import "database/sql"

type Int64 sql.NullInt64

func (i Int64) Type() Type {
	return TypeInt64
}

func (s Int64) SetValue(v interface{}) (Value, error) {
	val, ok := v.(int64)
	if !ok {
		return nil, errInvalidValueType(v, s.Type())
	}
	s.Int64 = val
	return s, nil
}

func (i Int64) GetValue() interface{} {
	return i.Int64
}

func (i Int64) IsNil() bool {
	return !i.Valid
}

type SlInt64 []Int64

func (sli SlInt64) Type() Type {
	return TypeSlInt64
}

func (sli SlInt64) IsNil() bool {
	for _, f := range sli {
		if !f.IsNil() {
			return false
		}
	}
	return true
}

func (sli SlInt64) GetValue() interface{} {
	var vals []interface{}
	for _, value := range sli {
		vals = append(vals, value.GetValue())
	}
	return vals
}

func (sli SlInt64) SetValue(v interface{}) (Value, error) {
	val, ok := v.([]int64)
	if !ok {
		return nil, errInvalidValueType(v, sli.Type())
	}
	var xsli SlInt64
	for _, value := range val {
		s := sql.NullInt64{Int64: value, Valid: true}
		xsli = append(xsli, Int64(s))
	}
	return xsli, nil
}
