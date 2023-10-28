package main

import (
	"log/slog"

	"github.com/austien/type-of-the-bored/http"
)

func main() {
	const addr = ":8080"
	slog.Info("Starting server", "addr", addr)
	if err := http.ListenAndServe(addr); err != nil {
		slog.Error("server failed", "addr", addr, "err", err)
	}
}
