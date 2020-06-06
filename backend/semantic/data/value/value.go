package value

import (
	"strconv"
)

type Value interface {
	Strings() []string

	GetValue() interface{}
	SetValue(interface{}) error

	Type() Type
	IsNil() bool
}

type Parser func(string) (interface{}, error)

// All defined parsers in order (to avoid false positives)
var Parsers = []Parser{
	Int64Parser,
	Float64Parser,
	BoolParser,
	StringParser,
}

func StringParser(s string) (interface{}, error) {
	return s, nil
}

func Float64Parser(s string) (interface{}, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func Int64Parser(s string) (interface{}, error) {
	f, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func BoolParser(s string) (interface{}, error) {
	f, err := strconv.ParseBool(s)
	if err != nil {
		return nil, err
	}
	return f, nil
}
