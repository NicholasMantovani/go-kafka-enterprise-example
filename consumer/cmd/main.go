package main

import "github.com/NicholasMantovani/go-kafka-enterprise-example/consumer/internal/config"

func main() {
	conn := config.ConfigureDatabase()
	logger := config.ConfigureLogging()
	pokemonReader := config.ConfigurePokemonReader()

}
