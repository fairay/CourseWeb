package repository

import (
	"api/recipes/objects"
	"strings"

	"github.com/jinzhu/gorm"
)

type IngredientsRep interface {
	Create(obj *objects.Ingredient) error
	Get(title string) (*objects.Ingredient, error)
	FindByRecipe(idRecipe int) ([]objects.RecipeIngredient, error)
	AddToRecipe(obj *objects.RecipeIngredient) error
	ReplaceInRecipe(id_rcp int, arr []objects.RecipeIngredient) error
	DelFromRecipe(idRecipe int, title string) error
}

type PGIngredientsRep struct {
	db *gorm.DB
}

func NewIngredientsRep(db *gorm.DB) *PGIngredientsRep {
	return &PGIngredientsRep{db}
}


func (this *PGIngredientsRep) Create(obj *objects.Ingredient) error {
	return this.db.Create(obj).Error
}
func (this *PGIngredientsRep) Get(title string) (*objects.Ingredient, error) {
	temp := new(objects.Ingredient)
	err := this.db.Where("LOWER(title) = ?", strings.ToLower(title)).Find(temp).Error
	return temp, err
}

func (this *PGIngredientsRep) FindByRecipe(idRecipe int) ([]objects.RecipeIngredient, error) {
	temp := []objects.RecipeIngredient{}
	//recipe := objects.Recipe{Id: id_rcp}

	//err := this.db.Model(&recipe).Association("Ingredients").Find(&temp).Error
	//err := this.db.Preload("Ingredients").Preload("Recipes").Find(&temp).Error
	err := this.db.Where(objects.RecipeIngredient{Recipe_id: idRecipe}).Find(&temp).Error
	return temp, err
}

func (this *PGIngredientsRep) AddToRecipe(obj *objects.RecipeIngredient) error {
	return this.db.Create(&obj).Error
}

func (this *PGIngredientsRep) ReplaceInRecipe(id_rcp int, arr []objects.RecipeIngredient) error {
	recipe := objects.Recipe{Id: id_rcp}
	return this.db.Model(&recipe).Association("RecipeIngredient").Replace(arr).Error
	// err := this.db.Model(&recipe).Association("Ingredients").Clear().Error
	// if err != nil {
	// 	return err
	// }
	// for _, obj := range arr {
	// 	err = this.db.Create(&obj).Error
	// }
	// return nil
}

func (this *PGIngredientsRep) DelFromRecipe(idRecipe int, title string) error {
	temp := &objects.RecipeIngredient{Recipe_id: idRecipe, Ingredient_id: title}
	return this.db.Delete(&temp).Error
}
