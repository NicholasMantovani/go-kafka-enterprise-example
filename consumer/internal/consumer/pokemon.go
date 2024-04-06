package consumer

import (
	"context"
	"github.com/NicholasMantovani/go-kafka-enterprise-example/consumer/internal/database"
	"github.com/NicholasMantovani/go-kafka-enterprise-example/consumer/internal/models"
)

type PokemonService interface {
	InsertPokemon(ctx context.Context, arg database.InsertPokemonParams) (models.Pokemon, error)
}
