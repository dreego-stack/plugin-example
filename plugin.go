package example

import (
	"fmt"
	"net/http"

	dreego "github.com/dreego-stack/dreego/core"
)

type Options struct {
	Prefix       string
	Greeting     string
	EnableLogging bool
}

func Register(app *dreego.App, options Options) error {
	if options.Prefix == "" {
		options.Prefix = "/example"
	}
	if options.Greeting == "" {
		options.Greeting = "Hello"
	}
	if options.EnableLogging {
		if err := app.Use(dreego.RequestLogging()); err != nil {
			return fmt.Errorf("example: register logging middleware: %w", err)
		}
	}
	if err := app.SetErrorHandler(http.StatusNotFound, notFoundHandler); err != nil {
		return fmt.Errorf("example: set error handler: %w", err)
	}
	if err := app.Register(http.MethodGet, options.Prefix+"/greet/{name}", greetHandler(options.Greeting)); err != nil {
		return fmt.Errorf("example: register greet route: %w", err)
	}
	if err := app.Register(http.MethodGet, options.Prefix+"/health", healthHandler); err != nil {
		return fmt.Errorf("example: register health route: %w", err)
	}
	return app.Register(http.MethodGet, "/{path...}", notFoundHandler)
}

func greetHandler(greeting string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "%s, %s!", greeting, name)
	}
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("ok"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "not found: %s", r.URL.Path)
}