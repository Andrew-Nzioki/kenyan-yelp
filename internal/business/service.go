package business

import (
	"context"

	"github.com/google/uuid"
)

type service struct {
    repo BusinessRepository
}

func NewService(repo BusinessRepository) BusinessService {
    return &service{
        repo: repo,
    }
}

func (s *service) CreateBusiness(ctx context.Context, business *Business) error {
    return s.repo.Store(ctx, business)
}


func (s *service) GetBusinessByID(ctx context.Context, id uuid.UUID) (*Business, error) {
    return s.repo.FindByID(ctx, id)
}