-- +goose Up
CREATE TABLE pokemons
(
    id          BIGSERIAL PRIMARY KEY,
    species     TEXT      NOT NULL,
    name        TEXT      NOT NULL,
    description TEXT      NOT NULL,
    updated_at  TIMESTAMP NOT NULL,
    created_at  TIMESTAMP NOT NULL

);

-- +goose Down
DROP TABLE pokemons;