package dbuilder

import "api/recipes/objects"

type CategoryBuilder struct {
	Title string
}
func newCategoryBuilder() *CategoryBuilder { 
	return new(CategoryBuilder) 
}
func (this *CategoryBuilder) Build() *objects.Category { 
	return &objects.Category{Title: this.Title} 
}

type CategoryMother struct{}
func (CategoryMother) Obj0() *objects.Category {
	b := newCategoryBuilder()
	b.Title = "супы"
	return b.Build()
}
func (CategoryMother) Obj1() *objects.Category {
	b := newCategoryBuilder()
	b.Title = "салаты"
	return b.Build()
}
