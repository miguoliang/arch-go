package keycloak

import (
	"context"
	"fmt"
	"github.com/miguoliang/keycloakadminclient"
	"log"
)

type UserService interface {
	GetUserById(userId string) (*keycloakadminclient.UserRepresentation, int, error)
	GetUserByUsername(username string) (*keycloakadminclient.UserRepresentation, int, error)
	ListUsers() (*[]keycloakadminclient.UserRepresentation, int, error)
	CreateUser(user *keycloakadminclient.UserRepresentation) (string, int, error)
	UpdateUser(user *keycloakadminclient.UserRepresentation) (*keycloakadminclient.UserRepresentation, int, error)
	DeleteUser(userId string) (int, error)
	ListGroups(userId string) (*[]keycloakadminclient.GroupRepresentation, int, error)
	JoinGroup(userId string, groupId string) (int, error)
	LeaveGroup(userId string, groupId string) (int, error)
}

type userService struct {
	keycloakClient *keycloakadminclient.APIClient
	realmName      string
}

func NewUserService(realmName string) UserService {
	return &userService{
		keycloakClient: GetAdminClient(),
		realmName:      realmName,
	}
}

func (u *userService) ListGroups(userId string) (*[]keycloakadminclient.GroupRepresentation, int, error) {
	groups, h, err := u.keycloakClient.UsersAPI.
		AdminRealmsRealmUsersUserIdGroupsGet(context.Background(), u.realmName, userId).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := CheckResponse(h, err)
	if err != nil {
		return nil, statusCode, err
	}
	return &groups, statusCode, nil
}

func (u *userService) GetUserById(userId string) (*keycloakadminclient.UserRepresentation, int, error) {
	user, h, err := u.keycloakClient.UsersAPI.
		AdminRealmsRealmUsersUserIdGet(context.Background(), u.realmName, userId).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := CheckResponse(h, err)
	if err != nil {
		return nil, statusCode, err
	}
	return user, statusCode, nil
}

func (u *userService) GetUserByUsername(username string) (*keycloakadminclient.UserRepresentation, int, error) {
	users, h, err := u.keycloakClient.UsersAPI.
		AdminRealmsRealmUsersGet(context.Background(), u.realmName).
		Username(username).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := CheckResponse(h, err)
	if err != nil {
		return nil, statusCode, err
	}
	if len(users) > 0 {
		return &users[0], statusCode, nil
	}
	return nil, statusCode, nil
}

func (u *userService) ListUsers() (*[]keycloakadminclient.UserRepresentation, int, error) {
	users, h, err := u.keycloakClient.UsersAPI.
		AdminRealmsRealmUsersGet(context.Background(), u.realmName).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := CheckResponse(h, err)
	if err != nil {
		return nil, statusCode, err
	}
	return &users, statusCode, nil
}

func (u *userService) CreateUser(user *keycloakadminclient.UserRepresentation) (string, int, error) {
	h, err := u.keycloakClient.UsersAPI.
		AdminRealmsRealmUsersPost(context.Background(), u.realmName).
		UserRepresentation(*user).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := CheckResponse(h, err)
	if err != nil {
		return "", statusCode, err
	}

	if h == nil {
		log.Println("http response is nil, but no error occurred.")
		return "", 500, fmt.Errorf("http response is nil, but no error occurred")
	} else if h.StatusCode != 201 {
		log.Println("Unexpected status code:", h.StatusCode)
		return "", h.StatusCode, fmt.Errorf("unexpected status code: %d", h.StatusCode)
	}

	location := h.Header.Get("Location")
	userId := location[len(location)-36:]
	return userId, statusCode, nil
}

func (u *userService) UpdateUser(user *keycloakadminclient.UserRepresentation) (*keycloakadminclient.UserRepresentation, int, error) {
	h, err := u.keycloakClient.UsersAPI.
		AdminRealmsRealmUsersUserIdPut(context.Background(), u.realmName, *user.Id).
		UserRepresentation(*user).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := CheckResponse(h, err)
	if err != nil {
		return nil, statusCode, err
	}
	return user, statusCode, nil
}

func (u *userService) DeleteUser(userId string) (int, error) {
	h, err := u.keycloakClient.UsersAPI.
		AdminRealmsRealmUsersUserIdDelete(context.Background(), u.realmName, userId).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := CheckResponse(h, err)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (u *userService) JoinGroup(userId string, groupId string) (int, error) {
	h, err := u.keycloakClient.UsersAPI.
		AdminRealmsRealmUsersUserIdGroupsGroupIdPut(context.Background(), u.realmName, userId, groupId).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := CheckResponse(h, err)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (u *userService) LeaveGroup(userId string, groupId string) (int, error) {
	h, err := u.keycloakClient.UsersAPI.
		AdminRealmsRealmUsersUserIdGroupsGroupIdDelete(context.Background(), u.realmName, userId, groupId).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := CheckResponse(h, err)
	if err != nil {
		return statusCode, err
	}
	return statusCode, nil
}
