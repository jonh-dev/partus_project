package main

import (
	"net"
	"os"

	"github.com/jonh-dev/go-logger/logger"
	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/config"
	"github.com/jonh-dev/partus_users/internal/encryption"
	"github.com/jonh-dev/partus_users/internal/repositories"
	"github.com/jonh-dev/partus_users/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	logger.Info("Iniciando o servidor...")

	envGetter := config.NewEnvVarGetter()

	certFile, err := envGetter.Get("SSL_CERT_FILE")

	if err != nil {
		logger.Fatal("Falha ao buscar o SSL_CERT_FILE: " + err.Error())
	}

	keyFile, err := envGetter.Get("SSL_KEY_FILE")
	if err != nil {
		logger.Fatal("Falha ao buscar o SSL_KEY_FILE: " + err.Error())
	}

	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		logger.Fatal("Certificado não encontrado: " + err.Error())
	}

	logger.Info("Carregando certificados...")
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		logger.Fatal("Falha ao setar o TLS: " + err.Error())
	}

	logger.Info("Criando servidor...")
	s := grpc.NewServer(grpc.Creds(creds))

	logger.Info("Registrando serviços...")
	dbService, err := config.NewDBService(envGetter)
	if err != nil {
		logger.Fatal("Falha ao criar o DBService: " + err.Error())
	}

	passwordEncryptor := &encryption.BcryptPasswordEncryptor{}
	repo := repositories.NewUserRepository(dbService)
	personalInfoRepo := repositories.NewPersonalInfoRepository(dbService)
	accountInfoRepo := repositories.NewAccountInfoRepository(dbService)

	personalInfoService := services.NewPersonalInfoService(personalInfoRepo)
	accountInfoService := services.NewAccountInfoService(accountInfoRepo, passwordEncryptor)
	service := services.NewUserService(repo, personalInfoService, accountInfoService)

	api.RegisterUserServiceServer(s, service)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Acessando o servidor na porta " + port)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Fatal("Failed to listen: " + err.Error())
	}

	if err := s.Serve(lis); err != nil {
		logger.Fatal("Falha ao inciar o servidor: " + err.Error())
	}
}
