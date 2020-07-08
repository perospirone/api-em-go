package main

import (
	"github.com/danitw/api-em-go/database"
	"github.com/danitw/api-em-go/routes"
	"gopkg.in/macaron.v1"
)

func main() {
	db := database.Connection()
	defer db.Close()

	database.Migrate(db)

	m := macaron.Classic()

	routes.Routes(m)

	m.Run()
}
