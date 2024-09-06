package hashing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wkdwilliams/go-blog/pkg/hashing"
)

func TestHashPassword(t *testing.T) {
	password, err := hashing.HashPassword("pass")

	assert.Nil(t, err)
	assert.Greater(t, len(password), 0)
}

func TestVerifyPasswordHash(t *testing.T) {
	password, err := hashing.HashPassword("pass")

	assert.Nil(t, err)
	assert.True(t, hashing.VerifyPassword("pass", password))
}
