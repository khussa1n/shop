package app

import (
	"github.com/khussa1n/shop/internal/config"
	"github.com/khussa1n/shop/internal/handler"
	"github.com/khussa1n/shop/internal/repository"
	"github.com/khussa1n/shop/internal/service"
	"github.com/khussa1n/shop/pkg/client/postgres"
	"log"
)

func Run(cfg *config.Config) error {
	db, err := postgres.New(
		postgres.WithHost(cfg.DB.Host),
		postgres.WithPort(cfg.DB.Port),
		postgres.WithDBName(cfg.DB.DBName),
		postgres.WithUsername(cfg.DB.Username),
		postgres.WithPassword(cfg.DB.Password),
	)
	if err != nil {
		log.Printf("connection to DB err: %s", err.Error())
		return err
	}
	defer db.Close()
	log.Println("connection success")

	migration := repository.NewMigrate(cfg)
	err = migration.Up()
	if err != nil {
		log.Printf("from migration")
		return err
	}
	log.Println("migration success")

	repos := repository.NewRepository(db.Pool)
	services := service.NewService(repos)
	h := handler.NewHandler(services)
	h.InitHandler()

	return nil
}
