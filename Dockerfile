# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

# Instalar dependências necessárias para o build
RUN apk add --no-cache git

WORKDIR /app

# Copiar apenas os arquivos de módulos primeiro para aproveitar o cache das camadas
COPY items-api/go.mod items-api/go.sum ./
RUN go mod download

# Copiar o restante do código fonte
COPY items-api/ ./

# Compilar a aplicação
RUN go build -o main main.go

# Porta padrão da aplicação
EXPOSE 9000

# Comando para rodar a aplicação
CMD ["./main"]
