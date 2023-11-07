package inmem

import (
	"errors"

	"github.com/austien/type-of-the-bored/rooms"
	"github.com/austien/type-of-the-bored/words"
)

type client struct {
	Rooms []*rooms.Room
}

var _ rooms.RoomClient = &client{}

func NewClient() *client {
	return &client{
		Rooms: []*rooms.Room{},
	}
}

func (c *client) CreateRoom() (*rooms.Room, error) {
	const maxRetries = 100
	var id string
	isUnique := false
	for retries := 0; retries < maxRetries; retries += 1 {
		id = rooms.NewID()
		for _, r := range c.Rooms {
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

	texts := map[int]rooms.Text{}

	for i := 0; i < 5; i++ {
		text, err := words.NewText(false, -1)
		if err != nil {
			return nil, err
		}

		chars := []string{}
		for _, c := range text.Word {
			chars = append(chars, string(c))
		}

		texts[i] = rooms.Text{
			Text:       text.Word,
			Characters: chars,
		}
	}

	room := &rooms.Room{
		ID:          id,
		Players:     []rooms.Player{},
		Texts:       texts,
		HasStarted:  false,
		HasFinished: false,
	}

	c.Rooms = append(c.Rooms, room)

	return room, nil
}

func (c *client) StartRoom(roomID string) error {
	room, err := c.GetRoom(roomID)
	if err != nil {
		return err
	}

	if room.HasStarted {
		return ErrRoomAlreadyStarted{RoomID: roomID}
	}

	if room.HasFinished {
		return ErrRoomAlreadyFinished{RoomID: roomID}
	}

	room.HasStarted = true
	return nil
}

func (c *client) JoinRoom(roomID, nickname string) error {
	room, err := c.GetRoom(roomID)
	if err != nil {
		return err
	}

	for _, p := range room.Players {
		if p.Nickname == nickname {
			return ErrPlayerAlreadyExists{Nickname: nickname}
		}
	}

	player := rooms.Player{
		Nickname:    nickname,
		CurrentText: 0,
		Progress:    newProgress(room.Texts),
		IsDone:      false,
	}

	room.Players = append(room.Players, player)
	return nil
}

func (c *client) LeaveRoom(roomID, nickname string) error {
	room, err := c.GetRoom(roomID)
	if err != nil {
		return err
	}

	for i, p := range room.Players {
		if p.Nickname == nickname {
			room.Players = append(room.Players[:i], room.Players[i+1:]...)
			return nil
		}
	}

	return ErrNicknameNotFound{Nickname: nickname}
}

func (c *client) GetRooms() []*rooms.Room {
	return c.Rooms
}

func (c *client) GetRoom(roomID string) (*rooms.Room, error) {
	for _, r := range c.Rooms {
		if r.ID == roomID {
			return r, nil
		}
	}

	return nil, ErrRoomIDNotFound{RoomID: roomID}
}

func newProgress(texts map[int]rooms.Text) map[int]rooms.TextProgress {
	progresses := map[int]rooms.TextProgress{}

	for i, t := range texts {
		progress := rooms.TextProgress{}
		for _, c := range t.Characters {
			progress.Characters = append(progress.Characters, rooms.Character{
				Character: c,
				IsCorrect: false,
			})
		}
		progresses[i] = progress
	}

	return progresses
}
