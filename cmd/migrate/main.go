package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gmhafiz/go8/config"
	"github.com/gmhafiz/go8/database"
	_ "github.com/lib/pq"
)

// Version is injected using ldflags during build time
var Version string

func main() {

	log.Printf("Version: %s\n", Version)

	cfg := config.Load(".env")

	fmt.Println(cfg)

	dsn := fmt.Sprintf("%s://%s:%d/%s?sslmode=%s&user=%s&password=%s",
		cfg.Database.Driver,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.SslMode,
		cfg.Database.User,
		cfg.Database.Pass)

	migrator := database.Migrator(database.WithDSN(dsn))

	// sqlDatabase := db.New(cfg)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Migrate failed %v", err)
	}
	migrator.DB = db

	if err := migrator.DB.Ping(); err != nil {
		log.Fatalf("Migrate Ping error %v", err)
	}

	// todo: accept cli flag for other operations
	migrator.Up()
}
