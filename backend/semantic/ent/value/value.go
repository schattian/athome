package value

type Value interface {
	GetValue() interface{}
	SetValue(interface{}) error

	Type() Type
	IsNil() bool
}
