package config

import "os"

const PROD = "prod"
const DEV = "dev"
const profile = "PROFILE"

func isProd() bool {
	return os.Getenv(profile) == PROD
}

func isDev(profile string) bool {
	return os.Getenv(profile) == DEV
}
