package business

import (
	"errors"
	"time"

	"github.com/google/uuid"
)


type Business struct {
    ID          uuid.UUID
    Name        string
    Description string
    Category    string
    Location    string
    Rating      float64
    ContactInfo string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

func (b *Business) Validate() error {
    if b.Name == "" {
        return errors.New("name is required")
    }
   
    return nil
}

func (b *Business) BeforeCreate() {
    b.CreatedAt = time.Now()
    b.UpdatedAt = time.Now()
}