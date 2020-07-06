package main

import (
	"github.com/danitw/api-em-go/conection"
	"github.com/danitw/api-em-go/routes"

	"gopkg.in/macaron.v1"
)



func main() {
	db := conection.Conection()
	defer db.Close()

	m := macaron.Classic()

	routes.Routes(m)

	m.Run()
}
