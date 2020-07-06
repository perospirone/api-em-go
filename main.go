package main

import (
	"encoding/json"
	"fmt"

	"github.com/danitw/api-em-go/conection"

	"gopkg.in/macaron.v1"
)

type Robo struct {
	Name string
	Age  int
}

type Email struct {
	Assunto  string `json:"assunto"`
	Mensagem string `json:"mensagem"`
}

func main() {
	conection.Conection()
	/*m := macaron.Classic()

	routes(m)

	m.Run()*/
}

func routes(m *macaron.Macaron) {
	m.Get("/", handler)

	m.Get("json", jason)

	m.Post("email", getEmail)

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

func getEmail(ctx *macaron.Context) []byte {
	data := ctx.Req.Body()

	var a Email

	d, _ := data.Bytes()

	err := json.Unmarshal(d, &a)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a.Assunto)
	fmt.Println(a.Mensagem)
	return d
}
