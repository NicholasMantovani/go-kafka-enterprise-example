package config

import "github.com/segmentio/kafka-go"

func ConfigurePokemonReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "pokemons",
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
}
