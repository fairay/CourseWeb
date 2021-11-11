package objects

import "encoding/json"

type RecipeIngredient struct {
	Recipe_id int    `json:"recipe" gorm:"primary_key"`
	Ingredient_id   string `json:"item" gorm:"primary_key"`
	Amount string `json:"amount"`
}

func (this *RecipeIngredient) ToDTO() *IngredientDTO {
	dto := new(IngredientDTO)
	jsonStr, _ := json.Marshal(this)
	json.Unmarshal(jsonStr, dto)

	dto.Title = this.Ingredient_id
	return dto
}
func (RecipeIngredient) ArrToDTO(src []RecipeIngredient) []IngredientDTO {
	dst := make([]IngredientDTO, len(src))
	for k, v := range src {
		dst[k] = *v.ToDTO()
	}
	return dst
}
