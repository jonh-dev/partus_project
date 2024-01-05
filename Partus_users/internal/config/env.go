package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

/*
Interface que representa um carregador de variáveis de ambiente.

- Load: Carrega as variáveis de ambiente de um arquivo.
*/
type EnvLoader interface {
	Load(string) error
}

/*
Estrutura de dados que representa um carregador de variáveis de ambiente de um arquivo.
*/
type FileEnvLoader struct{}

/*
Função que carrega as variáveis de ambiente de um arquivo.

@param path string

- Esta função é usada para carregar as variáveis de ambiente de um arquivo. Ela retorna um erro, caso ocorra algum problema ao carregar as variáveis de ambiente.
*/
func (f *FileEnvLoader) Load(path string) error {
	return godotenv.Load(path)
}

/*
Estrutura de dados que representa um carregador de variáveis de ambiente.

- envLoader: Carregador de variáveis de ambiente.
*/
type EnvVarGetter struct {
	envLoader EnvLoader
}

/*
Função que cria um novo carregador de variáveis de ambiente.

@param loader EnvLoader

@returns *EnvVarGetter

- Esta função é usada para criar um novo carregador de variáveis de ambiente. Ela retorna um carregador de variáveis de ambiente.
*/
func NewEnvVarGetter(loader EnvLoader) *EnvVarGetter {
	return &EnvVarGetter{
		envLoader: loader,
	}
}

/*
Função que retorna o valor de uma variável de ambiente.

@param key string

@returns string, error

- Esta função é usada para retornar o valor de uma variável de ambiente. Ela retorna um erro, caso ocorra algum problema ao retornar o valor da variável de ambiente.
*/
func (e *EnvVarGetter) Get(key string) (string, error) {
	val := os.Getenv(key)
	if val != "" {
		return val, nil
	}

	inContainer := os.Getenv("IN_CONTAINER")
	if inContainer == "true" {
		val = os.Getenv(key + "_CONTAINER")
		if val != "" {
			return val, nil
		}
	}

	env := e.getEnv()
	absEnvPath, err := e.getAbsEnvPath(env)
	if err != nil {
		return "", err
	}

	err = e.envLoader.Load(absEnvPath)
	if err != nil {
		return "", err
	}

	val = os.Getenv(key)
	if val == "" {
		return "", errors.New(key + " não encontrada")
	}

	return val, nil
}

/*
Função que retorna o valor da variável de ambiente APP_ENV.

@returns string

- Esta função é usada para retornar o valor da variável de ambiente APP_ENV. Ela retorna o valor da variável de ambiente APP_ENV.
*/
func (e *EnvVarGetter) getEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		return "development"
	}

	log.Println("Você está no ambiente de " + env)

	return env
}

/*
Função que retorna o caminho absoluto do arquivo de variáveis de ambiente.

@param env string

@returns string, error

- Esta função é usada para retornar o caminho absoluto do arquivo de variáveis de ambiente. Ela retorna um erro, caso ocorra algum problema ao retornar o caminho absoluto do arquivo de variáveis de ambiente.
*/
func (e *EnvVarGetter) getAbsEnvPath(env string) (string, error) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	envPath := filepath.Join(dir, fmt.Sprintf("../../.env.%s", env))
	return filepath.Abs(envPath)
}
