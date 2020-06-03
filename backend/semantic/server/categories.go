package server

import (
	"github.com/athomecomar/athome/backend/semantic/ent"
	"github.com/athomecomar/athome/backend/semantic/pb/pbsemantic"
)

func CategoryTreeToGetCategoriesResponse(t ent.CategoryTree) *pbsemantic.GetCategoriesResponse {
	get := &pbsemantic.GetCategoriesResponse{}
	for _, b := range t {
		get.Categories = append(get.Categories, CategoryBranchToPbsemanticCategory(b))
	}
	return get
}

func CategoryBranchToPbsemanticCategory(c *ent.CategoryBranch) *pbsemantic.Category {
	pbc := CategoryToPbsemanticCategory(c.Category)
	for _, child := range c.Leafs {
		pbc.Childs = append(pbc.Childs, CategoryToPbsemanticCategory(child))
	}
	return pbc
}

func CategoryToPbsemanticCategory(c ent.Category) *pbsemantic.Category {
	return &pbsemantic.Category{
		Name:     c.GetName(),
		Id:       c.GetId(),
		ParentId: c.GetParentId(),
	}
}
