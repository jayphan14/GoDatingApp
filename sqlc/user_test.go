package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:    util.randName(),
		Email:       "john@example.com",
		Password:    "secret123",
		Gender:      "male",
		University:  "Example University",
		Picture:     []byte("base64_encoded_image_data"),
		Bio:         "A short bio",
		BioPictures: []string{"bio_picture_1.jpg", "bio_picture_2.jpg"},
	}

	newUser, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	require.Equal(t, arg.Username, newUser.Username)
	require.Equal(t, arg.Email, newUser.Email)
	return newUser

}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	// newUser := createRandomUser(t)

}
