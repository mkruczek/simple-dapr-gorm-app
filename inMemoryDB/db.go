package main

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
)

const (
	defaultDsn = "postgres://postgres:secret@localhost:%s?sslmode=disable"
)

func main() {

	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:13",
		ExposedPorts: []string{"5432/tcp"},
		Env:          map[string]string{"POSTGRES_PASSWORD": "secret"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
	}

	var err error
	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	defer postgresContainer.Terminate(ctx)
	if err != nil {
		log.Fatalf("can't run postgres container: %s", err)
	}

	port, _ := postgresContainer.MappedPort(ctx, "5432")

	fmt.Printf("PORT: %s\n", port)

	dsn := fmt.Sprintf(defaultDsn, port.Port())

	fmt.Printf("DSN: %s", dsn)

	for {
	}
}
