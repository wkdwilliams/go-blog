package hashing_test

import (
	"testing"

	"github.com/wkdwilliams/go-blog/pkg/hashing"
)

func TestHashPassword(t *testing.T) {
	password, err := hashing.HashPassword("pass")

	if err != nil {
		t.Fatal(err)
	}

	if len(password) == 0 {
		t.Fatal("unexpected length of password")
	}
}

func TestVerifyPasswordHash(t *testing.T) {
	password, err := hashing.HashPassword("pass")

	if err != nil {
		t.Fatal(err)
	}

	if verified := hashing.VerifyPassword("pass", password); !verified {
		t.Fatal("could not verifiy password")
	}
}
