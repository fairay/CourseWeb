package objects

type Ingredients struct {
	Title string `json:"title" gorm:"primary_key"`
}
