package main

import (
	"log"
	"net"

	"github.com/jonh-dev/partus_users/api"
	"github.com/jonh-dev/partus_users/internal/config"
	"github.com/jonh-dev/partus_users/internal/repositories"
	"github.com/jonh-dev/partus_users/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	log.Println("Iniciando o servidor...")

	envGetter := config.NewEnvVarGetter(&config.FileEnvLoader{})

	certFile, err := envGetter.Get("SSL_CERT_FILE")

	if err != nil {
		log.Fatalf("Falha ao buscar o SSL_CERT_FILE: %v", err)
	}

	keyFile, err := envGetter.Get("SSL_KEY_FILE")
	if err != nil {
		log.Fatalf("Falha ao buscar o SSL_KEY_FILE: %v", err)
	}

	log.Println("Carregando certificados...")
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("Falha ao setar o TLS: %v", err)
	}

	log.Println("Criando servidor...")
	s := grpc.NewServer(grpc.Creds(creds))

	log.Println("Registrando serviços...")
	dbService, err := config.NewDBService()
	if err != nil {
		log.Fatalf("Falha ao criar o DBService: %v", err)
	}

	repo := repositories.NewUserRepository(dbService)
	personalInfoRepo := repositories.NewPersonalInfoRepository(dbService)
	accountInfoRepo := repositories.NewAccountInfoRepository(dbService)

	personalInfoService := services.NewPersonalInfoService(personalInfoRepo)
	accountInfoService := services.NewAccountInfoService(accountInfoRepo)
	service := services.NewUserService(repo, personalInfoService, accountInfoService)

	api.RegisterUserServiceServer(s, service)

	log.Println("Acesso ao servidor na porta 8080...")
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Falha ao escutar:  %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falha ao inciar o servidor: %v", err)
	}
}
