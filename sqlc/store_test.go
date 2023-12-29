package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// Check that likes can be created an updated concurrently
func TestLikeTx(t *testing.T) {

	store := NewStore(testDB)
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	user4 := createRandomUser(t)

	users := []User{user1, user2, user3, user4}
	// Run n concurent opertion Like
	n := len(users)

	errs := make(chan error)
	results := make(chan LikeTxResult)

	for i := 0; i < n-1; i++ {
		go func(i int) {
			fmt.Println("creating a like for user", users[i].ID, users[i+1].ID)
			result, err := store.LikeTx(context.Background(), LikeTxParams{
				SenderID:   users[i].ID,
				ReceiverID: users[i+1].ID,
			})

			errs <- err
			results <- result
		}(i)
	}

	for i := 0; i < n-1; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results

		require.NotEmpty(t, result)

		// Check transfer
		newLike := result.Like
		require.NotEmpty(t, newLike)
	}

}

func TestMatch(t *testing.T) {
	fmt.Println("Test2 Ran")
	store := NewStore(testDB)
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)

	result1, err1 := store.LikeTx(context.Background(), LikeTxParams{
		SenderID:   user1.ID,
		ReceiverID: user2.ID,
	})

	require.NoError(t, err1)
	require.NotEmpty(t, result1)

	// Check transfer
	newLike1 := result1.Like
	newMatch1 := result1.Match
	require.NotEmpty(t, newLike1)
	require.Equal(t, newLike1.SenderID, user1.ID)
	require.Equal(t, newLike1.ReceiverID, user2.ID)
	require.Empty(t, newMatch1)

	result2, err2 := store.LikeTx(context.Background(), LikeTxParams{
		SenderID:   user2.ID,
		ReceiverID: user1.ID,
	})

	newLike2 := result2.Like
	newMatch2 := result2.Match

	require.NoError(t, err2)
	require.NotEmpty(t, newLike2)
	require.Equal(t, newLike2.SenderID, user2.ID)
	require.Equal(t, newLike2.ReceiverID, user1.ID)
	require.Equal(t, newMatch2.User1id, user2.ID)
	require.Equal(t, newMatch2.User2id, user1.ID)
}
