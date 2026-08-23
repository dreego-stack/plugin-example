package example

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	dreego "github.com/dreego-stack/dreego/core"
)

func get(t *testing.T, srv *httptest.Server, path string) (int, string) {
	t.Helper()
	resp, err := http.Get(srv.URL + path)
	if err != nil {
		t.Fatalf("GET %s: %v", path, err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(body)
}

func TestRegisterDefaultOptions(t *testing.T) {
	app := dreego.New()
	if err := Register(app, Options{}); err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(app.Handler())
	defer srv.Close()

	if code, body := get(t, srv, "/example/greet/World"); code != 200 || body != "Hello, World!" {
		t.Fatalf("greet: code=%d body=%q want 200/Hello, World!", code, body)
	}
	if code, body := get(t, srv, "/example/health"); code != 200 || body != "ok" {
		t.Fatalf("health: code=%d body=%q want 200/ok", code, body)
	}
}

func TestRegisterCustomOptions(t *testing.T) {
	app := dreego.New()
	if err := Register(app, Options{Prefix: "/api", Greeting: "Hi"}); err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(app.Handler())
	defer srv.Close()

	if code, body := get(t, srv, "/api/greet/Ada"); code != 200 || body != "Hi, Ada!" {
		t.Fatalf("greet: code=%d body=%q want 200/Hi, Ada!", code, body)
	}
}

func TestRegisterEnableLogging(t *testing.T) {
	app := dreego.New()
	if err := Register(app, Options{EnableLogging: true}); err != nil {
		t.Fatalf("register with logging: %v", err)
	}
	srv := httptest.NewServer(app.Handler())
	defer srv.Close()
	if _, body := get(t, srv, "/example/health"); body != "ok" {
		t.Fatalf("health body=%q want ok", body)
	}
}

func TestRegisterAfterBuild(t *testing.T) {
	app := dreego.New()
	app.Build()
	if err := Register(app, Options{}); !errors.Is(err, dreego.ErrAppBuilt) {
		t.Fatalf("register after build error = %v, want ErrAppBuilt", err)
	}
}

func TestNotFoundHandler(t *testing.T) {
	app := dreego.New()
	if err := Register(app, Options{}); err != nil {
		t.Fatal(err)
	}
	srv := httptest.NewServer(app.Handler())
	defer srv.Close()

	if code, body := get(t, srv, "/nonexistent"); code != 404 || !strings.Contains(body, "not found:") {
		t.Fatalf("404: code=%d body=%q want 404 containing 'not found:'", code, body)
	}
}

func TestDuplicateRoute(t *testing.T) {
	app := dreego.New()
	if err := Register(app, Options{}); err != nil {
		t.Fatal(err)
	}
	if err := Register(app, Options{}); !errors.Is(err, dreego.ErrRouteConflict) {
		t.Fatalf("duplicate register error = %v, want ErrRouteConflict", err)
	}
}