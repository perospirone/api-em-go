package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/danitw/api-em-go/database"
	"github.com/jinzhu/gorm"
)

type Conta struct {
	Name string
	Email string
	Password string
}

func Register(db *gorm.DB, data []byte) bool {

	var conta Conta

	err := json.Unmarshal(data, &conta)

	if err != nil {
		fmt.Println("Deu pau no json: ")
		fmt.Print(err)
		return false
	}

	a := database.Conta{Name: conta.Name, Password:conta.Password, Email: conta.Email, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	db.Create(&a)

	return true
}