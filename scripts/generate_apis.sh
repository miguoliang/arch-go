#!/bin/bash

docker run --rm -v $(pwd):/code \
  -v $GOPATH/pkg/mod/github.com/miguoliang/keycloakadminclient@v0.0.0-20240416114625-bd88bf8cfb6b:/go/pkg/mod/github.com/miguoliang/keycloakadminclient@v0.0.0-20240416114625-bd88bf8cfb6b \
  ghcr.io/swaggo/swag:latest init -g ./cmd/arch-go/main.go -o /code/api -ot yaml -d ./,/go/pkg/mod/github.com/miguoliang/keycloakadminclient@v0.0.0-20240416114625-bd88bf8cfb6b

docker run --rm -v $(pwd):/local openapitools/openapi-generator-cli generate \
  -i /local/api/swagger.yaml -o /local/api -g openapi-yaml --additional-properties=x-extension.openapi=3.0.0,outputFile=openapi.yaml
