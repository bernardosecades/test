package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/bernardosecades/test/internal/kit/uuid"

	"github.com/bernardosecades/test/internal/purchase/entity"
)

var (
	// ErrInvalidID used when we receive an invalid ID
	ErrInvalidID = errors.New("is not a valid UUID")
	// ErrProductNoAvailable used when the product is not available
	ErrProductNoAvailable = errors.New("product no available")
	// ErrProductNotFound used when the product does not exist in database
	ErrProductNotFound = errors.New("product not found")
)

// PurchaseRepository interface determine which method should
// be implemented in order to suit the purchase service
type PurchaseRepository interface {
	GetProduct(ctx context.Context, ID string) (*entity.Product, error)
	SetToNoAvailable(ctx context.Context, ID string) error
}

// Purchase type encapsulates all the base for the
// purchase service
type Purchase struct {
	purchaseRepository PurchaseRepository
}

// NewPurchaseService method to instance a purchase service
func NewPurchaseService(p PurchaseRepository) Purchase {
	return Purchase{purchaseRepository: p}
}

// Make method to make a purchase
func (p Purchase) Make(ctx context.Context, ID string) error {
	if !uuid.IsValidUUID(ID) {
		return fmt.Errorf("'%s' %w", ID, ErrInvalidID)
	}

	product, err := p.purchaseRepository.GetProduct(ctx, ID)
	if err != nil {
		return err
	}

	if product == nil {
		return fmt.Errorf("id: '%s' %w", ID, ErrProductNotFound)
	}

	if !product.Available {
		return fmt.Errorf("id: '%s' %w", ID, ErrProductNoAvailable)
	}

	return p.purchaseRepository.SetToNoAvailable(ctx, ID)
}
