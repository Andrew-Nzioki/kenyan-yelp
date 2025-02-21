package business

import (
	"time"

	"github.com/google/uuid"
)

// Request DTOs
type CreateBusinessRequest struct {
    Name        string  `json:"name" binding:"required"`
    Description string  `json:"description"`
    Category    string  `json:"category" binding:"required"`
    Location    string  `json:"location" binding:"required"`
    OpenHours   string  `json:"open_hours"`
    ContactInfo string  `json:"contact_info"`
}

type UpdateBusinessRequest struct {
    Name        string  `json:"name,omitempty"`
    Description string  `json:"description,omitempty"`
    Category    string  `json:"category,omitempty"`
    Location    string  `json:"location,omitempty"`
    OpenHours   string  `json:"open_hours,omitempty"`
    ContactInfo string  `json:"contact_info,omitempty"`
}

// Response DTOs
type BusinessResponse struct {
    ID          uuid.UUID  `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Category    string  `json:"category"`
    Location    string  `json:"location"`
    Rating      float64 `json:"rating"`
    OpenHours   string  `json:"open_hours"`
    ContactInfo string  `json:"contact_info"`
    CreatedAt   string  `json:"created_at"`
    UpdatedAt   string  `json:"updated_at"`
}

type ListBusinessesResponse struct {
    Businesses []BusinessResponse `json:"businesses"`
    Total     int                `json:"total"`
    Page      int                `json:"page"`
    PageSize  int                `json:"page_size"`
}
type ErrorResponse struct {
    Error string `json:"error"`
}

// Conversion methods
func (b *Business) ToResponse() BusinessResponse {
    return BusinessResponse{
        ID:          b.ID,
        Name:        b.Name,
        Description: b.Description,
        Category:    b.Category,
        Location:    b.Location,
        Rating:      b.Rating,

        CreatedAt:   b.CreatedAt.Format(time.RFC3339),
        UpdatedAt:   b.UpdatedAt.Format(time.RFC3339),
    }
}

func (req *CreateBusinessRequest) ToEntity() *Business {
    return &Business{
        Name:        req.Name,
        Description: req.Description,
        Category:    req.Category,
        Location:    req.Location,
        Rating:      0,

    }
}