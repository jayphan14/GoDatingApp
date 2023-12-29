// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: like.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createLike = `-- name: CreateLike :one
INSERT INTO likes (
  sender_id, receiver_id
) VALUES (
  $1, $2
)
RETURNING id, sender_id, receiver_id, created_at
`

type CreateLikeParams struct {
	SenderID   pgtype.UUID `json:"sender_id"`
	ReceiverID pgtype.UUID `json:"receiver_id"`
}

func (q *Queries) CreateLike(ctx context.Context, arg CreateLikeParams) (Like, error) {
	row := q.db.QueryRow(ctx, createLike, arg.SenderID, arg.ReceiverID)
	var i Like
	err := row.Scan(
		&i.ID,
		&i.SenderID,
		&i.ReceiverID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteLike = `-- name: DeleteLike :exec
DELETE FROM likes
WHERE sender_id = $1 AND receiver_id = $2
`

type DeleteLikeParams struct {
	SenderID   pgtype.UUID `json:"sender_id"`
	ReceiverID pgtype.UUID `json:"receiver_id"`
}

func (q *Queries) DeleteLike(ctx context.Context, arg DeleteLikeParams) error {
	_, err := q.db.Exec(ctx, deleteLike, arg.SenderID, arg.ReceiverID)
	return err
}

const getLike = `-- name: GetLike :one
SELECT id, sender_id, receiver_id, created_at FROM likes
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetLike(ctx context.Context, id pgtype.UUID) (Like, error) {
	row := q.db.QueryRow(ctx, getLike, id)
	var i Like
	err := row.Scan(
		&i.ID,
		&i.SenderID,
		&i.ReceiverID,
		&i.CreatedAt,
	)
	return i, err
}

const getLikeByUsers = `-- name: GetLikeByUsers :one
SELECT id, sender_id, receiver_id, created_at FROM likes
WHERE sender_id = $1 AND receiver_id = $2
`

type GetLikeByUsersParams struct {
	SenderID   pgtype.UUID `json:"sender_id"`
	ReceiverID pgtype.UUID `json:"receiver_id"`
}

func (q *Queries) GetLikeByUsers(ctx context.Context, arg GetLikeByUsersParams) (Like, error) {
	row := q.db.QueryRow(ctx, getLikeByUsers, arg.SenderID, arg.ReceiverID)
	var i Like
	err := row.Scan(
		&i.ID,
		&i.SenderID,
		&i.ReceiverID,
		&i.CreatedAt,
	)
	return i, err
}

const listLikesReceived = `-- name: ListLikesReceived :many
SELECT id, sender_id, receiver_id, created_at FROM likes
WHERE receiver_id = $1
`

func (q *Queries) ListLikesReceived(ctx context.Context, receiverID pgtype.UUID) ([]Like, error) {
	rows, err := q.db.Query(ctx, listLikesReceived, receiverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Like
	for rows.Next() {
		var i Like
		if err := rows.Scan(
			&i.ID,
			&i.SenderID,
			&i.ReceiverID,
			&i.CreatedAt,
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

const listLikesSent = `-- name: ListLikesSent :many
SELECT id, sender_id, receiver_id, created_at FROM likes
WHERE sender_id = $1
`

func (q *Queries) ListLikesSent(ctx context.Context, senderID pgtype.UUID) ([]Like, error) {
	rows, err := q.db.Query(ctx, listLikesSent, senderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Like
	for rows.Next() {
		var i Like
		if err := rows.Scan(
			&i.ID,
			&i.SenderID,
			&i.ReceiverID,
			&i.CreatedAt,
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
