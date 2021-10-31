package objects

type Step struct {
	Recipe      int    `json:"recipe" gorm:"primary_key;foreignkey:Id"`
	Num         int    `json:"num" gorm:"type:serial;primary_key"`
	Description string `json:"description"`
	Title       string `json:"title"`
}
func (Step) TableName() string {
	return "steps"
}

type StepDTO struct {
	Recipe      int    `json:"recipe"`
	Num         int    `json:"num"`
	Description string `json:"description"`
	Title       string `json:"title"`
}
