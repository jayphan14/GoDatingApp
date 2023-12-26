package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username:    "john_doe",
		Email:       "john@example.com",
		Password:    "secret123",
		Gender:      "male",
		University:  "Example University",
		Picture:     []byte("base64_encoded_image_data"),
		Bio:         "A short bio about John Doe",
		BioPictures: []string{"bio_picture_1.jpg", "bio_picture_2.jpg"},
	}

	newAccount, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, newAccount)

	require.Equal(t, arg.Username, newAccount.Username)
	require.Equal(t, arg.Email, newAccount.Email)
}
