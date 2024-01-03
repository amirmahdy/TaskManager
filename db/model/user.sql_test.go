package db

import (
	"context"
	"taskmanager/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func createFakeDBUser() (username string, params CreateUserParams, err error) {
	username = utils.CreateRandomName()
	email := utils.CreateRandomEmail()
	fullName := utils.CreateRandomName()
	password := utils.CreateRandomString(8)
	hash, _ := utils.CreateHashPassword(password)

	params = CreateUserParams{
		Username:       username,
		FullName:       fullName,
		Email:          email,
		HashedPassword: hash,
	}

	// Call the create user function
	_, err = testStore.CreateUser(context.Background(), params)
	return
}

func TestCreateAndGetUser(t *testing.T) {
	username, params, err := createFakeDBUser()
	require.NoError(t, err)

	// Call the get user function
	user, err := testStore.GetUser(context.Background(), username)
	require.NoError(t, err)

	// Verify the result
	require.Equal(t, params.Username, user.Username, "expected username %s, got %s", params.Username, user.Username)
	// Add more assertions as needed

}
