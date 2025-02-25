// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Business struct {
	ID          uuid.UUID
	Name        string
	Description sql.NullString
	Category    string
	Location    string
	Rating      sql.NullString
	ContactInfo sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
