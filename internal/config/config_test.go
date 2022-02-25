package config_test

import (
	"os"
	"testing"

	"github.com/matias179/go-fiber-clone/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestSetEnvironmentVariables(t *testing.T) {
	port := "3563"
	appEnv := "test"

	os.Setenv("PORT", port)
	os.Setenv("APP_ENV", appEnv)

	config.SetEnvironmentVariables()

	assert.Equal(t, port, config.Envs.PORT)
	assert.Equal(t, config.TestEnv, config.Envs.AppEnv)
}
