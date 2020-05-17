package ent

type Validable interface {
	Validate() error
	Name() string
}
