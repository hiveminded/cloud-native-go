package main

import (
	"os"

	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"

	"github.com/hiveminded/cloud-native-go/api"
)

func main() {
	app := ion.New()

	app.Get("/", index)

	apiGroup := app.Party("/api")
	{
		apiGroup.Get("/echo", api.EchoHandler)
		apiGroup.Get("/hello", api.HelloHandler)
		apiGroup.Get("/books", api.AllBooksHandler)
		apiGroup.Get("/books/{isbn:string}", api.GetBookHandler)
		apiGroup.Post("/books", api.CreateBookHandler)
		apiGroup.Put("/books/{isbn:string}", api.UpdateBookHandler)
		apiGroup.Delete("/books/{isbn:string}", api.DeleteBookHandler)

	}

	app.Run(ion.Addr(":" + port()))
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(ctx context.Context) {
	ctx.Writef("Welcome to Cloud Native")
}
