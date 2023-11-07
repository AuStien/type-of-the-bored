package rooms

import (
	"errors"
)

var Rooms []*Room

var ErrPlayerAlreadyExists = errors.New("player already exists")

type RoomClient interface {
	CreateRoom() (*Room, error)
	StartRoom(roomID string) error
	JoinRoom(roomID, nickname string) error
	LeaveRoom(roomID, nickname string) error
	GetRooms() []*Room
	GetRoom(roomID string) (*Room, error)
}

type Room struct {
	ID          string // Combination of <adjective>-<noun>
	Players     []Player
	Texts       map[int]Text // Map from ID (0-X) to Text
	HasStarted  bool
	HasFinished bool
}

type Text struct {
	Text       string   // The full text
	Characters []string // Each individual character
}

type Player struct {
	Nickname    string // Chosen nickname
	CurrentText int    // Which texts the player is on
	Progress    map[int]TextProgress
	IsDone      bool
}

type TextProgress struct {
	Characters []Character
	IsDone     bool
}

type Character struct {
	Character string
	IsCorrect bool
}

// func NewRoom() (*Room, error) {
// 	const maxRetries = 100
// 	var id string
// 	isUnique := false
// 	for retries := 0; retries < maxRetries; retries += 1 {
// 		id = NewID()
// 		for _, r := range Rooms {
// 			if r.ID == id {
// 				continue
// 			}
// 		}
// 		isUnique = true
// 		break
// 	}

// 	if !isUnique {
// 		return nil, errors.New("failed to generate unique room ID")
// 	}

// 	texts := map[int]Text{}

// 	for i := 0; i < 5; i++ {
// 		text, err := words.NewText(false, -1)
// 		if err != nil {
// 			return nil, err
// 		}

// 		chars := []string{}
// 		for _, c := range text.Word {
// 			chars = append(chars, string(c))
// 		}

// 		texts[i] = Text{
// 			Text:       text.Word,
// 			Characters: chars,
// 		}
// 	}

// 	room := &Room{
// 		ID:          id,
// 		Players:     []Player{},
// 		Texts:       texts,
// 		HasStarted:  false,
// 		HasFinished: false,
// 	}

// 	Rooms = append(Rooms, room)

// 	return room, nil
// }

// func GetRoom(id string) (*Room, error) {
// 	for _, r := range Rooms {
// 		if r.ID == id {
// 			return r, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("room %q not found", id)
// }

// func AddPlayer(roomID, name string) error {
// 	room, err := GetRoom(roomID)
// 	if err != nil {
// 		return err
// 	}
// 	for _, p := range room.Players {
// 		if p.Nickname == name {
// 			return ErrPlayerAlreadyExists
// 		}
// 	}

// 	player := Player{
// 		Nickname:    name,
// 		CurrentText: 0,
// 		Progress:    NewProgress(room.Texts),
// 		IsDone:      false,
// 	}

// 	room.Players = append(room.Players, player)
// 	return nil
// }

// func NewProgress(texts map[int]Text) map[int]TextProgress {
// 	progresses := map[int]TextProgress{}

// 	for i, t := range texts {
// 		progress := TextProgress{}
// 		for _, c := range t.Characters {
// 			progress.Characters = append(progress.Characters, Character{
// 				Character: c,
// 				IsCorrect: false,
// 			})
// 		}
// 		progresses[i] = progress
// 	}

// 	return progresses
// }
