#!/bin/bash

docker run --rm -v $(pwd):/code ghcr.io/swaggo/swag:latest init -g /cmd/arch-go/main.go -o /code/api -ot yaml

docker run --rm -v $(pwd):/local openapitools/openapi-generator-cli generate \
  -i /local/api/swagger.yaml -o /local/api -g openapi-yaml --additional-properties=x-extension.openapi=3.0.0,outputFile=openapi.yaml