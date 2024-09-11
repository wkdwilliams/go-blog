package web_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/wkdwilliams/go-blog/internal/adapters/primary/web"
	"github.com/wkdwilliams/go-blog/internal/adapters/secondary/database"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
	"github.com/wkdwilliams/go-blog/internal/infrastructure"
)

func TestStartAndStopServer(t *testing.T) {
	mockDB, _ := infrastructure.NewMysqlMock()
	
	server := web.NewApp(
		services.NewUserService(database.NewUserRepository(mockDB)),
		services.NewPostService(database.NewPostRepository(mockDB)),
		web.WithHideBanner(),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}

	errCh := make(chan error)

	go func() {
		errCh <- server.Start()
	}()

	select {
		// Test the shutting down of the server
		if err := server.Stop(ctx); err != nil {
			t.Fatalf("Error stopping the server: %v", err)
		}
	case err := <-errCh:
		t.Fatalf("Error starting the server: %v", err)
	}
}
