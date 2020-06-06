package value

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type String sql.NullString

func (s *String) Type() Type {
	return TypeString
}

func (s *String) IsNil() bool {
	if s == nil {
		return true
	}
	return !s.Valid
}

var ErrInvalidValueType = errors.New("invalid value type given")

func errInvalidValueType(given interface{}, expected Type) error {
	return fmt.Errorf("%w; given: %T, expected: %s", ErrInvalidValueType, given, expected)
}

func (s *String) SetValue(v interface{}) error {
	val, ok := v.(string)
	if !ok {
		return errInvalidValueType(v, s.Type())
	}
	s.String, s.Valid = val, true
	return nil
}

func (s *String) GetValue() interface{} {
	return s.String
}

type SlString []*String

func (sli *SlString) Type() Type {
	return TypeSlString
}

func (sli *SlString) GetValue() interface{} {
	var vals []interface{}
	for _, value := range *sli {
		vals = append(vals, value.GetValue())
	}
	return vals
}

func (sli *SlString) SetValue(v interface{}) error {
	val, ok := v.([]string)
	if !ok {
		return errInvalidValueType(v, sli.Type())
	}
	for _, value := range val {
		s := &String{String: value, Valid: true}
		*sli = append(*sli, s)
	}
	return nil
}

func (s *String) Strings() (strs []string) {
	return []string{s.String}
}

func (sli *SlString) Strings() (strs []string) {
	for _, s := range *sli {
		strs = append(strs, s.String)
	}
	return
}

func (sli *SlString) IsNil() bool {
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
