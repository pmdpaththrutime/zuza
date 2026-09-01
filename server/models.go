package main

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username string
	Nickname string
}

type RecordingRoom struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	RoomCode   string    `gorm:"uniqueIndex"`
	OwnerID    uuid.UUID `gorm:"type:uuid;index"`
	Owner      User      `gorm:"foreignKey:OwnerID"`
	Performers []User    `gorm:"many2many:recording_room_performers"`
}

type Recording struct {
	ID         uuid.UUID     `gorm:"type:uuid;primaryKey"`
	RoomID     uuid.UUID     `gorm:"type:uuid;index"`
	Room       RecordingRoom `gorm:"foreignKey:RoomID"`
	Performers []User        `gorm:"many2many:recording_performers"`
	Tracks     []Track       `gorm:"foreignKey:RecordingID"`
}

type Track struct {
	ID          uuid.UUID   `gorm:"type:uuid;primaryKey"`
	RecordingID uuid.UUID   `gorm:"type:uuid;index"`
	PerformerID uuid.UUID   `gorm:"type:uuid;index"`
	Performer   User        `gorm:"foreignKey:PerformerID"`
	File        []AudioFile `gorm:"foreignKey:TrackID"`
}

type AudioFile struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	TrackID  uuid.UUID `gorm:"type:uuid;index"`
	MimeType string
	Location string
}
