package value

import (
	"database/sql"
	"strconv"
)

type Bool sql.NullBool

func (b *Bool) Type() Type {
	return TypeBool
}

func (b *Bool) IsNil() bool {
	if b == nil {
		return true
	}

	return !b.Valid
}

func (f *Bool) SetValue(v interface{}) error {
	val, ok := v.(bool)
	if !ok {
		return errInvalidValueType(v, f.Type())
	}
	f.Bool, f.Valid = val, true
	return nil
}

func (b *Bool) Strings() (strs []string) {
	return []string{strconv.FormatBool(b.Bool)}
}

func (f *Bool) GetValue() interface{} {
	return f.Bool
}
