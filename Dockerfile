# Use uma imagem base do Golang como ponto de partida
FROM golang:1.17

# Defina o diretório de trabalho no contêiner
WORKDIR /app

# Copie todo o código-fonte da aplicação Go para o contêiner
COPY . .

# Instale as dependências e construa a aplicação
RUN go get -d -v ./...
RUN go build -o myapp

# Exponha a porta em que a aplicação Go irá rodar (por padrão, a porta 8080)
EXPOSE 8080

# Comando para iniciar a aplicação Go
CMD ["./myapp"]
