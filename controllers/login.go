package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Conta2 struct {
	Email    string
	Password string
}

func Login(db *gorm.DB, data []byte) string {
	var c Conta2

	err := json.Unmarshal(data, &c)

	if err != nil {
		fmt.Println("Deu pau no json: ")
		fmt.Print(err)
	}

	fmt.Println(c)

	return "true"

}
