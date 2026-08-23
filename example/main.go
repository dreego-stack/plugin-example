package main

import (
	"log"
	"os"

	dreego "github.com/dreego-stack/dreego/core"
	example "github.com/dreego-stack/plugin-example"
)

func main() {
	app := dreego.New()
	if err := example.Register(app, example.Options{
		Prefix:        "/api",
		Greeting:      "Hi",
		EnableLogging: true,
	}); err != nil {
		log.Fatal(err)
	}
	addr := ":8080"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}