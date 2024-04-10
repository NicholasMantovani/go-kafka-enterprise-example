package main

import (
	"context"
	"github.com/NicholasMantovani/go-kafka-enterprise-example/consumer/internal/config"
	consumer2 "github.com/NicholasMantovani/go-kafka-enterprise-example/consumer/internal/consumer"
	"github.com/NicholasMantovani/go-kafka-enterprise-example/consumer/internal/database"
	"github.com/NicholasMantovani/go-kafka-enterprise-example/consumer/internal/service"
	"os/signal"
	"syscall"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	conn := config.ConfigureDatabase()

	defer conn.Close(ctx)

	pokemonReader := config.ConfigurePokemonReader()

	defer pokemonReader.Close()

	queries := database.New(conn)

	logger := config.ConfigureLogging()

	svc := service.NewPokemonServiceImpl(queries, logger)

	consumer := consumer2.NewPokemonConsumer(svc, pokemonReader, logger)

	go consumer.Consume(ctx)

	select {
	case <-ctx.Done():
		done()
	}
}
