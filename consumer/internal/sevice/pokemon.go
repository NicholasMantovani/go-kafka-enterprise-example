package sevice

import (
	"context"
	"github.com/NicholasMantovani/go-kafka-enterprise-example/consumer/internal/database"
	"github.com/NicholasMantovani/go-kafka-enterprise-example/consumer/internal/models"
	"go.uber.org/zap"
)

type PokemonServiceImpl struct {
	queries *database.Queries
	logger  *zap.Logger
}

func New(queries *database.Queries, logger *zap.Logger) *PokemonServiceImpl {
	return &PokemonServiceImpl{queries: queries, logger: logger}
}

func (p *PokemonServiceImpl) InsertPokemon(ctx context.Context, arg database.InsertPokemonParams) (models.Pokemon, error) {

	p.logger.Info("Inserting new pokemon", getPokemonLoggingInfo(arg)...)

	res, err := p.queries.InsertPokemon(ctx, arg)
	if err != nil {
		p.logger.Error("Error while inserting new pokemon", append(getPokemonLoggingInfo(arg), zap.Error(err))...)
		return models.Pokemon{}, err
	}
	p.logger.Info("Inserted pokemon", getPokemonLoggingInfo(arg)...)

	return fromDbToDto(res), err

}

func fromDbToDto(from database.Pokemon) models.Pokemon {
	return models.Pokemon{
		ID:          from.ID,
		Species:     from.Species,
		Name:        from.Name,
		Description: from.Description,
	}
}

func getPokemonLoggingInfo(arg database.InsertPokemonParams) []zap.Field {
	var fields []zap.Field

	fields = append(fields, zap.Int64("ID", arg.ID))
	return append(fields, zap.String("NAME", arg.Name))
}
