package main

import (
	"context"
	"log"

	"github.com/Jonss/cartaman/pkg/adapters/repository/pg"
	"github.com/Jonss/cartaman/pkg/ports/httprest"
	"github.com/Jonss/cartaman/pkg/usecase/deck"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	var err error
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("error getting config: error=(%v)", err)
	}

	conn, err := pg.NewConnection(cfg.DBURL)
	if err != nil {
		log.Fatalf("error on connect=(%q)", err)
	}

	err = pg.Migrate(conn, cfg.DBName, cfg.DBMigrationPath)
	if err != nil {
		log.Fatalf("error on migrate=(%q)", err)
	}

	cardRepo := pg.NewPGCardRepository(conn)
	err = cardRepo.SeedCards(context.Background())
	if err != nil {
		log.Fatalf("error on seed cards=(%q)", err)
	}

	deckRepo := pg.NewPGDeckRepository(conn)
	deckService := deck.NewDeckService(deckRepo, cardRepo)

	r := httprest.NewApp(fiber.New(), &deckService)
	r.Routes()

	log.Fatal(r.FiberApp.Listen(":" + cfg.Port))
}

type Config struct {
	Port            string `mapstructure:"PORT"`
	DBURL           string `mapstructure:"DATABASE_URL"`
	DBName          string `mapstructure:"DATABASE_NAME"`
	DBMigrationPath string `mapstructure:"DATABASE_MIGRATION_PATH"`
}

func loadConfig() (Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, nil
		}
		return config, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}
	return config, nil
}
