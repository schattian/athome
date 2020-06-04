package value

type Type string

const (
	TypeInt64   Type = "int64"
	TypeBool    Type = "bool"
	TypeFloat64 Type = "float64"
	TypeString  Type = "string"

	TypeSlInt64   Type = "slint64"
	TypeSlString  Type = "slstring"
	TypeSlFloat64 Type = "slfloat64"
)
