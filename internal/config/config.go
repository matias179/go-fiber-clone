package config

import (
	"os"
)

type AppEnvType string

const (
	ProdEnv  AppEnvType = "prod"
	StageEnv AppEnvType = "stage"
	DevEnv   AppEnvType = "dev"
	TestEnv  AppEnvType = "test"
)

// Envs - exported environment variables to use any part of the API.
var Envs envs

type envs struct {
	AppEnv AppEnvType
	PORT   string
}

func setAppEnv() AppEnvType {
	value := os.Getenv("APP_ENV")
	if value == "" {
		return DevEnv
	}

	return AppEnvType(value)
}

// SetEnvironmentVariables - grab all environment variables.
func SetEnvironmentVariables() {
	Envs = envs{
		AppEnv: setAppEnv(),
		PORT:   os.Getenv("PORT"),
	}
}
