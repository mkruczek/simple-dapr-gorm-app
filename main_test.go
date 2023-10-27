package main

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"testing"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGORMWithPostgresContainer(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:13",
		ExposedPorts: []string{"5432/tcp"},
		Env:          map[string]string{"POSTGRES_PASSWORD": "secret"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
	}

	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("could not start container: %s", err)
	}
	defer postgresContainer.Terminate(ctx)

	host, _ := postgresContainer.Host(ctx)
	port, _ := postgresContainer.MappedPort(ctx, "5432")
	dsn := "postgres://postgres:secret@" + host + ":" + port.Port() + "?sslmode=disable"

	// Użyj GORM do połączenia z bazą danych
	db := InitializeDatabase(dsn)

	// Wykonaj operacje CRUD za pomocą GORM
	// Na przykład:
	CreateProduct(db, "ABC123", 100)
	product, err := GetProduct(db, 1)
	if err != nil || product.Code != "ABC123" {
		t.Fatalf("unexpected result from GetProduct: %v, error: %v", product, err)
	}
}
