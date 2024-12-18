package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/baxromumarov/cloud-storage/config"
	"github.com/baxromumarov/cloud-storage/internal/api"
	"github.com/baxromumarov/cloud-storage/internal/db"
	"github.com/baxromumarov/cloud-storage/internal/pkg/logger"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "cloud-storage")
	postgres, err := db.InitPostgres(cfg)
	if err != nil {
		log.Fatal("error while connecting postgres: " + err.Error())
	}


	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	srv := api.New(&api.RouterOptions{
		Log: log,
		Cfg: cfg,
		Db:  postgres,
	})

	fmt.Println("Starting server on :8080")
	if err := srv.Run(":" + cfg.HttpPort); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Server error: %s\n", err)
	}

}
