package v1

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/austien/type-of-the-bored/rooms"
)

func CreateRoom(roomClient rooms.RoomClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		room, err := roomClient.CreateRoom()
		if err != nil {
			slog.Error("failed to create room", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rooms.Rooms = append(rooms.Rooms, room)

		w.Header().Set("HX-Redirect", fmt.Sprintf("/room/%s", room.ID))
		w.WriteHeader(http.StatusCreated)
	}
}

func GetRooms(roomClient rooms.RoomClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type response struct {
			Rooms []*rooms.Room `json:"rooms"`
		}

		currRooms := roomClient.GetRooms()
		if currRooms == nil {
			currRooms = []*rooms.Room{}
		}

		if err := json.NewEncoder(w).Encode(response{Rooms: currRooms}); err != nil {
			slog.Error("failed to encode response", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
