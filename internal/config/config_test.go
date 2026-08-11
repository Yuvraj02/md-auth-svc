package config_test

import (
	"testing"

	"github.com/marketing-digest/auth-service/internal/config"
)

func TestLoadRequiresDatabaseEnv(t *testing.T) {
	t.Setenv("DATABASE_HOST", "")
	t.Setenv("DATABASE_NAME", "")
	t.Setenv("DATABASE_USER", "")
	t.Setenv("DATABASE_PASSWORD", "")

	_, err := config.Load()
	if err == nil {
		t.Fatal("expected error when database env is missing")
	}
}

func TestLoadOK(t *testing.T) {
	t.Setenv("DATABASE_HOST", "localhost")
	t.Setenv("DATABASE_PORT", "5432")
	t.Setenv("DATABASE_NAME", "marketing_digest_auth")
	t.Setenv("DATABASE_USER", "auth")
	t.Setenv("DATABASE_PASSWORD", "secret")
	t.Setenv("GRPC_PORT", "50051")

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Service != "auth-service" {
		t.Fatalf("service=%s", cfg.Service)
	}
	if cfg.DatabaseName != "marketing_digest_auth" {
		t.Fatalf("db=%s", cfg.DatabaseName)
	}
}
