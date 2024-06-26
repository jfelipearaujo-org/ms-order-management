package update

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jfelipearaujo-org/ms-order-management/internal/shared/custom_error"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	t.Run("Should return nil when dto is valid without items", func(t *testing.T) {
		// Arrange
		dto := UpdateOrderDto{
			OrderId: uuid.NewString(),
			State:   1,
		}

		// Act
		err := dto.Validate()

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return nil when dto is valid with items", func(t *testing.T) {
		// Arrange
		dto := UpdateOrderDto{
			OrderId: uuid.NewString(),
			State:   1,
			Items: []UpdateOrderItemDto{
				{
					ItemId:    uuid.NewString(),
					Name:      "name",
					UnitPrice: 10,
					Quantity:  1,
				},
			},
		}

		// Act
		err := dto.Validate()

		// Assert
		assert.NoError(t, err)
	})

	t.Run("Should return error when dto is invalid", func(t *testing.T) {
		// Arrange
		dto := UpdateOrderDto{}

		// Act
		err := dto.Validate()

		// Assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, custom_error.ErrRequestNotValid)
	})
}
