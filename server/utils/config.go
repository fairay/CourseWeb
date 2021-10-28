package utils

import (
	"encoding/json"
	"os"
)

type DBConfiguration struct {
	Type	string	`json:"type"`
	Name	string	`json:"name"`

	User	string	`json:"user"`
	Password string	`json:"password"`

	Host	string	`json:"host"`
	Port   	string	`json:"port"`
}

type Configuration struct {
	DB				DBConfiguration		`json:"db"`
	LogFile			string				`json:"log_file"`
	Port			string				`json:"port"`
	TokenPassword	string				`json:"token_password"`
	TokenLiveTime	int64				`json:"token_livetime"`
}

var (
	Config Configuration
)

// TODO: returnable error
func InitConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		panic(err)
	}
}
