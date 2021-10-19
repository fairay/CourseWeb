package objects

type Steps struct {
	Recipe      int    `json:"recipe" gorm:"primary_key;foreignkey:Id"`
	Num         int    `json:"num" gorm:"type:serial;primary_key"`
	Description string `json:"description"`
	Title       string `json:"title"`
}
