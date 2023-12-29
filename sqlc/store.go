package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	*Queries
	connPool *pgxpool.Pool
}

func NewStore(connPool *pgxpool.Pool) *Store {
	return &Store{
		Queries:  New(connPool),
		connPool: connPool,
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}

type LikeTxParams struct {
	SenderID   pgtype.UUID `json:"sender_id"`
	ReceiverID pgtype.UUID `json:"receiver_id"`
}
type LikeTxResult struct {
	Like  Like  `json:"like"`
	Match Match `json:"match"`
}

// LikeTx perform a like on one account to another
// check if like already exist, create a like record, check if there is a match, then create a match record if there is
func (store *Store) LikeTx(ctx context.Context, arg LikeTxParams) (LikeTxResult, error) {
	var result LikeTxResult
	// call the generic execTx
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// Check if the like already exist
		_, errQueryingExistingLike := q.GetLikeByUsers(ctx, GetLikeByUsersParams{
			SenderID:   arg.SenderID,
			ReceiverID: arg.ReceiverID,
		})

		if errQueryingExistingLike == nil {
			return nil // the like already exist, do nothing
		}

		// Create a new Like entry
		result.Like, err = q.CreateLike(ctx, CreateLikeParams{
			SenderID:   arg.SenderID,
			ReceiverID: arg.ReceiverID,
		})
		if err != nil {
			return err
		}

		// Check if the other person also like you
		_, errQueryingOtherLike := q.GetLikeByUsers(ctx, GetLikeByUsersParams{
			SenderID:   arg.ReceiverID,
			ReceiverID: arg.SenderID,
		})

		if errQueryingOtherLike != nil {
			return nil
		}

		// They both like each other, create new match
		var errCreatingNewMatch error
		result.Match, errCreatingNewMatch = q.CreateMatch(ctx, CreateMatchParams{
			User1id: arg.SenderID,
			User2id: arg.ReceiverID,
		})

		if errCreatingNewMatch != nil {
			return errCreatingNewMatch
		}

		return nil
	})

	return result, err
}
