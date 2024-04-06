package config

import "go.uber.org/zap"

func ConfigureLogging() *zap.Logger {
	if isProd() {
		logger, err := zap.NewProduction()
		if err != nil {
			panic(err)
		}
		return logger
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return logger
}
