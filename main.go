package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/macaron.v1"
)

type Robo struct {
	Name string
	Age int
}

func main() {
	m := macaron.Classic()

	routes(m)

	m.Run()
}

func routes(m *macaron.Macaron) {
	m.Get("/", handler)

	m.Post("/", func() string {
		return "dani"
	})

	m.Get("json", jason)
}

func handler(ctx *macaron.Context) string {
	return ctx.Req.RequestURI
}

func jason(ctx *macaron.Context) []byte {
	robo := &Robo{Name: "Dani", Age: 15}

	b, err := json.Marshal(robo)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return b
}