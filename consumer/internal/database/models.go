// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Pokemon struct {
	ID          int64
	Species     string
	Name        string
	Description string
	UpdatedAt   pgtype.Timestamp
	CreatedAt   pgtype.Timestamp
}
