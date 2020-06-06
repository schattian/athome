package value

import "github.com/pkg/errors"

type Type string

const (
	TypeInt64     Type = "int64"
	TypeBool      Type = "bool"
	TypeFloat64   Type = "float64"
	TypeString    Type = "string"
	TypeSlInt64   Type = "slint64"
	TypeSlString  Type = "slstring"
	TypeSlFloat64 Type = "slfloat64"
)

func Parse(t Type, s ...string) (v interface{}, err error) {
	switch len(s) {
	case 0:
		err = errors.New("no values given to parser")
	case 1:
		v, err = parseSingle(t, s[0])
	default:
		v, err = parseMultiple(t, s)
	}
	return
}

func parseMultiple(t Type, strs []string) (v interface{}, err error) {
	var sl []interface{}
	var parser Parser
	switch t {
	case TypeSlInt64:
		parser = Int64Parser
	case TypeSlString:
		parser = StringParser
	case TypeSlFloat64:
		parser = Float64Parser
	}
	for _, str := range strs {
		v, err = parser(str)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing with %s type given", t)
		}
		sl = append(sl, v)
	}
	return sl, nil
}

func parseSingle(t Type, s string) (v interface{}, err error) {
	switch t {
	case TypeInt64:
		v, err = Int64Parser(s)
	case TypeBool:
		v, err = BoolParser(s)
	case TypeFloat64:
		v, err = Float64Parser(s)
	case TypeString:
		v, err = StringParser(s)
	}
	return
}

var (
	NilFloat64 = &Float64{Valid: false}
	NilString  = &String{Valid: false}
	NilBool    = &Bool{Valid: false}
	NilInt64   = &Int64{Valid: false}

	NilSlInt64   = &SlInt64{}
	NilSlFloat64 = &SlFloat64{}
	NilSlString  = &SlString{}
)
