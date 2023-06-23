# Estágio de compilação
FROM golang:1.17-alpine AS build

# Definindo diretório de trabalho
WORKDIR /src

# Adicionando arquivos Go
ADD . .

# Compilando o programa
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app

# Estágio de execução
FROM alpine:latest

# Adicionando o binário compilado do estágio de compilação
COPY --from=build /bin/app /bin/app

# ...
ADD run.sh /run.sh
RUN chmod +x /run.sh


# Definindo a porta que será exposta pelo container
EXPOSE 3000

# Comando que será executado quando o container for iniciado
CMD ["/bin/app", "-port", "3000"]
