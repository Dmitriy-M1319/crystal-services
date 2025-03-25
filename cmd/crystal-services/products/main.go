package main

import (
	"flag"
	"github.com/Dmitriy-M1319/crystal-services/internal/config"
	"github.com/Dmitriy-M1319/crystal-services/internal/products/db"
	"github.com/Dmitriy-M1319/crystal-services/internal/products/server"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := config.ReadConfigYML("config.yaml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}

	cfg := config.GetConfigInstance()

	migration := flag.Bool("migration", true, "Defines the migration start option")
	flag.Parse()

	var err error
	conn, err := db.NewConnection(cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init postgres")
	}
	defer db.Close(conn)

	if *migration {
		err = goose.Up(conn.DB, cfg.Database.Migrations)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to up migrations")
		}
	}

	if err := server.NewGrpcServer(conn).Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")
		return
	}

}
