# Use a imagem oficial do Golang como base
FROM golang:1.20.6

# Defina o diretório de trabalho no contêiner
WORKDIR /app

# Copie o go mod e o go sum files
COPY go.mod go.sum ./

# Baixe todas as dependências. 
# Se as dependências do go.mod e do go.sum não mudarem, então essa camada do Docker será reutilizada na próxima build
RUN go mod download

# Copie o código da fonte para o contêiner
COPY . .

# Construa o Go app
RUN go build -o main ./cmd/server

# Exponha a porta 50051 para o mundo fora deste contêiner
EXPOSE 50051

# Execute o binário compilado
CMD ["./main"]