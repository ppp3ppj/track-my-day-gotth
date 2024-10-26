package main

import (
	"fmt"
	"log"

	"github.com/ppp3ppj/track-my-day-gotth/internal/config"
	"github.com/ppp3ppj/track-my-day-gotth/internal/infrastructure/database"
)

func main() {
    cfg := config.MustLoadConfig()
    fmt.Printf("%+v\n", cfg.AppInfo)
    fmt.Printf("%+v\n", cfg.Server)
    fmt.Printf("%+v\n", cfg.Database)
    fmt.Printf("%+v\n", cfg.Database.CurrentDBPath)

    db := database.NewSqliteDatabase(cfg.Database)

    defer func() {
        if err := db.Close(); err != nil {
            log.Fatalf("Failed to close database connection: %v", err)
        }
    }()
}
