package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/austien/type-of-the-bored/http"
	"github.com/austien/type-of-the-bored/inmem"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	roomClient := inmem.NewClient()

	addr := fmt.Sprintf(":%s", port)

	slog.Info("Starting server", "addr", addr)
	if err := http.ListenAndServe(addr, roomClient); err != nil {
		slog.Error("server failed", "addr", addr, "err", err)
	}
}
