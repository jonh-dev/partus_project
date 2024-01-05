package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
Estrutura de dados que representa o serviço de banco de dados.

- Client: Cliente do MongoDB.

- DBName: Nome do banco de dados.
*/
type DBService struct {
	Client *mongo.Client
	DBName string
}

/*
Função que cria um novo serviço de banco de dados.

@returns *DBService, error

- Esta função é usada para criar um novo serviço de banco de dados. Ele retorna um erro, caso ocorra algum problema ao criar o serviço.
*/
func NewDBService() (*DBService, error) {
	envGetter := NewEnvVarGetter(&FileEnvLoader{})

	inContainer, err := envGetter.Get("IN_CONTAINER")
	if err != nil {
		return nil, err
	}

	var uri string
	if inContainer == "true" {
		uri, err = envGetter.Get("DOCKER_MONGO_URI")
		if err != nil {
			return nil, err
		}
	} else {
		uri, err = envGetter.Get("MONGO_URI")
		if err != nil {
			return nil, err
		}
	}

	dbName, err := envGetter.Get("DB_NAME")
	if err != nil {
		return nil, err
	}

	clientOptions := options.Client().ApplyURI(uri)

	// Try to connect to MongoDB multiple times before giving up
	var client *mongo.Client
	for i := 0; i < 5; i++ {
		client, err = mongo.Connect(context.Background(), clientOptions)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to MongoDB, retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return nil, err
	}

	return &DBService{Client: client, DBName: dbName}, nil
}

/*
Função que conecta ao MongoDB.

@params uri string

@params clientOptions *options.ClientOptions

@returns *mongo.Client, error

- Esta função é usada para conectar ao MongoDB. Ela usa a string de conexão para se conectar ao MongoDB. Ela retorna um erro, caso ocorra algum problema ao se conectar ao MongoDB.
*/
func connectToMongoDB(uri string, clientOptions *options.ClientOptions) (*mongo.Client, error) {
	return mongo.Connect(context.Background(), clientOptions)
}
