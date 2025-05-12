package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/theCodeBoy12/oktaWebhook/internal/server/handlers"
	"github.com/theCodeBoy12/oktaWebhook/internal/server/middlewere"
)

var logger *slog.Logger

var (
	port         string = os.Getenv("PORT")
	token        string = os.Getenv("TOKEN")
	ConfFilePath string = os.Getenv("CONF_FILE_PATH")
)

func init() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	slog.SetDefault(logger)
	if port == "" {
		slog.Error("PORT is not set")
		os.Exit(1)
	}
	if token == "" {
		slog.Error("TOKEN is not set")
		os.Exit(1)
	}
	if ConfFilePath == "" {
		slog.Error("CONF_FILE_PATH is not set")
		os.Exit(1)
	}

}

func main() {
	//TODO: ratelimiting
	router := http.NewServeMux()
	router.HandleFunc("GET /{$}", handlers.VerificationHandler)
	router.Handle("POST /{$}", handlers.OktaWebhookRouter(ConfFilePath))
	authMiddlWere := middlewere.AuthMiddleware(token)

	srv := &http.Server{
		Addr: ":" + port,
		Handler: middlewere.LoggingMiddleware(
			authMiddlWere(router),
		),
		MaxHeaderBytes: 1 << 20,
	}

	slog.Info(fmt.Sprintf("Starting webhook on port %s", port))
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("Failed to start webhook server", "error", err)
		os.Exit(1)
	}

}
