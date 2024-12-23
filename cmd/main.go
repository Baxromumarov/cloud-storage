package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/baxromumarov/cloud-storage/config"
	"github.com/baxromumarov/cloud-storage/internal/api"
	"github.com/baxromumarov/cloud-storage/internal/pkg/logger"
	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "cloud-storage")
	// postgres, err := helper.InitPostgres(cfg)
	// if err != nil {
	// log.Fatal("error while connecting postgres: " + err.Error())
	// }
	err := log.Error("Error while inserting file into database", logger.Error( fmt.Errorf("error-12")))
	fmt.Println(err)
	return
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	srv := api.New(&api.RouterOptions{
		Log: log,
		Cfg: cfg,
		Db:  &sqlx.DB{},
	})

	httpServer := &http.Server{
		Addr:    ":" + cfg.HttpPort,
		Handler: srv,
	}

	go func() {
		fmt.Println("Starting server on :8080")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %s\n", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	<-signalChan
	fmt.Println("Received shutdown signal, shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown failed: %v\n", err)
	} else {
		fmt.Println("Server gracefully stopped")
	}
}
