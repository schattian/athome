package schema

import (
	"errors"

	"github.com/athomecomar/athome/backend/semantic/pb/pbsemantic"
	"github.com/athomecomar/storeql"
)

type Category interface {
	storeql.Storable

	GetName() string
	SetName(string)

	GetParentId() uint64
	SetParentId(uint64)
}

func CategoryToPb(c Category) *pbsemantic.Category {
	return &pbsemantic.Category{
		Name:     c.GetName(),
		ParentId: c.GetParentId(),
	}
}

func (t CategoryTree) ToPb() map[uint64]*pbsemantic.Category {
	cats := make(map[uint64]*pbsemantic.Category)
	for _, b := range t {
		cats[b.GetId()] = b.ToPb()
	}
	return cats
}

func (b *CategoryBranch) ToPb() *pbsemantic.Category {
	pbc := CategoryToPb(b.Category)
	pbc.Childs = make(map[uint64]*pbsemantic.Category)
	for _, child := range b.Leafs {
		pbc.Childs[child.GetId()] = CategoryToPb(child)
	}
	return pbc
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
