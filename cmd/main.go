package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AnkitBishen/fileManagerApp/internal/crosMid"
	getFile "github.com/AnkitBishen/fileManagerApp/internal/handlers"
)

func main() {

	// get config

	// manage routes
	router := http.NewServeMux()
	router.HandleFunc("POST /api/getlist", getFile.List())
	router.HandleFunc("POST /api/createDir", getFile.CreateDir())
	router.HandleFunc("POST /api/delete", getFile.Delete())
	router.HandleFunc("POST /api/rename", getFile.Rename())
	router.HandleFunc("POST /api/read", getFile.Read())

	// handle cros errors
	crosMux := crosMid.CorsMiddleware(router)

	// start server in go routine ----
	server := http.Server{
		Addr:    ":8080",
		Handler: crosMux,
	}

	slog.Info("running server", slog.String("address", server.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shut down server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully...")

}
