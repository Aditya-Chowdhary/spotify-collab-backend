// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: events.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO events (user_uuid, name, event_code)
VALUES ($1, $2, $3)
RETURNING user_uuid, event_uuid, created_at, updated_at, name, event_code
`

type CreateEventParams struct {
	UserUuid  uuid.UUID `json:"user_uuid"`
	Name      string    `json:"name"`
	EventCode string    `json:"event_code"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error) {
	row := q.db.QueryRow(ctx, createEvent, arg.UserUuid, arg.Name, arg.EventCode)
	var i Event
	err := row.Scan(
		&i.UserUuid,
		&i.EventUuid,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.EventCode,
	)
	return i, err
}

const deleteEvent = `-- name: DeleteEvent :exec
DELETE FROM events
WHERE user_uuid = $1 AND event_uuid = $2
`

type DeleteEventParams struct {
	UserUuid  uuid.UUID `json:"user_uuid"`
	EventUuid uuid.UUID `json:"event_uuid"`
}

func (q *Queries) DeleteEvent(ctx context.Context, arg DeleteEventParams) error {
	_, err := q.db.Exec(ctx, deleteEvent, arg.UserUuid, arg.EventUuid)
	return err
}

const getAllEvents = `-- name: GetAllEvents :many
SELECT user_uuid, event_uuid, created_at, updated_at, name, event_code
FROM events
WHERE user_uuid = $1
`

func (q *Queries) GetAllEvents(ctx context.Context, userUuid uuid.UUID) ([]Event, error) {
	rows, err := q.db.Query(ctx, getAllEvents, userUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.UserUuid,
			&i.EventUuid,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.EventCode,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getEvent = `-- name: GetEvent :one
SELECT user_uuid, event_uuid, created_at, updated_at, name, event_code
FROM events
WHERE user_uuid = $1 AND event_uuid = $2
`

type GetEventParams struct {
	UserUuid  uuid.UUID `json:"user_uuid"`
	EventUuid uuid.UUID `json:"event_uuid"`
}

func (q *Queries) GetEvent(ctx context.Context, arg GetEventParams) (Event, error) {
	row := q.db.QueryRow(ctx, getEvent, arg.UserUuid, arg.EventUuid)
	var i Event
	err := row.Scan(
		&i.UserUuid,
		&i.EventUuid,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.EventCode,
	)
	return i, err
}

const getEventUUIDByCode = `-- name: GetEventUUIDByCode :one
Select event_uuid
FROM events
WHERE event_code = $1
`

func (q *Queries) GetEventUUIDByCode(ctx context.Context, eventCode string) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, getEventUUIDByCode, eventCode)
	var event_uuid uuid.UUID
	err := row.Scan(&event_uuid)
	return event_uuid, err
}

const getEventUUIDByName = `-- name: GetEventUUIDByName :one
SELECT event_uuid
FROM events
WHERE name = $1
`

func (q *Queries) GetEventUUIDByName(ctx context.Context, name string) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, getEventUUIDByName, name)
	var event_uuid uuid.UUID
	err := row.Scan(&event_uuid)
	return event_uuid, err
}

const updateEventName = `-- name: UpdateEventName :one
UPDATE events
SET name = $1
WHERE user_uuid = $2 AND event_uuid = $3
RETURNING user_uuid, event_uuid, created_at, updated_at, name, event_code
`

type UpdateEventNameParams struct {
	Name      string    `json:"name"`
	UserUuid  uuid.UUID `json:"user_uuid"`
	EventUuid uuid.UUID `json:"event_uuid"`
}

func (q *Queries) UpdateEventName(ctx context.Context, arg UpdateEventNameParams) (Event, error) {
	row := q.db.QueryRow(ctx, updateEventName, arg.Name, arg.UserUuid, arg.EventUuid)
	var i Event
	err := row.Scan(
		&i.UserUuid,
		&i.EventUuid,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.EventCode,
	)
	return i, err
}
