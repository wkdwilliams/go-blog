package services_test

import (
	"math/rand"
	"sync"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
	"github.com/wkdwilliams/go-blog/mocks"
	"github.com/wkdwilliams/go-blog/pkg/hashing"
)

func TestCanCreateUser(t *testing.T) {
	mockUserRepository := mocks.NewUserRepository(t)
	mockUserRepository.On("Create", mock.Anything).Return(nil)

	userService := services.NewUserService(mockUserRepository)

	const (
		username = "testuser"
		password = "testpass"
		name     = "testname"
	)

	user, err := userService.CreateAccount(username, password, name)

	assert.Nil(t, err)                          // Assert that the returned error is nil
	assert.Equal(t, user.Username, username)    // Assert that the returned username is equal to the username we inputted
	assert.NotEqual(t, user.Password, password) // Assert that the returned password is NOT equal to the password we inputted (password is hashed)
	assert.Equal(t, user.Name, name)            // Assert that the returned name is equal to the name we inputted

	mockUserRepository.AssertCalled(t, "Create", mock.Anything)
}

// This test is slow due to the hashing function
func TestCanCreateUserFuzzer(t *testing.T) {
	mockUserRepository := mocks.NewUserRepository(t)
	mockUserRepository.On("Create", mock.Anything).Return(nil)

	userService := services.NewUserService(mockUserRepository)

	faker := gofakeit.New(0)

	wg := &sync.WaitGroup{}

	for i := 0; i < 50; i++ {

		wg.Add(1)

		// We have to run inside a go routine because this just takes too long...
		go func() {
			var (
				username = faker.Username()
				password = faker.Password(true, true, true, true, true, rand.Intn(50-1+1)+1)
				name     = faker.Name()
			)

			user, err := userService.CreateAccount(username, password, name)

			assert.Nil(t, err)                          // Assert that the returned error is nil
			assert.Equal(t, user.Username, username)    // Assert that the returned username is equal to the username we inputted
			assert.NotEqual(t, user.Password, password) // Assert that the returned password is NOT equal to the password we inputted (password is hashed)
			assert.Equal(t, user.Name, name)            // Assert that the returned name is equal to the name we inputted

			mockUserRepository.AssertCalled(t, "Create", mock.Anything)
			wg.Done()
		}()
	}

	wg.Wait()
}

func TestCanGetUserById(t *testing.T) {
	expectedUser := &models.User{
		ID: uuid.New(),
	}

	mockUserRepository := mocks.NewUserRepository(t)
	mockUserRepository.On("GetById", mock.Anything).Return(expectedUser, nil)

	userService := services.NewUserService(mockUserRepository)
	user, err := userService.GetById(expectedUser.ID)

	assert.Nil(t, err)
	assert.Equal(t, expectedUser.ID, user.ID)

	mockUserRepository.AssertCalled(t, "GetById", expectedUser.ID)
}

func TestCanGetByUsername(t *testing.T) {
	mockUserRepository := mocks.NewUserRepository(t)

	expectedUser := &models.User{Username: "testuser"}
	mockUserRepository.On("GetByUsername", expectedUser.Username).Return(expectedUser, nil)

	userService := services.NewUserService(mockUserRepository)
	user, err := userService.GetByUsername(expectedUser.Username)

	// Assert no errors and the correct user is returned
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.Username, user.Username)

	mockUserRepository.AssertCalled(t, "GetByUsername", expectedUser.Username)
}

func TestGetAllUsers(t *testing.T) {
	mockUserRepository := mocks.NewUserRepository(t)

	// Mocking multiple users
	mockUsers := []models.User{
		{Username: "user1"},
		{Username: "user2"},
	}
	mockUserRepository.On("GetAll").Return(mockUsers, nil)

	userService := services.NewUserService(mockUserRepository)
	users, err := userService.GetAll()

	// Assert no errors and the users are returned
	assert.Nil(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, mockUsers[0].Username, users[0].Username)
	assert.Equal(t, mockUsers[1].Username, users[1].Username)

	// Ensure that the mock was called
	mockUserRepository.AssertCalled(t, "GetAll")
}

func TestShouldAuthenticateUser(t *testing.T) {
	mockUserRepository := mocks.NewUserRepository(t)

	const password = "testpass"
	hashedPassword, _ := hashing.HashPassword(password)

	// Mocking user and password verification
	expectedUser := &models.User{Username: "testuser", Password: hashedPassword}
	mockUserRepository.On("GetByUsername", "testuser").Return(expectedUser, nil)

	userService := services.NewUserService(mockUserRepository)
	user, err := userService.Authenticate("testuser", password)

	// Assert no errors and the user is returned
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.Username)

	// Ensure that the mock was called
	mockUserRepository.AssertCalled(t, "GetByUsername", "testuser")
}

func TestShouldNotAuthenticateUser(t *testing.T) {
	mockUserRepository := mocks.NewUserRepository(t)

	const password = "testpass"
	hashedPassword, _ := hashing.HashPassword(password)

	// Mocking user and password verification
	expectedUser := &models.User{Username: "testuser", Password: hashedPassword}
	mockUserRepository.On("GetByUsername", "testuser").Return(expectedUser, nil)

	userService := services.NewUserService(mockUserRepository)
	user, err := userService.Authenticate("testuser", "WrongPassword") // We insert the wrong password here

	// Assert that there ARE errors and the user is nil
	assert.NotNil(t, err)
	assert.Nil(t, user)

	// Ensure that the mock was called
	mockUserRepository.AssertCalled(t, "GetByUsername", "testuser")
}

func TestGetTotalUserCount(t *testing.T) {
	mockUserRepository := mocks.NewUserRepository(t)

	// Mocking the total count
	mockUserRepository.On("GetTotalCount").Return(int64(10))

	userService := services.NewUserService(mockUserRepository)
	count := userService.GetTotalCount()

	// Assert the count is correct
	assert.Equal(t, int64(10), count)

	// Ensure that the mock was called
	mockUserRepository.AssertCalled(t, "GetTotalCount")
}
