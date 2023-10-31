package rooms

import (
	"errors"
	"fmt"

	"github.com/austien/type-of-the-bored/words"
)

type rooms []*Room

var Rooms rooms

var ErrPlayerAlreadyExists = errors.New("player already exists")

type Room struct {
	ID      string
	Players []Player
	Text    []words.Text
}

type Player struct {
	Name string
}

func NewRoom() (*Room, error) {
	const maxRetries = 100
	var id string
	isUnique := false
	for retries := 0; retries < maxRetries; retries += 1 {
		id = NewID()
		for _, r := range Rooms {
			if r.ID == id {
				continue
			}
		}
		isUnique = true
		break
	}

	if !isUnique {
		return nil, errors.New("failed to generate unique room ID")
	}

	return &Room{
		ID:      id,
		Players: []Player{},
		Text:    []words.Text{},
	}, nil
}

func GetRoom(id string) (*Room, error) {
	for _, r := range Rooms {
		if r.ID == id {
			return r, nil
		}
	}

	return nil, fmt.Errorf("room %q not found", id)
}

func (r rooms) AddPlayer(roomID, name string) error {
	room, err := GetRoom(roomID)
	if err != nil {
		return err
	}
	for _, p := range room.Players {
		if p.Name == name {
			return ErrPlayerAlreadyExists
		}
	}

	room.Players = append(room.Players, Player{Name: name})
	return nil
}
