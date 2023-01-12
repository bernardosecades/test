//go:build unit
// +build unit

package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/bernardosecades/test/internal/purchase/entity"
	"github.com/stretchr/testify/mock"
)

func TestErrorInvalidUUID(t *testing.T) {
	mock := purchaseRepositoryMock{}

	s := NewPurchaseService(&mock)

	ctx := context.TODO()
	uuid := "1d5309a3-c288"
	product := entity.Product{
		ID:        uuid,
		Available: true,
	}

	mock.ShouldGetProduct(ctx, uuid, &product, 0)
	mock.ShouldSetToNoAvailable(ctx, uuid, 0)

	// Execute
	err := s.Make(ctx, uuid)

	assert.ErrorIs(t, err, ErrInvalidID)
}

func TestErrorProductNoAvailable(t *testing.T) {
	mock := purchaseRepositoryMock{}

	s := NewPurchaseService(&mock)

	ctx := context.TODO()
	uuid := "1bd6d9c5-a1fb-428b-b2f8-2d366c4fc224"
	product := entity.Product{
		ID:        uuid,
		Available: false,
	}

	mock.ShouldGetProduct(ctx, uuid, &product, 1)
	mock.ShouldSetToNoAvailable(ctx, uuid, 0)

	// Execute
	err := s.Make(ctx, uuid)

	assert.ErrorIs(t, err, ErrProductNoAvailable)
}

func TestSuccessMakePurchase(t *testing.T) {
	mock := purchaseRepositoryMock{}

	s := NewPurchaseService(&mock)

	ctx := context.TODO()
	uuid := "1bd6d9c5-a1fb-428b-b2f8-2d366c4fc224"
	product := entity.Product{
		ID:        uuid,
		Available: true,
	}

	mock.ShouldGetProduct(ctx, uuid, &product, 1)
	mock.ShouldSetToNoAvailable(ctx, uuid, 1)

	// Execute
	err := s.Make(ctx, uuid)

	assert.Nil(t, err)
}

type purchaseRepositoryMock struct {
	mock.Mock
}

func (m *purchaseRepositoryMock) GetProduct(ctx context.Context, ID string) (*entity.Product, error) {
	args := m.Called(ctx, ID)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (m *purchaseRepositoryMock) SetToNoAvailable(ctx context.Context, ID string) error {
	args := m.Called(ctx, ID)
	return args.Error(0)
}

func (m *purchaseRepositoryMock) ShouldGetProduct(ctx context.Context, ID string, product *entity.Product, times int) {
	m.
		On("GetProduct", ctx, ID).
		Times(times).
		Return(product, nil)
}

func (m *purchaseRepositoryMock) ShouldSetToNoAvailable(ctx context.Context, ID string, times int) {
	m.
		On("SetToNoAvailable", ctx, ID).
		Times(times).
		Return(nil)
}
