package business

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

// BusinessFilter for searching/filtering businesses
type BusinessFilter struct {
    Category    string
    Location    string
    MinRating   float64
    SearchQuery string
    Limit       int
    Offset      int
}



type BusinessRepository interface {
    Store(ctx context.Context, business *Business) error
    FindByID(ctx context.Context, id uuid.UUID) (*Business, error)
    List(ctx context.Context, filter BusinessFilter) ([]Business, error)
    Update(ctx context.Context, business *Business) error
    Delete(ctx context.Context, id uuid.UUID) error
}

type BusinessService interface {
    CreateBusiness(ctx context.Context, business *Business) error
    GetBusinessByID(ctx context.Context, id uuid.UUID) (*Business, error)
}

var (
    ErrBusinessNotFound = errors.New("business not found")
    ErrInvalidBusinessData = errors.New("invalid business data")
)