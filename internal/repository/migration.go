package repository

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/khussa1n/shop/internal/config"
	"log"
)

type Migrate struct {
	Config  *config.Config
	Migrate *migrate.Migrate
}

func NewMigrate(config *config.Config) *Migrate {
	m := new(Migrate)

	log.Println("MigrationPath: ", config.DB.MigrationPath)
	migr, err := migrate.New(
		fmt.Sprintf("file://%s", config.DB.MigrationPath),
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.DBName))
	if err != nil {
		log.Fatal(err)
	}

	m.Migrate = migr
	m.Config = config

	return m
}

func (m *Migrate) Up() error {
	if err := m.Migrate.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			return fmt.Errorf("migrate up err: %w", err)
		}
	}

	return nil
}
