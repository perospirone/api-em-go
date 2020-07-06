package conection

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Conection() *gorm.DB {
	db, err := gorm.Open("postgres", "host=0.0.0.0 port=5432 user=postgres dbname=teste password=postgres sslmode=disable")

	if err != nil {
		fmt.Println("Erro: ")
		fmt.Println(err)
	}

	return db
}
