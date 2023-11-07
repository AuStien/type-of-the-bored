package inmem

import (
	"fmt"
)

type ErrRoomAlreadyStarted struct {
	RoomID string
}

func (e ErrRoomAlreadyStarted) Error() string {
	return fmt.Sprintf("room %q already started", e.RoomID)
}

type ErrRoomAlreadyFinished struct {
	RoomID string
}

func (e ErrRoomAlreadyFinished) Error() string {
	return fmt.Sprintf("room %q already finished", e.RoomID)
}

type ErrPlayerAlreadyExists struct {
	Nickname string
}

func (e ErrPlayerAlreadyExists) Error() string {
	return fmt.Sprintf("player %q already exists", e.Nickname)
}

type ErrNicknameNotFound struct {
	Nickname string
}

func (e ErrNicknameNotFound) Error() string {
	return fmt.Sprintf("nickname %q not found", e.Nickname)
}

type ErrRoomIDNotFound struct {
	RoomID string
}

func (e ErrRoomIDNotFound) Error() string {
	return fmt.Sprintf("room with ID %q not found", e.RoomID)
}
