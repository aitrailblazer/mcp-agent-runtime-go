// mcp-agent-runtime-go: Starter Scaffold (net/http native)

// Directory: cmd/server/main.go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aitrailblazer/mcp-agent-runtime-go/internal/config"
	"github.com/aitrailblazer/mcp-agent-runtime-go/internal/router"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", router.HealthHandler)
	mux.HandleFunc("/ping", router.PingHandler)
	mux.HandleFunc("/invoke", router.InvokeHandler)
	mux.HandleFunc("/schema", router.SchemaHandler)

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	go func() {
		log.Printf("🚀 MCP Agent Runtime starting on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ Server failed: %s", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("🔻 Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("❌ Server Shutdown Failed: %+v", err)
	}
	log.Println("✅ Server exited cleanly")
}
