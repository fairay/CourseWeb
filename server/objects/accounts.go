package objects

type Accounts struct {
	Login          string `json:"login" gorm:"primary_key"`
	Role           string `json:"role"`
	Salt           string `json:"salt"`
	HashedPassword string `json:"hashed_password"`
}
