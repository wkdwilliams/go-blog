package services_test

import (
	"math/rand"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
	"github.com/wkdwilliams/go-blog/mocks"
)

func TestCanCreatePost(t *testing.T) {
	expectedPost := models.NewPost("test title", "test content", uuid.New())

	mockPostRepository := mocks.NewPostRepository(t)
	mockPostRepository.On("Create", mock.Anything).Return(nil)

	postService := services.NewPostService(mockPostRepository)
	post, err := postService.Create(expectedPost.Title, expectedPost.Content, expectedPost.UserID)

	assert.Nil(t, err)
	assert.Equal(t, expectedPost.Title, post.Title)
	assert.Equal(t, expectedPost.Content, post.Content)
	assert.Equal(t, expectedPost.UserID, post.UserID)
}

func TestCanCreatePostFuzzer(t *testing.T) {
	mockPostRepository := mocks.NewPostRepository(t)
	mockPostRepository.On("Create", mock.Anything).Return(nil)

	postService := services.NewPostService(mockPostRepository)

	faker := gofakeit.New(0)

	for i := 0; i < 200; i++ {
		var (
			title   = faker.BookTitle()
			content = faker.Sentence(rand.Intn(50-1+1) + 1)
			userId  = uuid.New()
		)

		post, err := postService.Create(title, content, userId)

		assert.Nil(t, err)                     // Assert that the returned error is nil
		assert.Equal(t, post.Title, title)     // Assert that the returned post title is equal to the title we inputted
		assert.Equal(t, post.Content, content) // Assert that the returned content is equal to the content we inputted
		assert.Equal(t, post.UserID, userId)   // Assert that the returned user id is equal to the user id we inputted

		mockPostRepository.AssertCalled(t, "Create", mock.Anything)
	}
}

func TestCanGetPostById(t *testing.T) {
	faker := gofakeit.New(0)

	expectedPost := models.NewPost(faker.BookTitle(), faker.Sentence(rand.Intn(50-1+1)+1), uuid.New())

	mockPostRepository := mocks.NewPostRepository(t)
	mockPostRepository.On("GetById", expectedPost.ID).Return(&expectedPost, nil)

	postService := services.NewPostService(mockPostRepository)
	post, err := postService.GetById(expectedPost.ID)

	assert.Nil(t, err)                   // Assert that the returned error is nil
	assert.Equal(t, post, &expectedPost) // Assert that the struct and the returned struct are the same
}

func TestGetAllPosts(t *testing.T) {
	mockPostRepository := mocks.NewPostRepository(t)
	faker := gofakeit.New(0)
	// Mocking multiple users
	mockPosts := []models.Post{
		models.NewPost(faker.BookTitle(), faker.Sentence(rand.Intn(50-1+1)+1), uuid.New()),
		models.NewPost(faker.BookTitle(), faker.Sentence(rand.Intn(50-1+1)+1), uuid.New()),
		models.NewPost(faker.BookTitle(), faker.Sentence(rand.Intn(50-1+1)+1), uuid.New()),
		models.NewPost(faker.BookTitle(), faker.Sentence(rand.Intn(50-1+1)+1), uuid.New()),
	}
	mockPostRepository.On("GetAll").Return(mockPosts, nil)

	postService := services.NewPostService(mockPostRepository)
	posts, err := postService.GetAll()

	// Assert no errors and the users are returned
	assert.Nil(t, err)
	assert.Len(t, posts, 4)
	assert.Equal(t, mockPosts[0], posts[0])
	assert.Equal(t, mockPosts[1], posts[1])
	assert.Equal(t, mockPosts[2], posts[2])
	assert.Equal(t, mockPosts[3], posts[3])

	// Ensure that the mock was called
	mockPostRepository.AssertCalled(t, "GetAll")
}

func TestCanDeletePost(t *testing.T) {
	mockPostRepository := mocks.NewPostRepository(t)
	mockPostRepository.On("Delete", mock.Anything).Return(nil)

	mockPostService := services.NewPostService(mockPostRepository)
	err := mockPostService.Delete(uuid.New())

	assert.Nil(t, err)

	mockPostRepository.AssertCalled(t, "Delete", mock.Anything)
}
