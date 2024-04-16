package main

import (
	"github.com/miguoliang/keycloakadminclient"
	"log"
)

var (
	// Keycloak API client
	keycloakApiClient *keycloakadminclient.APIClient
)

func main() {

	SetupLog()

	log.Println("Started!")
}
