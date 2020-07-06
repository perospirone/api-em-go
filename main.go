package main

import (
	"net/http"

	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.Classic()
	macaron.Logger()
	routes(m)

	m.Run()
}

func routes(m *macaron.Macaron) {
	m.Get("/", handler)

	m.Post("/", func() string {
		return "dani"
	})

	m.Get("/da", func(resp http.ResponseWriter, req *http.Request) {
		resp.WriteHeader(300)
		
	})
}

func handler(ctx *macaron.Context) string {
	return ctx.Req.RequestURI
}
