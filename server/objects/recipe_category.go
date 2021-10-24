package objects

type RecipeCategory struct {
	Recipe   int    `gorm:"primary_key;foreignkey:Id" json:"recipe"`
	Category string `gorm:"primary_key;foreignkey:Id" json:"category"`
}
