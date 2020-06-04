package value

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type String sql.NullString

func (s String) Type() Type {
	return TypeString
}
func (s String) IsNil() bool {
	return !s.Valid
}

var ErrInvalidValueType = errors.New("invalid value type given")

func errInvalidValueType(given interface{}, expected Type) error {
	return fmt.Errorf("%w; given: %T, expected: %s", ErrInvalidValueType, given, expected)
}

func (s String) SetValue(v interface{}) error {
	val, ok := v.(string)
	if !ok {
		return errInvalidValueType(v, s.Type())
	}
	s.String = val
	return nil
}

func (s String) GetValue() interface{} {
	return s.String
}

type SlString []String

func (sli SlString) Type() Type {
	return TypeSlString
}

func (sli SlString) GetValue() interface{} {
	var vals []interface{}
	for _, value := range sli {
		vals = append(vals, value.GetValue())
	}
	return vals
}

func (sli SlString) SetValue(v interface{}) error {
	val, ok := v.([]string)
	if !ok {
		return errInvalidValueType(v, sli.Type())
	}
	var xsli SlString
	for _, value := range val {
		s := sql.NullString{String: value, Valid: true}
		xsli = append(xsli, String(s))
	}
	sli = xsli
	return nil
}

func (sli SlString) IsNil() bool {
	for _, f := range sli {
		if !f.IsNil() {
			return false
		}
	}
	return true
}
