package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"

	"github.com/joho/godotenv"
	"github.com/wkdwilliams/go-blog/internal/adapters/secondary/database"
	"github.com/wkdwilliams/go-blog/internal/domain/models"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
	"github.com/wkdwilliams/go-blog/internal/infrastructure"
)

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := infrastructure.NewMysql()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := database.NewUserRepository(db)
	usersService := services.NewUserService(userRepo)

	postRepo := database.NewPostRepository(db)
	postService := services.NewPostService(postRepo)

	// ##################### Seed users #####################

	userFile, err := os.Open("cmd/seed/data/users.json")
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	var createdUsers []models.User

	if err := json.NewDecoder(userFile).Decode(&users); err != nil {
		log.Fatal(err)
	}

	for _, v := range users {
		user, err := usersService.CreateAccount(v.Username, v.Password, v.Name)
		if err != nil {
			log.Fatal(err)
		}

		createdUsers = append(createdUsers, *user)
	}

	// ##################### Seed posts #####################

	postsFile, err := os.Open("cmd/seed/data/posts.json")
	if err != nil {
		log.Fatal(err)
	}

	var posts []Post

	if err := json.NewDecoder(postsFile).Decode(&posts); err != nil {
		log.Fatal(err)
	}

	for _, post := range posts {
		if len(createdUsers) == 1 {
			postService.Create(post.Title, post.Content, createdUsers[0].ID)
			continue
		}
		postService.Create(post.Title, post.Content, createdUsers[rand.Intn((len(createdUsers)-1)-0)+0].ID)
	}
}
