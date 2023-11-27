package main

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"os"
	"os/signal"
)

const (
	configPath = "../port.config"
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
	mustSavePortToSharedFile(port.Port())
	defer os.Remove(configPath)

	dsn := fmt.Sprintf(defaultDsn, port.Port())
	fmt.Printf("DSN: %s\n", dsn)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	fmt.Println("awaiting signal...")
	s := <-c
	fmt.Println("Got signal:", s)
}

// func will save port to file where other application cane read it
func mustSavePortToSharedFile(port string) {
	file, err := os.Create(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(port)
	if err != nil {
		log.Fatal(err)
	}
}
