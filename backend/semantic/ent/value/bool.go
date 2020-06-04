package value

import "database/sql"

type Bool sql.NullBool

func (b Bool) Type() Type {
	return TypeBool
}
func (b Bool) IsNil() bool {
	return !b.Valid
}

func (f Bool) SetValue(v interface{}) error {
	val, ok := v.(bool)
	if !ok {
		return errInvalidValueType(v, f.Type())
	}
	f.Bool = val
	return nil
}

func (f Bool) GetValue() interface{} {
	return f.Bool
}
