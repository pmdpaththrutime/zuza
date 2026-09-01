package main

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Username string
	Nickname string
}

type RecordingRoom struct {
	ID         uuid.UUID
	RoomCode   string
	Owner      User
	Performers []User // the users authorized to access this room
}

type Recording struct {
	ID         uuid.UUID
	Room       RecordingRoom
	Performers []User
	Tracks     []Track
}

type Track struct {
	ID        uuid.UUID
	Performer User
	File      []AudioFile // one or more (but usually one) files that represent this track
}

type AudioFile struct {
	ID       uuid.UUID
	MimeType string
	Location string
}
