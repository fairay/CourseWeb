package objects

import "encoding/json"

type Category struct {
	Title string 		`json:"title" gorm:"primary_key"`

	Recipes []*Recipe	`gorm:"many2many:recipe_category;"`
}

type CategoryDTO struct {
	Title string `json:"title"`
}

func (Category) TableName() string {
	return "categories"
}
func (this *Category) ToDTO() *CategoryDTO {
	dto := new(CategoryDTO)
	jsonStr, _ := json.Marshal(this)
	json.Unmarshal(jsonStr, dto)
	return dto
}
func (Category) ArrToDTO(src []Category) []CategoryDTO {
	dst := make([]CategoryDTO, len(src))
	for k, v := range src {
		dst[k] = *v.ToDTO()
	}
	return dst
}

func (this *CategoryDTO) ToModel() *Category {
	mod := new(Category)

	jsonStr, _ := json.Marshal(this)
	json.Unmarshal(jsonStr, mod)
	return mod
}

func ToArrModel(src []CategoryDTO) *[]Category {
	dst := make([]Category, len(src))
	for i, value := range src {
		dst[i] = *value.ToModel()
	}
	return &dst
}