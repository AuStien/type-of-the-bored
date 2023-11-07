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
	"github.com/austien/type-of-the-bored/rooms"
)

//go:embed assets
var assets embed.FS

func ListenAndServe(addr string, roomClient rooms.RoomClient) error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Handle("/", templ.Handler(components.Page(components.Text())))
	r.HandleFunc("/text", v1.TextComponentGet())
	r.Route("/room", func(r chi.Router) {
		r.Route("/{roomID}", func(r chi.Router) {
			r.Get("/", v1.RoomComponentGet(roomClient))
			r.Post("/", v1.RoomComponentPost(roomClient))
			r.Get("/start", v1.RoomComponentStart(roomClient))
		})
	})

	r.Route("/v1", func(r chi.Router) {
		r.Post("/room", v1.CreateRoom(roomClient))
		r.Get("/rooms", v1.GetRooms(roomClient))
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
