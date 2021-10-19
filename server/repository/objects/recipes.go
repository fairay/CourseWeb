package objects

import (
	"time"
)

type Recipes struct {
	Id          int           `json:"id" gorm:"primary_key;type:serial"`
	Author      string        `json:"author" gorm:"foreignkey:Login"`
	Title       string        `json:"title"`
	CreatedAt   time.Time     `json:"created_at"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"`
	PortionNum  int           `json:"portion_num"`
}
