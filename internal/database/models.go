// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Event struct {
	UserUuid  uuid.UUID          `json:"user_uuid"`
	EventUuid uuid.UUID          `json:"event_uuid"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	Name      string             `json:"name"`
	EventCode string             `json:"event_code"`
}

type Playlist struct {
	EventUuid  uuid.UUID          `json:"event_uuid"`
	PlaylistID string             `json:"playlist_id"`
	Name       string             `json:"name"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
}

type Song struct {
	SongUri    string `json:"song_uri"`
	PlaylistID string `json:"playlist_id"`
	Count      int32  `json:"count"`
}

type Token struct {
	Hash     []byte             `json:"hash"`
	UserUuid uuid.UUID          `json:"user_uuid"`
	Expiry   pgtype.Timestamptz `json:"expiry"`
	Scope    string             `json:"scope"`
}

type User struct {
	ID           int64              `json:"id"`
	UserUuid     uuid.UUID          `json:"user_uuid"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
	Name         string             `json:"name"`
	Email        interface{}        `json:"email"`
	PasswordHash []byte             `json:"password_hash"`
	Activated    bool               `json:"activated"`
	Version      int32              `json:"version"`
}