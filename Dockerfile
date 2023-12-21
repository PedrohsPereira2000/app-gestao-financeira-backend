# stage de build
FROM golang:1.21 AS build
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./src/main.go
# stage imagem final
# FROM scratch
# WORKDIR /app
# COPY --from=build /app/api ./
EXPOSE 8080
CMD [ "./api" ]