package consumer

import (
	"context"
	"github.com/NicholasMantovani/go-kafka-enterprise-example/consumer/internal/database"
	"github.com/NicholasMantovani/go-kafka-enterprise-example/consumer/internal/models"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type PokemonService interface {
	InsertPokemon(ctx context.Context, arg database.InsertPokemonParams) (models.Pokemon, error)
}

type PokemonConsumer struct {
	svc    PokemonService
	reader *kafka.Reader
	logger *zap.Logger
}

func NewPokemonConsumer(svc PokemonService, reader *kafka.Reader, logger *zap.Logger) *PokemonConsumer {
	return &PokemonConsumer{svc: svc, reader: reader, logger: logger}
}

func (c *PokemonConsumer) Consume(ctx context.Context) {

	for {
		message, err := c.reader.ReadMessage(ctx)
		if err != nil {
			c.logger.Error("error while fetching messages", zap.Error(err))
		}

		if _, err := c.svc.InsertPokemon(message); err != nil {

		}
	}

}
