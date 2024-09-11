package web_test

import (
	"context"
	"testing"
	"time"

	"github.com/wkdwilliams/go-blog/internal/adapters/primary/web"
	"github.com/wkdwilliams/go-blog/internal/adapters/secondary/database"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
)

func TestStartAndStopServer(t *testing.T) {
	server := web.NewApp(
		services.NewUserService(&database.MemoryUserRepository{}),
		services.NewPostService(&database.MemoryPostRepository{}),
		web.WithHideBanner(),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	errCh := make(chan error)

	go func() {
		errCh <- server.Start()
	}()

	select {
	case <-ctx.Done():
		// If there's no error after
		// the context has timed out,
		// then we have no error starting
		// the server.

		// Test the shutting down of the server
		if err := server.Stop(ctx); err != nil {
			t.Fatalf("Error stopping the server: %v", err)
		}
	case err := <-errCh:
		t.Fatalf("Error starting the server: %v", err)
	}
}
