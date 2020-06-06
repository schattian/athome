package server

import (
	"github.com/athomecomar/athome/backend/semantic/pb/pbsemantic"
	"github.com/athomecomar/athome/backend/semantic/schema"
)

func CategoryTreeToRetrieveCategoriesResponse(t schema.CategoryTree) *pbsemantic.RetrieveCategoriesResponse {
	get := &pbsemantic.RetrieveCategoriesResponse{}
	for _, b := range t {
		get.Categories = append(get.Categories, CategoryBranchToPbsemanticCategory(b))
	}
	return get
}

func CategoryBranchToPbsemanticCategory(c *schema.CategoryBranch) *pbsemantic.Category {
	pbc := CategoryToPbsemanticCategory(c.Category)
	for _, child := range c.Leafs {
		pbc.Childs = append(pbc.Childs, CategoryToPbsemanticCategory(child))
	}
	return pbc
}

func CategoryToPbsemanticCategory(c schema.Category) *pbsemantic.Category {
	return &pbsemantic.Category{
		Name:     c.GetName(),
		Id:       c.GetId(),
		ParentId: c.GetParentId(),
	}
}
