package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jayphan14/GoDatingApp/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:    util.RandName(),
		Email:       util.RandEmail(),
		Password:    util.RandPassword(),
		Gender:      util.RandGender(),
		University:  util.RandUniversity(),
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
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Gender, user2.Gender)
	require.Equal(t, user1.University, user2.University)
	require.Equal(t, user1.Picture, user2.Picture)
	require.Equal(t, user1.Bio, user2.Bio)
	require.Equal(t, user1.BioPictures, user2.BioPictures)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)
	updateArg := UpdateUserParams{
		ID:          user1.ID,
		Username:    util.RandName(),
		Email:       util.RandEmail(),
		Password:    util.RandPassword(),
		Gender:      util.RandGender(),
		University:  util.RandUniversity(),
		Picture:     []byte("base64_encoded_image_data"),
		Bio:         "A short bio",
		BioPictures: []string{"bio_picture_1.jpg", "bio_picture_2.jpg"},
	}
	user2, err := testQueries.UpdateUser(context.Background(), updateArg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, updateArg.ID, user2.ID)
	require.Equal(t, updateArg.Username, user2.Username)
	require.Equal(t, updateArg.Email, user2.Email)
	require.Equal(t, updateArg.Password, user2.Password)
	require.Equal(t, updateArg.Gender, user2.Gender)
	require.Equal(t, updateArg.University, user2.University)
	require.Equal(t, updateArg.Picture, user2.Picture)
	require.Equal(t, updateArg.Bio, user2.Bio)
	require.Equal(t, updateArg.BioPictures, user2.BioPictures)
}

func TestDeleteUser(t *testing.T) {
	account1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), account1.ID)
	require.NoError(t, err)
	account2, err := testQueries.GetUser(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

}
