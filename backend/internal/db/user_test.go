package db

import (
	"context"
	"os"
	"testing"

	"github.com/LostProgrammer1010/InventorySystem/internal/authentication"
	"github.com/LostProgrammer1010/InventorySystem/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	connect()
	userCollection = client.Database("InventorySystem_test").Collection("User_test")
	createUniqueIndexes(userCollection, []string{"email", "username"})
	code := m.Run()
	userCollection.Drop(context.TODO())
	os.Exit(code)
}

func TestAddUser(t *testing.T) {
	test_user := models.User{
		Username:     "LostProgrammer1010",
		Password:     "123",
		FirstName:    "Dustin",
		LastName:     "Meyer",
		Email:        "check@gmail.com",
		Organization: []models.Organization{{Name: "NSVC", Role: "ADMIN"}},
	}
	err := AddUser(test_user)

	assert.NoError(t, err)

	found_user, err := GetUserByUsername(test_user.Username)

	assert.NoError(t, err)
	assert.Equal(t, test_user.Username, found_user.Username)
	assert.Equal(t, test_user.Email, found_user.Email)
	assert.True(t, authentication.VerifyPassword(found_user.Password, test_user.Password))
	assert.Equal(t, test_user.FirstName, found_user.FirstName)
	assert.Equal(t, test_user.LastName, found_user.LastName)
	assert.Equal(t, test_user.Organization[0].Name, found_user.Organization[0].Name)
	assert.Equal(t, test_user.Organization[0].Role, found_user.Organization[0].Role)

}

func TestAddSameUser(t *testing.T) {
	test_user := models.User{
		Username:     "LostProgrammer1010",
		Password:     "123",
		FirstName:    "Dustin",
		LastName:     "Meyer",
		Email:        "check@gmail.com",
		Organization: []models.Organization{{Name: "NSVC", Role: "ADMIN"}},
	}
	err := AddUser(test_user)
	assert.Error(t, err)

	found_user, err := GetUserByUsername(test_user.Username)
	assert.NoError(t, err)

	assert.Equal(t, test_user.Username, found_user.Username)
	assert.Equal(t, test_user.Email, found_user.Email)
	assert.True(t, authentication.VerifyPassword(found_user.Password, test_user.Password))
	assert.Equal(t, test_user.FirstName, found_user.FirstName)
	assert.Equal(t, test_user.LastName, found_user.LastName)
	assert.Equal(t, test_user.Organization[0].Name, found_user.Organization[0].Name)
	assert.Equal(t, test_user.Organization[0].Role, found_user.Organization[0].Role)

}

func TestUsernameAlreadyTaken(t *testing.T) {
	test_user := models.User{
		Username:     "LostProgrammer1010",
		Password:     "test",
		FirstName:    "test",
		LastName:     "test",
		Email:        "test",
		Organization: []models.Organization{{Name: "NSVC", Role: "ADMIN"}},
	}

	err := AddUser(test_user)
	assert.Error(t, err)

	found_user, err := GetUserByUsername(test_user.Username)
	assert.NoError(t, err)

	assert.Equal(t, test_user.Username, found_user.Username)
	assert.NotEqual(t, test_user.Email, found_user.Email)
	assert.False(t, authentication.VerifyPassword(found_user.Password, test_user.Password))
	assert.NotEqual(t, test_user.FirstName, found_user.FirstName)
	assert.NotEqual(t, test_user.LastName, found_user.LastName)
	assert.Equal(t, test_user.Organization[0].Name, found_user.Organization[0].Name)
	assert.Equal(t, test_user.Organization[0].Role, found_user.Organization[0].Role)
}

func TestEmailAlreadyTaken(t *testing.T) {
	test_user := models.User{
		Username:     "test",
		Password:     "test",
		FirstName:    "test",
		LastName:     "test",
		Email:        "check@gmail.com",
		Organization: []models.Organization{{Name: "NSVC", Role: "ADMIN"}},
	}

	err := AddUser(test_user)
	assert.Error(t, err)

	found_user, err := GetUserByUsername("LostProgrammer1010")
	assert.NoError(t, err)

	assert.NotEqual(t, test_user.Username, found_user.Username)
	assert.Equal(t, test_user.Email, found_user.Email)
	assert.False(t, authentication.VerifyPassword(found_user.Password, test_user.Password))
	assert.NotEqual(t, test_user.FirstName, found_user.FirstName)
	assert.NotEqual(t, test_user.LastName, found_user.LastName)
	assert.Equal(t, test_user.Organization[0].Name, found_user.Organization[0].Name)
	assert.Equal(t, test_user.Organization[0].Role, found_user.Organization[0].Role)
}
