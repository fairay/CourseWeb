package objects

type RecipeIngredient struct {
	Recipe int    `gorm:"primary_key;foreignkey:Id" json:"recipe"`
	Item   string `gorm:"primary_key;foreignkey:Title" json:"item"`
}
