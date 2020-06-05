package schema

import (
	"errors"

	"github.com/athomecomar/storeql"
)

type Category interface {
	storeql.Storable

	GetName() string
	SetName(string)

	GetParentId() uint64
	SetParentId(uint64)
}

type CategoryBranch struct {
	Category

	Leafs []Category
}

type CategoryTree []*CategoryBranch

func AddCategoryToTree(t CategoryTree, c Category) (CategoryTree, error) {
	if c.GetParentId() == 0 {
		t = append(t, &CategoryBranch{Category: c})
		return t, nil
	}
	for _, b := range t {
		if b.GetId() == c.GetId() {
			return nil, errors.New("category id repeated on the tree")
		}
		if b.GetId() == c.GetParentId() {
			b.Leafs = append(b.Leafs, c)
			return t, nil
		}
	}

	return t, nil
}
