package dbuilder

import "api/recipes/objects"

type IngredientBuilder struct {
	Title          string 
}
func newIngredientBuilder() *IngredientBuilder { 
	return new(IngredientBuilder) 
}
func (this *IngredientBuilder) Build() *objects.Ingredient { 
	return &objects.Ingredient{Title: this.Title} 
}

type IngredientMother struct{}
func (IngredientMother) Obj0() *objects.Ingredient {
	b := newIngredientBuilder()
	b.Title = "молоко"
	return b.Build()
}
func (IngredientMother) Obj1() *objects.Ingredient {
	b := newIngredientBuilder()
	b.Title = "говядина"
	return b.Build()
}
