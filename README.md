# test-examen

## Requisitos

* Go 1.17
* Swaggo( go get -u -v github.com/swaggo/swag/cmd/swag )
* Echo Swagger ( go get -u -v github.com/swaggo/echo-swagger )
* Template Text ( go get github.com/alecthomas/template )

* `swag init -g api/api.go -o api/docs`


## Compilación y Ejecución

* `BASE_URL=https://farmanet.minsal.cl/index.php/ws/getLocales go run api/api.go`

## Ejecución de Tests

* `BASE_URL=https://farmanet.minsal.cl/index.php/ws/getLocales go test ./...`

## Docker

Comandos para generación de contenedor de API. No es necesario para ambiente local.

* `docker build -t exam-test:1.0 .`
* `docker run -p 1323:1323 --name exam-test exam-test:1.0`
