package stub

import "api/recipes/objects"

type CategoriesStub struct {
	arr []objects.Category
}
func NewCategotiesStub(arr ...objects.Category) *CategoriesStub {
	return &CategoriesStub{arr}
}

/*
List() []objects.Category
Find(ctg string) ([]objects.Category, error)
Get(ctg string) (objects.Category, error)
FindRecipes(ctg string) ([]objects.Recipe, error)
FindByRecipe(id_rcp int) ([]objects.Category, error)

Create(obj *objects.Category) error
AddToRecipe(id_rcp int, ctg string) error
ReplaceInRecipe(id_rcp int, arr []objects.Category) error
DelFromRecipe(id_rcp int, ctg string) error
*/