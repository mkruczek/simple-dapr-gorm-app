package common

import (
	"context"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const DefaultDsn = "default"

func InitializeDatabase(dsn string) *gorm.DB {

	var err error

	if dsn == "" || dsn == DefaultDsn {
		if dsn, err = testPostgresDSN(); err != nil {
			log.Fatal(err)
		}
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Nie można połączyć się z bazą danych")
	}
	db.AutoMigrate(&Product{})
	return db
}

func testPostgresDSN() (string, error) {
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
		return "", err
	}

	host, _ := postgresContainer.Host(ctx)
	port, _ := postgresContainer.MappedPort(ctx, "5432")

	return "postgres://postgres:secret@" + host + ":" + port.Port() + "?sslmode=disable", nil
}
