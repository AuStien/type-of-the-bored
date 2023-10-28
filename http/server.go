package http

import (
	"embed"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/austien/type-of-the-bored/components"
	"github.com/austien/type-of-the-bored/words"
)

//go:embed assets
var assets embed.FS

func ListenAndServe(addr string) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/", templ.Handler(components.Page()))
	r.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		text, err := words.NewText(false, -1)
		if err != nil {
			slog.Error("failed to generate text", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var letters []string
		for _, l := range text.Word {
			letters = append(letters, string(l))
		}

		c := components.Text(letters, text.Description)
		if err := c.Render(r.Context(), w); err != nil {
			slog.Error("failed to render text", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Serve static assets
	r.Handle("/assets/*", http.FileServer(http.FS(assets)))

	return http.ListenAndServe(addr, r)
}
