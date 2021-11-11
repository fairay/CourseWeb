package models_test

import (
	"api/recipes/errors"
	"api/recipes/mocks"
	"api/recipes/models"
	"api/recipes/objects"
	_ "api/recipes/objects"
	"api/recipes/objects/dbuilder"
	"api/recipes/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

/// CLASSIC STYLE (Stub)

/*
Get all category - sucсess
*/
func TestGetAll(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.CategoryMother{}.All()

	mockRep := repository.NewCategotiesRep(db)
	for _, obj := range objArr {
		err := mockRep.Create(&obj);
		if  err != nil {
			panic(err)
		}
	}

	model := models.NewCategory(mockRep, nil)
	resArr := model.GetAll()

	assert.ElementsMatch(t, resArr, objArr)
}

/*
Get category - category exists
*/
func TestGet(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.CategoryMother{}.All()
	objCat := dbuilder.CategoryMother{}.Obj0()

	mockRep := repository.NewCategotiesRep(db)
	for _, obj := range objArr {
		err := mockRep.Create(&obj);
		if  err != nil {
			panic(err)
		}
	}

	model := models.NewCategory(mockRep, nil)
	res, err := model.Get(objCat.Title)

	assert.Nil(t, err, "Get has unexpected error")
	assert.Equal(t, res, objCat, "Get has unexpected error")
}

/*
Find categories - input: "" result: all categories
*/
func TestFind(t *testing.T) {
	db, err := stubConnecton()
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
	for _, obj := range objArr {
		err := mockRep.Create(&obj);
		if  err != nil {
			panic(err)
		}
	}

	model := models.NewCategory(mockRep, nil)
	resArr, err := model.Find(strFind)

	assert.Nil(t, err, "Get has unexpected error")
	assert.ElementsMatch(t, resArr, objCatArr)	
}

/*
Get recipes by categories - category and recipe with such category exist
*/
/*func TestGetRecipes(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.CategoryMother{}.All()
	objCat := dbuilder.CategoryMother{}.Obj0()
	objRcpArr := []objects.Recipe{
		*dbuilder.CategoryMother{}.Obj0(),
		*dbuilder.CategoryMother{}.Obj1(),
	}

	mockRep := repository.NewCategotiesStub(db)
	for _, obj := range objArr {
		err := mockRep.Create(&obj);
		if  err != nil {
			panic(err)
		}
	}
	for _, obj := range objArr {
		err := mockRep.Create(&obj);
		if  err != nil {
			panic(err)
		}
	}

	model := models.NewCategory(mockRep, nil)
	resArr, err := model.Find(strFind)

	assert.Nil(t, err, "Get has unexpected error")
	assert.ElementsMatch(t, resArr, objCatArr)	
}*/


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
