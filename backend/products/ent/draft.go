package ent

import "github.com/athomecomar/athome/backend/products/ent/stage"

type Draft struct {
	Id    uint64
	Stage stage.Stage

	// First
	Title      string
	CategoryId uint64

	// Second

	// Third
}

func (d *Draft) ValidateByStage() error {
	switch d.Stage {
	case stage.First:
	case stage.Second:
	case stage.Fourth:
	case stage.Nil:
	}
	return nil
}
