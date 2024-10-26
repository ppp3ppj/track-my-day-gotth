package main

import (
	"fmt"

	"github.com/ppp3ppj/track-my-day-gotth/internal/config"
)

func main() {
    cfg := config.MustLoadConfig()
    fmt.Printf("%+v\n", cfg.AppInfo)
    fmt.Printf("%+v\n", cfg.Server)
    fmt.Printf("%+v\n", cfg.Database)
    fmt.Printf("%+v\n", cfg.Database.CurrentDBPath)
}
