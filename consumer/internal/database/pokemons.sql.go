// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: pokemons.sql

package database

import (
	"context"
)

const insertPokemon = `-- name: InsertPokemon :one
INSERT INTO pokemons (id, species, name, description, updated_at, created_at)
VALUES ($1, $2, $3, $4, NOW(), NOW())
RETURNING id, species, name, description, updated_at, created_at
`

type InsertPokemonParams struct {
	ID          int64
	Species     string
	Name        string
	Description string
}

func (q *Queries) InsertPokemon(ctx context.Context, arg InsertPokemonParams) (Pokemon, error) {
	row := q.db.QueryRow(ctx, insertPokemon,
		arg.ID,
		arg.Species,
		arg.Name,
		arg.Description,
	)
	var i Pokemon
	err := row.Scan(
		&i.ID,
		&i.Species,
		&i.Name,
		&i.Description,
		&i.UpdatedAt,
		&i.CreatedAt,
	)
	return i, err
}
