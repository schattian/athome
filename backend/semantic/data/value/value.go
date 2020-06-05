package value

type Value interface {
	GetValue() interface{}
	SetValue(interface{}) (Value, error)

	Type() Type
	IsNil() bool
}
