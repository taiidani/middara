package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	redis "github.com/redis/go-redis/v9"
	"github.com/taiidani/middara/internal/cache"
	"github.com/taiidani/middara/internal/models"
	"github.com/taiidani/middara/internal/server"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	slog.SetLogLoggerLevel(slog.LevelDebug)

	// Set up the caching backend
	rds := cache.NewClient(ctx)

	// Set up the relational database
	if err := models.InitDB(ctx); err != nil {
		log.Fatalf("database init: %s", err)
	}

	// Serve until interrupted
	if err := serve(ctx, rds); err != nil {
		log.Fatal(err)
	}
}

func serve(ctx context.Context, rds *redis.Client) error {
	srv := server.NewServer(ctx, rds)

	go func() {
		slog.Info("Server starting", "dev", server.DevMode)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			slog.Error("Unclean server shutdown encountered", "error", err)
		}
	}()

	<-ctx.Done()

	// Gracefully shut down over 60 seconds
	slog.Info("Server shutting down")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Minute)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return err
	}

	slog.Info("Server exited")
	return nil
}
