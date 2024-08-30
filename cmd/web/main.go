package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/wkdwilliams/go-blog/internal/adapters/primary/web"
	"github.com/wkdwilliams/go-blog/internal/adapters/secondary/mysql"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
	"github.com/wkdwilliams/go-blog/internal/infrastructure"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Create instance of database
	db, err := infrastructure.NewMysql()
	if err != nil {
		log.Fatal(err)
	}
	// Initialise secondary port implementations (Secondary adapters)
	userRepo := mysql.NewUserRepository(db) // <- this is swappable since its just a repo implementation
	// Initialise core service layer
	usersService := services.NewUserService(userRepo) // core business logic doesn't change.

	postRepo := mysql.NewPostRepository(db)
	postService := services.NewPostService(postRepo)

	// Init primary (driving) adapter
	// this is swapple since we can spin up another primary adapter, and inject business logic
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Fatal(err)
	}

	srv := web.NewApp(usersService, postService, web.WithPort(port))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()

	log.Println("shutting down...")
	if err := srv.Stop(context.Background()); err != nil {
		log.Fatal(err)
	}
}
