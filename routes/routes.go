package routes

import (
	"encoding/json"
	"fmt"
	"github.com/danitw/api-em-go/controllers"
	"github.com/danitw/api-em-go/database"
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



func Routes(m *macaron.Macaron) {
	m.Get("json", jason)

	m.Post("email", getEmail)

	m.Post("register", register)
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

	fmt.Println(a)
	return d
}

func register(ctx *macaron.Context) string{
	data := ctx.Req.Body()

	d, _ := data.Bytes()

	bug := controllers.Register(database.Connection(), d)

	if bug == false {
		return "Deu errado"
	}

	return "Deu certo"
}