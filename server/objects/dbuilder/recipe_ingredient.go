package dbuilder

import "api/recipes/objects"

type RecipeIngredientBuilder struct {
	Recipe_id 		int
	Ingredient_id   string
	Amount 			string
}
func newRecipeIngredientBuilder() *RecipeIngredientBuilder { 
	return new(RecipeIngredientBuilder) 
}
func (this *RecipeIngredientBuilder) Build() *objects.RecipeIngredient { 
	return &objects.RecipeIngredient{Recipe_id: this.Recipe_id, 
		Ingredient_id: this.Ingredient_id, Amount: this.Amount} 
}

type RecipeIngredientMother struct{}
func (RecipeIngredientMother) Obj0() *objects.RecipeIngredient {
	b := newRecipeIngredientBuilder()
	b.Recipe_id = 1
	b.Ingredient_id = "молоко"
	b.Amount = "1 стакан"
	return b.Build()
}
func (RecipeIngredientMother) Obj1() *objects.RecipeIngredient {
	b := newRecipeIngredientBuilder()
	b.Recipe_id = 1
	b.Ingredient_id = "соль"
	b.Amount = "1 щепотка"
	return b.Build()
}
