package objects

import "encoding/json"

type Ingredient struct {
	Title string 		`json:"title" gorm:"primary_key"`

	Recipes []*Recipe	`gorm:"many2many:recipe_category;"`
}

func (Ingredient) TableName() string {
	return "ingredients"
}
func (this *Ingredient) ToDTO() *IngredientDTO {
	dto := new(IngredientDTO)
	jsonStr, _ := json.Marshal(this)
	json.Unmarshal(jsonStr, dto)
	return dto
}
func (Ingredient) ArrToDTO(src []Ingredient) []IngredientDTO {
	dst := make([]IngredientDTO, len(src))
	for k, v := range src {
		dst[k] = *v.ToDTO()
	}
	return dst
}

type IngredientDTO struct {
	Title  string `json:"title"`
	Amount string `json:"amount"`
}

func (this *IngredientDTO) ToModel() *Ingredient {
	mod := new(Ingredient)

	jsonStr, _ := json.Marshal(this)
	json.Unmarshal(jsonStr, mod)
	return mod
}
