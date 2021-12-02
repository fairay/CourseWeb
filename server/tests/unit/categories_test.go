package models_test

import (
	"api/recipes/errors"
	"api/recipes/mocks"
	"api/recipes/models"
	"api/recipes/objects"
	_ "api/recipes/objects"
	"api/recipes/objects/dbuilder"
	"api/recipes/repository"
	"api/recipes/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

/// CLASSIC STYLE (Stub)

/*
Check category exists - category exists
*/
func TestExistsCategory(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.CategoryMother{}.All()
	objCat := dbuilder.CategoryMother{}.Obj0()

	mockRep := repository.NewCategotiesRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil {
		panic(err)
	}

	model := models.NewCategory(mockRep, nil)
	res := model.IsExists(objCat.Title)

	assert.Equal(t, res, true, "Wrong exists bool")
}

/*
category not exists
*/
func TestNotExistsCategory(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.CategoryMother{}.All()

	mockRep := repository.NewCategotiesRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil {
		panic(err)
	}

	model := models.NewCategory(mockRep, nil)
	res := model.IsExists(nWord)

	assert.Equal(t, res, false, "Wrong exists bool")
}

/*
Get all category - sucсess
*/
func TestGetAllCategory(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.CategoryMother{}.All()

	mockRep := repository.NewCategotiesRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil {
		panic(err)
	}

	model := models.NewCategory(mockRep, nil)
	resArr := model.GetAll()

	assert.ElementsMatch(t, resArr, objArr)
}

/*
Get category - category exists
*/
func TestGetCategory(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.CategoryMother{}.All()
	objCat := dbuilder.CategoryMother{}.Obj0()

	mockRep := repository.NewCategotiesRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil {
		panic(err)
	}

	model := models.NewCategory(mockRep, nil)
	res, err := model.Get(objCat.Title)

	assert.Nil(t, err, "Get has unexpected error")
	assert.Equal(t, res, objCat, "Get has unexpected error")
}

/*
Find categories - input: "" result: all categories
*/
func TestFindCategory(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.CategoryMother{}.All()
	objCatArr := []objects.Category{
		*dbuilder.CategoryMother{}.Obj0(),
		*dbuilder.CategoryMother{}.Obj1(),
	}
	strFind := ""

	mockRep := repository.NewCategotiesRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil {
		panic(err)
	}

	model := models.NewCategory(mockRep, nil)
	resArr, err := model.Find(strFind)

	assert.Nil(t, err, "Get has unexpected error")
	assert.ElementsMatch(t, resArr, objCatArr)
}
func TestFindOneCategory(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.CategoryMother{}.All()
	objCat := dbuilder.CategoryMother{}.Obj0()
	expArr := []objects.Category{*objCat}

	mockRep := repository.NewCategotiesRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil {
		panic(err)
	}

	strFind := objCat.Title
	model := models.NewCategory(mockRep, nil)
	resArr, err := model.Find(strFind)

	assert.Nil(t, err, "Get has unexpected error")
	assert.ElementsMatch(t, resArr, expArr)
}

/*
Get categories by recipe - category and recipe with such recipe exist
*/
func TestGetByRecipeCategory(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	catArr := dbuilder.CategoryMother{}.All()
	recArr := dbuilder.RecipeMother{}.All()

	mockCat := repository.NewCategotiesRep(db)
	err = mockCat.CreateList(catArr)
	if err != nil {
		panic(err)
	}

	mockRec := repository.NewRecipesRep(db)
	err = mockRec.CreateList(recArr)
	if err != nil {
		panic(err)
	}

	mockCat.AddToRecipe(recArr[0].Id, catArr[0].Title)
	mockCat.AddToRecipe(recArr[1].Id, catArr[0].Title)
	mockCat.AddToRecipe(recArr[1].Id, catArr[1].Title)
	expArr := []objects.Category{catArr[0]}

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Category = models.NewCategory(mockCat, allM)
	resArr, err := allM.Category.GetByRecipe(recArr[0].Id)

	assert.Nil(t, err, "Get has unexpected error")
	assert.ElementsMatch(t, resArr, expArr)
}

/*
Get recipes by categories - category and recipe with such category exist
*/
func TestGetRecipesCategory(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	catArr := dbuilder.CategoryMother{}.All()
	recArr := dbuilder.RecipeMother{}.All()

	mockCat := repository.NewCategotiesRep(db)
	err = mockCat.CreateList(catArr)
	if err != nil {
		panic(err)
	}

	mockRec := repository.NewRecipesRep(db)
	err = mockRec.CreateList(recArr)
	if err != nil {
		panic(err)
	}

	mockCat.AddToRecipe(recArr[0].Id, catArr[0].Title)
	mockCat.AddToRecipe(recArr[1].Id, catArr[0].Title)
	mockCat.AddToRecipe(recArr[1].Id, catArr[1].Title)
	expArr := []objects.Recipe{recArr[0], recArr[1]}

	model := models.NewCategory(mockCat, nil)
	resArr, err := model.GetRecipes(catArr[0].Title)

	assert.Nil(t, err, "Get has unexpected error")
	tests.CompareRecipes(t, resArr, expArr, "Get has unexpected object array")
}

/// LONDON STYLE (Mock)

/*
Create recipe - successful operation
*/
func TestCreateCategory(t *testing.T) {

	mockRep := new(mocks.CategoriesRep)
	model := models.NewCategory(mockRep, nil)
	obj := dbuilder.CategoryMother{}.Obj0()

	mockRep.On("Get", obj.Title).Return(nil, errors.RecordNotFound).Once()
	mockRep.On("Create", obj).Return(nil).Once()

	err := model.Create(obj)

	assert.Nil(t, err, "Create category have unexpected error")
	mockRep.AssertExpectations(t)
}

/*
Add category - successful operation, new category does not exist
*/
func TestAddCategoryRecipe(t *testing.T) {
	mockRec := new(mocks.RecipesRep)
	mockCat := new(mocks.CategoriesRep)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Category = models.NewCategory(mockCat, allM)

	objCat := dbuilder.CategoryMother{}.Obj0()
	objRcp := dbuilder.RecipeMother{}.Obj0()

	mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()

	mockCat.On("Get", objCat.Title).Return(nil, errors.RecordNotFound).Twice()

	mockCat.On("Create", objCat).Return(nil).Once()

	mockCat.On("AddToRecipe", objRcp.Id, objCat.Title).Return(nil).Once()

	err := allM.Category.AddToRecipe(objRcp.Id, objCat)

	assert.Nil(t, err, "Addition a new category has unexpected error")
	mockRec.AssertExpectations(t)
	mockCat.AssertExpectations(t)
}

/*
Delete category - successful operation (recipe/category exist)
*/
func TestDelCategoryRecipe(t *testing.T) {
	mockRec := new(mocks.RecipesRep)
	mockCat := new(mocks.CategoriesRep)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Category = models.NewCategory(mockCat, allM)

	objCat := dbuilder.CategoryMother{}.Obj0()
	objRcp := dbuilder.RecipeMother{}.Obj0()

	mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()

	mockCat.On("Get", objCat.Title).Return(objCat, nil).Once()

	mockCat.On("DelFromRecipe", objRcp.Id, objCat.Title).Return(nil).Once()

	err := allM.Category.DelFromRecipe(objRcp.Id, objCat)

	assert.Nil(t, err, "Deletion a category from recipe has unexpected error")
	mockRec.AssertExpectations(t)
	mockCat.AssertExpectations(t)
}

/*
Post category to recipe - successful operation
*/
func TestPostCategoryRecipe(t *testing.T) {
	mockRec := new(mocks.RecipesRep)
	mockCat := new(mocks.CategoriesRep)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Category = models.NewCategory(mockCat, allM)

	objCatArr := []objects.Category{
		*dbuilder.CategoryMother{}.Obj0(),
	}

	objRcp := dbuilder.RecipeMother{}.Obj0()

	mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()

	for _, category := range objCatArr {
		mockCat.On("Get", category.Title).Return(nil, errors.RecordNotFound).Twice()
		mockCat.On("Create", &category).Return(nil).Once()
	}

	mockCat.On("ReplaceInRecipe", objRcp.Id, objCatArr).Return(nil).Once()

	err := allM.Category.PostToRecipe(objRcp.Id, &objCatArr)

	assert.Nil(t, err, "Post a new category/ies has unexpected error")
	mockRec.AssertExpectations(t)
	mockCat.AssertExpectations(t)
}
