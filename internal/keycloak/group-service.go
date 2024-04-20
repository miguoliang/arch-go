package keycloak

import (
	"context"
	"fmt"
	"github.com/miguoliang/keycloakadminclient"
	"log"
)

type GroupService interface {
	CreateGroup(group *keycloakadminclient.GroupRepresentation) (string, int, error)
	GetGroup(groupId string) (*keycloakadminclient.GroupRepresentation, int, error)
	UpdateGroup(groupId string, group *keycloakadminclient.GroupRepresentation) (int, error)
	DeleteGroup(groupId string) (int, error)
	ListGroups() (*[]keycloakadminclient.GroupRepresentation, int, error)
}

type groupService struct {
	keycloakClient *keycloakadminclient.APIClient
	realmName      string
}

// CreateGroup creates a new group.
func (g *groupService) CreateGroup(group *keycloakadminclient.GroupRepresentation) (string, int, error) {
	h, err := g.keycloakClient.GroupsAPI.
		AdminRealmsRealmGroupsPost(context.TODO(), g.realmName).
		GroupRepresentation(*group).
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
	groupId := location[len(location)-36:]
	return groupId, statusCode, nil
}

// GetGroup gets a group by its id.
func (g *groupService) GetGroup(groupId string) (*keycloakadminclient.GroupRepresentation, int, error) {
	groupRepresentation, h, err := g.keycloakClient.GroupsAPI.
		AdminRealmsRealmGroupsGroupIdGet(context.TODO(), g.realmName, groupId).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := CheckResponse(h, err)
	if err != nil {
		return nil, statusCode, err
	}
	return groupRepresentation, statusCode, nil
}

// UpdateGroup updates a group.
func (g *groupService) UpdateGroup(groupId string, group *keycloakadminclient.GroupRepresentation) (int, error) {
	h, err := g.keycloakClient.GroupsAPI.
		AdminRealmsRealmGroupsGroupIdPut(context.TODO(), g.realmName, groupId).
		GroupRepresentation(*group).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	return CheckResponse(h, err)
}

// DeleteGroup deletes a group by its id.
func (g *groupService) DeleteGroup(groupId string) (int, error) {
	h, err := g.keycloakClient.GroupsAPI.
		AdminRealmsRealmGroupsGroupIdDelete(context.TODO(), g.realmName, groupId).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	return CheckResponse(h, err)
}

// ListGroups gets all groups.
func (g *groupService) ListGroups() (*[]keycloakadminclient.GroupRepresentation, int, error) {
	groups, h, err := g.keycloakClient.GroupsAPI.
		AdminRealmsRealmGroupsGet(context.TODO(), g.realmName).
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

func NewGroupService(realmName string) GroupService {
	return &groupService{
		keycloakClient: GetAdminClient(),
		realmName:      realmName,
	}
}
