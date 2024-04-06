-- name: InsertPokemon :one
INSERT INTO pokemons (id, species, name, description, updated_at, created_at)
VALUES ($1, $2, $3, $4, NOW(), NOW())
RETURNING *;
