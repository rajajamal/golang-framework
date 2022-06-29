package configs

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Test_Config_Debug(t *testing.T) {
	err := godotenv.Load(".env.test.debug")
	if err != nil {
		godotenv.Load("../.env.test.debug")
	}

	Load()

	assert.Equal(t, Env.Debug, true)

	err = godotenv.Overload(".env.test.debug-false")
	if err != nil {
		godotenv.Overload("../.env.test.debug-false")
	}

	Load()

	assert.Equal(t, Env.Debug, false)
}
