// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Playlist struct {
	UserUuid     uuid.UUID          `json:"user_uuid"`
	PlaylistUuid uuid.UUID          `json:"playlist_uuid"`
	PlaylistID   string             `json:"playlist_id"`
	Name         string             `json:"name"`
	PlaylistCode string             `json:"playlist_code"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
}

type Song struct {
	SongUri      string    `json:"song_uri"`
	PlaylistUuid uuid.UUID `json:"playlist_uuid"`
	Count        int32     `json:"count"`
}

type Token struct {
	UserUuid uuid.UUID          `json:"user_uuid"`
	Refresh  []byte             `json:"refresh"`
	Access   []byte             `json:"access"`
	Expiry   pgtype.Timestamptz `json:"expiry"`
}

type User struct {
	ID           int64              `json:"id"`
	UserUuid     uuid.UUID          `json:"user_uuid"`
	SpotifyID    string             `json:"spotify_id"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
	Name         string             `json:"name"`
	Email        interface{}        `json:"email"`
	PasswordHash []byte             `json:"password_hash"`
	Activated    bool               `json:"activated"`
	Version      int32              `json:"version"`
}
