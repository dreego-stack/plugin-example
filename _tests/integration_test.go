package tests

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	dreego "github.com/dreego-stack/dreego/core"
	example "github.com/dreego-stack/plugin-example"
)

func TestExamplePluginIntegration(t *testing.T) {
	app := dreego.New()
	if err := example.Register(app, example.Options{
		Prefix:        "/api",
		Greeting:      "Hi",
		EnableLogging: false,
	}); err != nil {
		t.Fatalf("Register: %v", err)
	}

	server := httptest.NewServer(app.Handler())
	defer server.Close()

	tests := []struct {
		path     string
		wantCode int
		wantBody string
	}{
		{"/api/greet/World", 200, "Hi, World!"},
		{"/api/greet/Ada", 200, "Hi, Ada!"},
		{"/api/health", 200, "ok"},
	}

	for _, tt := range tests {
		resp, err := http.Get(server.URL + tt.path)
		if err != nil {
			t.Fatalf("GET %s: %v", tt.path, err)
		}
		if resp.StatusCode != tt.wantCode {
			t.Errorf("GET %s: status = %d, want %d", tt.path, resp.StatusCode, tt.wantCode)
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		if !strings.Contains(string(body), tt.wantBody) {
			t.Errorf("GET %s: body = %q, want %q", tt.path, string(body), tt.wantBody)
		}
	}
}

func TestExamplePluginDefaultOptions(t *testing.T) {
	app := dreego.New()
	if err := example.Register(app, example.Options{}); err != nil {
		t.Fatalf("Register with defaults: %v", err)
	}
	server := httptest.NewServer(app.Handler())
	defer server.Close()

	resp, err := http.Get(server.URL + "/example/greet/Test")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	if !strings.Contains(string(body), "Hello, Test!") {
		t.Errorf("body = %q, want 'Hello, Test!'", string(body))
	}
}

func TestExamplePluginErrAppBuilt(t *testing.T) {
	app := dreego.New()
	_ = app.Handler()
	err := example.Register(app, example.Options{})
	if err == nil {
		t.Fatal("expected ErrAppBuilt after build")
	}
}