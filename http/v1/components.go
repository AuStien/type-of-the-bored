package v1

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"

	"github.com/austien/type-of-the-bored/components"
	"github.com/austien/type-of-the-bored/rooms"
	"github.com/austien/type-of-the-bored/words"
)

func TextComponentGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		c := components.TextBox(letters, text.Description)
		if err := c.Render(r.Context(), w); err != nil {
			slog.Error("failed to render text", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func RoomComponentGet(roomClient rooms.RoomClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isHTMX := r.Header.Get("HX-Request") == "true"
		roomID := chi.URLParam(r, "roomID")
		if roomID == "" {
			http.Error(w, "missing room ID", http.StatusBadRequest)
			return
		}

		if _, err := roomClient.GetRoom(roomID); err != nil {
			component := pageOrChild(isHTMX, components.Error(err.Error()))
			if err := component.Render(r.Context(), w); err != nil {
				slog.Error("failed to render error", "err", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		component := pageOrChild(isHTMX, components.EnteringRoom(roomID))
		if err := component.Render(r.Context(), w); err != nil {
			slog.Error("failed to render room", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func RoomComponentPost(roomClient rooms.RoomClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isHTMX := r.Header.Get("HX-Request") == "true"
		roomID := chi.URLParam(r, "roomID")
		if roomID == "" {
			http.Error(w, "missing room ID", http.StatusBadRequest)
			return
		}

		room, err := roomClient.GetRoom(roomID)
		if err != nil {
			components := pageOrChild(isHTMX, components.Error(err.Error()))
			if err := components.Render(r.Context(), w); err != nil {
				slog.Error("failed to render error", "err", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		if err := r.ParseForm(); err != nil {
			slog.Error("failed to parse form", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		username := r.Form.Get("username")

		if username == "" {
			http.Error(w, "missing username", http.StatusBadRequest)
		}

		if err := roomClient.JoinRoom(roomID, username); err != nil {
			slog.Error("failed to add player", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		component := pageOrChild(isHTMX, components.RoomAsUser(roomID, username, room.HasStarted))
		if err := component.Render(r.Context(), w); err != nil {
			slog.Error("failed to render room", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func RoomComponentStart(roomClient rooms.RoomClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isHTMX := r.Header.Get("HX-Request") == "true"
		username := r.Header.Get("user")
		roomID := chi.URLParam(r, "roomID")
		if roomID == "" {
			http.Error(w, "missing room ID", http.StatusBadRequest)
			return
		}

		room, err := roomClient.GetRoom(roomID)
		if err != nil {
			components := pageOrChild(isHTMX, components.Error(err.Error()))
			if err := components.Render(r.Context(), w); err != nil {
				slog.Error("failed to render error", "err", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		if err := roomClient.StartRoom(roomID); err != nil {
			slog.Error("failed to start room", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		component := pageOrChild(isHTMX, components.RoomAsUser(roomID, username, room.HasStarted))
		if err := component.Render(r.Context(), w); err != nil {
			slog.Error("failed to render room", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func pageOrChild(isHTMX bool, child templ.Component) templ.Component {
	if isHTMX {
		return child
	}
	return components.Page(child)
}
