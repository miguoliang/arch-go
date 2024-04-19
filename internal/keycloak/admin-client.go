package keycloak

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/miguoliang/keycloakadminclient"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	keycloakServerURL = "http://localhost:8080/auth"
	realmName         = "master"
	clientID          = "admin-cli"
	username          = "admin"
	password          = "admin"
)

var (
	mutex             sync.Mutex
	accessToken       string
	refreshToken      string
	expiryTime        time.Time
	keycloakApiClient *keycloakadminclient.APIClient
)

func refreshAccessToken() {
	mutex.Lock()
	defer mutex.Unlock()

	if accessToken == "" {
		initAccessToken()
		return
	}

	if time.Now().Before(expiryTime.Add(-5 * time.Minute)) { // Refresh before expiry - adjust time as needed
		tokenData, err := retrieveTokenFromRemote()
		if err != nil {
			return
		}

		newAccessToken, ok := tokenData["access_token"].(string)
		if !ok {
			log.Fatalf("access_token missing in response")
			return
		}
		accessToken = newAccessToken

		refreshToken, ok = tokenData["refresh_token"].(string)
		if !ok {
			log.Fatalf("refresh_token missing in response")
			return
		}

		expiry, ok := tokenData["expires_in"].(float64)
		if !ok {
			log.Fatalf("expires_in missing in response")
			return
		}

		expiryTime = time.Now().Add(time.Duration(expiry) * time.Second)
	}
}

func retrieveTokenFromRemote() (map[string]interface{}, error) {
	tokenURL := fmt.Sprintf("%s/auth/realms/%s/protocol/openid-connect/token", keycloakServerURL, realmName)

	formData := url.Values{}
	formData.Set("grant_type", "refresh_token") // Assuming you have a refresh token mechanism
	formData.Set("client_id", clientID)
	formData.Set("refresh_token", refreshToken)

	requestBody := strings.NewReader(formData.Encode())

	req, err := http.NewRequest(http.MethodPost, tokenURL, requestBody)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("failed to refresh token: status code %d", resp.StatusCode)
		return nil, err
	}

	var tokenData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&tokenData)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	return tokenData, nil
}

func initAccessToken() {

	form := url.Values{}
	form.Add("client_id", clientID)
	form.Add("username", username)
	form.Add("password", password)
	form.Add("grant_type", "password")

	endpoint := fmt.Sprintf("%s/realms/%s/protocol/openid-connect/token", keycloakServerURL, realmName)
	resp, err := http.PostForm(endpoint, form)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(responseBody, &jsonData)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	accessToken = jsonData["access_token"].(string)
	refreshToken = jsonData["refresh_token"].(string)
	expiry := jsonData["expires_in"].(float64)
	expiryTime = time.Now().Add(time.Duration(expiry) * time.Second)
}

func GetAdminClient() *keycloakadminclient.APIClient {

	refreshAccessToken()

	if keycloakApiClient != nil {
		keycloakApiClient.GetConfig().DefaultHeader["Authorization"] = "Bearer " + accessToken
		return keycloakApiClient
	}

	configuration := keycloakadminclient.NewConfiguration()
	configuration.AddDefaultHeader("Authorization", "Bearer "+accessToken)
	configuration.Servers = keycloakadminclient.ServerConfigurations{
		{
			URL: keycloakServerURL,
		},
	}
	keycloakApiClient = keycloakadminclient.NewAPIClient(configuration)
	return keycloakApiClient
}

func CheckResponse(h *http.Response, err error) (int, error) {
	if err == nil {
		return 0, nil
	}

	var apiError *keycloakadminclient.GenericOpenAPIError
	if errors.As(err, &apiError) {
		return h.StatusCode, err
	}

	return 500, err
}
