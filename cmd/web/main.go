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
	"github.com/wkdwilliams/go-blog/internal/adapters/secondary/database"
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

	// Initialize the services
	userRepo := database.NewUserRepository(db)
	usersService := services.NewUserService(userRepo)

	postRepo := database.NewPostRepository(db)
	postService := services.NewPostService(postRepo)

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Fatal(err)
	}

	srv := web.NewApp(usersService, postService, web.WithPort(port))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server.
	<-ctx.Done()

	if err := srv.Stop(context.Background()); err != nil {
		log.Fatal(err)
	}

	d, err := db.DB()

	if err != nil {
		log.Fatal(err)
	}

	if err := d.Close(); err != nil {
		log.Fatal(err)
	}
}
