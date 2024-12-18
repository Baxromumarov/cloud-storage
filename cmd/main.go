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
	"github.com/baxromumarov/cloud-storage/internal/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "cloud-storage")
	fmt.Println(log)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Create a server instance with the Gin router
	srv := &http.Server{
		Addr:    ":8080", // Port to listen on
		Handler: r,       // Gin handler
	}

	// Run the server in a separate goroutine
	go func() {
		fmt.Println("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %s\n", err)
		}
	}()

	// Block until we receive an interrupt signal
	<-stop
	fmt.Println("\nShutting down server gracefully...")

	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shut down: %s\n", err)
	} else {
		fmt.Println("Server exited cleanly.")
	}
}
