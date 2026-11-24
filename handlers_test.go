package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	rec := httptest.NewRecorder()
	helloHandler(Config{Env: "staging"})(rec, httptest.NewRequest(http.MethodGet, "/hello", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if got, want := rec.Body.String(), "Hello! I am running in staging\n"; got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestHealthz(t *testing.T) {
	rec := httptest.NewRecorder()
	healthzHandler(rec, httptest.NewRequest(http.MethodGet, "/healthz", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
}

func TestLoadConfigDefaults(t *testing.T) {
	t.Setenv("ENV", "")
	t.Setenv("DB_HOST", "db.local")
	t.Setenv("DB_PASSWORD", "secret")

	cfg, err := loadConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Env != "development" {
		t.Errorf("Env = %q, want %q", cfg.Env, "development")
	}
}

func TestLoadConfigRequiresDBSettings(t *testing.T) {
	cases := map[string]map[string]string{
		"missing host":     {"DB_HOST": "", "DB_PASSWORD": "secret"},
		"missing password": {"DB_HOST": "db.local", "DB_PASSWORD": ""},
	}

	for name, env := range cases {
		t.Run(name, func(t *testing.T) {
			for k, v := range env {
				t.Setenv(k, v)
			}
			if _, err := loadConfig(); err == nil {
				t.Fatal("expected an error, got nil")
			}
		})
	}
}
