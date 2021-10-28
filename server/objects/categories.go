package objects

type Category struct {
	Title string `json:"title" gorm:"primary_key"`
}
func (Category) TableName() string {
	return "categories"
}