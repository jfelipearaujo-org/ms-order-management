package server

import (
	"testing"

	"github.com/jfelipearaujo-org/ms-order-management/internal/environment"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	t.Run("Should return a new server", func(t *testing.T) {
		// Arrange
		config := &environment.Config{
			ApiConfig: &environment.ApiConfig{
				Port: 8080,
			},
			DbConfig: &environment.DatabaseConfig{
				Url: "postgres://host:1234",
			},
		}

		// Act
		server := NewServer(config)

		// Assert
		assert.NotNil(t, server)
		assert.Equal(t, ":8080", server.Addr)
	})
}
