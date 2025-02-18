package business

import "github.com/google/uuid"

type Business struct {
    ID          string
    Name        string
    OwnerID     uuid.UUID
    LocationID  uuid.UUID
    Status      string
    PriceRange  string   
    Rating      float64
}
// Core business rules/validations
