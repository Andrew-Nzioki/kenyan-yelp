package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/Andrew-Nzioki/kenyan-yelp/internal/business"
	"github.com/google/uuid"
)

type postgresRepository struct {
    q *Queries
}

func NewPostgresRepository(db *sql.DB) business.BusinessRepository {
    return &postgresRepository{
        q: New(db),
    }
}

func (r *postgresRepository) Store(ctx context.Context, b *business.Business) error {
    return r.q.CreateBusiness(ctx, CreateBusinessParams{
        ID:          b.ID,
        Name:        b.Name,
        Description: sql.NullString{String: b.Description, Valid: b.Description != ""},
        Category:    b.Category,
        Location:    b.Location,
        Rating:      sql.NullString{String: fmt.Sprintf("%.2f", b.Rating), Valid: true},
        ContactInfo: sql.NullString{String: b.ContactInfo, Valid: b.ContactInfo != ""},
        CreatedAt:   b.CreatedAt,
        UpdatedAt:   b.UpdatedAt,
    })
}

func (r *postgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*business.Business, error) {
    b, err := r.q.GetBusiness(ctx, id)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, business.ErrBusinessNotFound
        }
        return nil, err
    }
    
    rating, _ := strconv.ParseFloat(b.Rating.String, 64)
    return &business.Business{
        ID:          b.ID,
        Name:        b.Name,
        Description: b.Description.String,
        Category:    b.Category,
        Location:    b.Location,
        Rating:      rating,
        ContactInfo: b.ContactInfo.String,
        CreatedAt:   b.CreatedAt,
        UpdatedAt:   b.UpdatedAt,
    }, nil
}

func (r *postgresRepository) List(ctx context.Context, filter business.BusinessFilter) ([]business.Business, error) {
    businesses, err := r.q.ListBusinesses(ctx, ListBusinessesParams{
  
        Limit:    int32(filter.Limit),
        Offset:   int32(filter.Offset),
    })
    if err != nil {
        return nil, err
    }

    result := make([]business.Business, len(businesses))
    for i, b := range businesses {
        rating, _ := strconv.ParseFloat(b.Rating.String, 64)
        result[i] = business.Business{
            ID:          b.ID,
            Name:        b.Name,
            Description: b.Description.String,
            Category:    b.Category,
            Location:    b.Location,
            Rating:      rating,
            ContactInfo: b.ContactInfo.String,
            CreatedAt:   b.CreatedAt,
            UpdatedAt:   b.UpdatedAt,
        }
    }
    return result, nil
}

func (r *postgresRepository) Update(ctx context.Context, b *business.Business) error {
    err := r.q.UpdateBusiness(ctx, UpdateBusinessParams{
        ID:          b.ID,
        Name:        b.Name,
        Description: sql.NullString{String: b.Description, Valid: b.Description != ""},
        Category:    b.Category,
        Location:    b.Location,
        Rating:      sql.NullString{String: fmt.Sprintf("%.2f", b.Rating), Valid: true},
        ContactInfo: sql.NullString{String: b.ContactInfo, Valid: b.ContactInfo != ""},
        UpdatedAt:   b.UpdatedAt,
    })
    if err == sql.ErrNoRows {
        return business.ErrBusinessNotFound
    }
    return err
}

func (r *postgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
    err := r.q.DeleteBusiness(ctx, id)
    if err == sql.ErrNoRows {
        return business.ErrBusinessNotFound
    }
    return err
}