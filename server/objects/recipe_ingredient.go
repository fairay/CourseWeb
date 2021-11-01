package objects

type RecipeIngredient struct {
	/*Recipe int    `gorm:"primary_key;foreignkey:Id" json:"recipe"`
	Item   string `gorm:"primary_key;foreignkey:Title" json:"item"`*/

	Recipe int    `json:"recipe" gorm:"primary_key"`
	Item   string `json:"item" gorm:"primary_key"`
	Amount string `json:"amount"`
}
