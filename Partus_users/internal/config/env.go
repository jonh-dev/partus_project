package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/jonh-dev/go-logger/logger"
)

type EnvVarGetter struct {
	env string
}

func NewEnvVarGetter() *EnvVarGetter {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	logger.Info(fmt.Sprintf("Você está usando o ambiente %s", env))

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	envPath := filepath.Join(dir, fmt.Sprintf("../../.env.%s", env))

	err := godotenv.Load(envPath)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Erro ao carregar o arquivo .env.%s", env))
	}

	return &EnvVarGetter{
		env: env,
	}
}

func (e *EnvVarGetter) Get(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("%s não encontrada", key)
	}
	return val, nil
}
