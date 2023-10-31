package http

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/austien/type-of-the-bored/components"
	v1 "github.com/austien/type-of-the-bored/http/v1"
)

//go:embed assets
var assets embed.FS

func ListenAndServe(addr string) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/", templ.Handler(components.Page(components.Text())))
	r.HandleFunc("/text", v1.TextComponentGet())
	r.Route("/room", func(r chi.Router) {
		r.Get("/{roomID}", v1.RoomComponentGet())
		r.Post("/{roomID}", v1.RoomComponentPost())
	})

	r.Route("/v1", func(r chi.Router) {
		r.Post("/room", v1.CreateRoom())
		r.Get("/rooms", v1.GetRooms())
	})

	// Serve static assets
	r.Handle("/assets/*", http.FileServer(http.FS(assets)))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		msg := fmt.Sprintf("page %q doesn't exist, yo", r.URL.Path)

		var c templ.Component
		if r.Header.Get("HX-Request") == "true" {
			c = components.Error(msg)
		} else {
			c = components.Page(components.Error(msg))
		}
		if err := c.Render(r.Context(), w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return http.ListenAndServe(addr, r)
}
