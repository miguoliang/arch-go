package keycloak

import (
	"context"
	"fmt"
	"github.com/miguoliang/keycloakadminclient"
	"log"
)

type RoleService interface {
	ListRoles() (*[]keycloakadminclient.RoleRepresentation, int, error)
	GetRoleById(roleId string) (*keycloakadminclient.RoleRepresentation, int, error)
	GetRoleByName(roleName string) (*keycloakadminclient.RoleRepresentation, int, error)
	CreateRole(role *keycloakadminclient.RoleRepresentation) (string, int, error)
	UpdateRole(roleId string, role *keycloakadminclient.RoleRepresentation) (*keycloakadminclient.RoleRepresentation, int, error)
	DeleteRole(roleId string) (int, error)
}

type roleService struct {
	client    *keycloakadminclient.APIClient
	realmName string
}

func NewRoleService(realmName string) RoleService {
	return &roleService{
		client:    GetAdminClient(),
		realmName: realmName,
	}
}

func (r *roleService) ListRoles() (*[]keycloakadminclient.RoleRepresentation, int, error) {
	roles, h, err := r.client.RolesAPI.
		AdminRealmsRealmRolesGet(context.Background(), r.realmName).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := CheckResponse(h, err)
	if err != nil {
		return nil, statusCode, err
	}
	return &roles, statusCode, nil
}

func (r *roleService) GetRoleById(roleId string) (*keycloakadminclient.RoleRepresentation, int, error) {
	role, h, err := r.client.RolesByIDAPI.
		AdminRealmsRealmRolesByIdRoleIdGet(context.Background(), r.realmName, roleId).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := CheckResponse(h, err)
	if err != nil {
		return nil, statusCode, err
	}
	return role, statusCode, nil
}

func (r *roleService) GetRoleByName(roleName string) (*keycloakadminclient.RoleRepresentation, int, error) {
	roles, _, err := r.ListRoles()
	if err != nil {
		return nil, 500, err
	}
	for _, role := range *roles {
		if *role.Name == roleName {
			return &role, 200, nil
		}
	}
	return nil, 404, fmt.Errorf("role name %s not found", roleName)
}

func (r *roleService) CreateRole(role *keycloakadminclient.RoleRepresentation) (string, int, error) {
	h, err := r.client.RolesAPI.
		AdminRealmsRealmRolesPost(context.Background(), r.realmName).
		RoleRepresentation(*role).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := CheckResponse(h, err)
	if h == nil {
		log.Println("http response is nil, but no error occurred.")
		return "", 500, fmt.Errorf("http response is nil, but no error occurred")
	} else if h.StatusCode != 201 {
		log.Println("Unexpected status code:", h.StatusCode)
		return "", h.StatusCode, fmt.Errorf("unexpected status code: %d", h.StatusCode)
	}

	newRole, statusCode, err := r.GetRoleByName(*role.Name)
	if err != nil {
		return "", statusCode, err
	}
	return newRole.GetId(), 201, nil
}

func (r *roleService) UpdateRole(roleId string, role *keycloakadminclient.RoleRepresentation) (*keycloakadminclient.RoleRepresentation, int, error) {
	h, err := r.client.RolesByIDAPI.
		AdminRealmsRealmRolesByIdRoleIdPut(context.Background(), r.realmName, roleId).
		RoleRepresentation(*role).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := CheckResponse(h, err)
	if err != nil {
		return nil, statusCode, err
	}
	return role, statusCode, nil
}

func (r *roleService) DeleteRole(roleId string) (int, error) {
	h, err := r.client.RolesByIDAPI.
		AdminRealmsRealmRolesByIdRoleIdDelete(context.Background(), r.realmName, roleId).
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
