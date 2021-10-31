package objects

import "encoding/json"

type Category struct {
	Title string `json:"title" gorm:"primary_key"`
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

type CategoryDTO struct {
	Title string `json:"title"`
}
