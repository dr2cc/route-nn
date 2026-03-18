package main

import (
	"fmt"
	"log"
	"os"
	"route-nn/internal/app"
	"route-nn/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	// Env
	// Пытаемся загрузить .env, если он есть.
	// Ошибку игнорируем, так как CONFIG_PATH уже может быть в системе!
	_ = godotenv.Load()

	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}

	// Run
	if err := app.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
