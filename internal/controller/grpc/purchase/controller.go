package purchase

import (
	"context"
	"github.com/bernardosecades/test/proto/go/protobuf/api"
)

// PurchaseService interface determine which method should
// be implemented in order to suit the purchase controller
type PurchaseService interface {
	Make(ctx context.Context, ID string) error
}

// Controller type encapsulates all the base for the
// purchase handlers
type Controller struct {
	PurchaseService PurchaseService
}

// Purchase make a purchase
func (s *Controller) Purchase(ctx context.Context, req *api.PurchaseRequest) (*api.PurchaseResponse, error) {
	err := s.PurchaseService.Make(ctx, req.Id)
	if err != nil {
		return &api.PurchaseResponse{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &api.PurchaseResponse{
		Success: true,
		Error:   "",
	}, nil
}
