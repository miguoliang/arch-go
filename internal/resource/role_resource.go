package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/arch-go/internal/dto"
	"github.com/miguoliang/arch-go/internal/keycloak"
	"github.com/miguoliang/keycloakadminclient"
)

// GetRoleHandler get role by id
// @Summary Get role by id
// @Description Get role by id
// @Tags role
// @Accept json
// @Produce json
// @Param roleId path string true "Role ID"
// @Success 200 {object} keycloakadminclient.RoleRepresentation
// @Failure 400
// @Failure 404
// @Router /roles/{roleId} [get]
func GetRoleHandler(c *gin.Context) {
	service := keycloak.NewRoleService(CustomRealmName)
	roleId := c.Param("id")
	role, statusCode, err := service.GetRoleById(roleId)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(statusCode, role)
}

// ListRolesHandler list all roles
// @Summary List all roles
// @Description List all roles
// @Tags role
// @Accept json
// @Produce json
// @Success 200 {array} keycloakadminclient.RoleRepresentation
// @Failure 400
// @Failure 404
// @Router /roles [get]
func ListRolesHandler(c *gin.Context) {
	service := keycloak.NewRoleService(CustomRealmName)
	roles, statusCode, err := service.ListRoles()
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(statusCode, roles)
}

// CreateRoleHandler create a new role
// @Summary Create a new role
// @Description Create a new role
// @Tags role
// @Accept json
// @Produce json
// @Param role body object true "Role"
// @Success 201 {object} dto.CreatedResponse
// @Failure 400
// @Failure 409
// @Router /roles [post]
func CreateRoleHandler(c *gin.Context) {
	service := keycloak.NewRoleService(CustomRealmName)
	var role keycloakadminclient.RoleRepresentation
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(400, dto.ErrorResponse{Message: err.Error()})
		return
	}
	roleId, statusCode, err := service.CreateRole(&role)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(statusCode, dto.CreatedResponse{Id: roleId})
}

// DeleteRoleHandler delete role by id
// @Summary Delete role by id
// @Description Delete role by id
// @Tags role
// @Accept json
// @Produce json
// @Param roleId path string true "Role ID"
// @Success 204
// @Failure 400
// @Failure 404
// @Router /roles/{roleId} [delete]
func DeleteRoleHandler(c *gin.Context) {
	service := keycloak.NewRoleService(CustomRealmName)
	roleId := c.Param("id")
	statusCode, err := service.DeleteRole(roleId)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.Status(statusCode)
}

// UpdateRoleHandler update role by id
// @Summary Update role by id
// @Description Update role by id
// @Tags role
// @Accept json
// @Produce json
// @Param roleId path string true "Role ID"
// @Param role body object true "Role"
// @Success 200 {object} keycloakadminclient.RoleRepresentation
// @Failure 400
// @Failure 404
// @Router /roles/{roleId} [put]
func UpdateRoleHandler(c *gin.Context) {
	service := keycloak.NewRoleService(CustomRealmName)
	var role keycloakadminclient.RoleRepresentation
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(400, dto.ErrorResponse{Message: err.Error()})
		return
	}
	roleId := c.Param("id")
	r, statusCode, err := service.UpdateRole(roleId, &role)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(statusCode, r)
}

// CheckRoleHandler check role name
// @Summary Check role name
// @Description Check role name
// @Tags role
// @Accept json
// @Produce json
// @Param roleName query string true "Role Name"
// @Success 204
// @Failure 400
// @Failure 404
// @Router /roles/check [get]
func CheckRoleHandler(c *gin.Context) {
	service := keycloak.NewRoleService(CustomRealmName)
	roleName := c.Query("roleName")
	_, statusCode, err := service.GetRoleByName(roleName)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.Status(204)
}
