package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/arch-go/internal/dto"
	"github.com/miguoliang/arch-go/internal/keycloak"
	"github.com/miguoliang/arch-go/pkg/str"
	"github.com/miguoliang/keycloakadminclient"
)

// GetUserHandler get user by id
// @Summary Get user by id
// @Description Get user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} keycloakadminclient.UserRepresentation
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /users/{id} [get]
func GetUserHandler(c *gin.Context) {
	service := keycloak.NewUserService(RealmName)
	userID := c.Param("id")
	user, statusCode, err := service.GetUserById(userID)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(statusCode, user)
}

// CreateUserHandler create user
// @Summary Create user
// @Description Create user
// @Tags user
// @Accept json
// @Produce json
// @Param user body keycloakadminclient.UserRepresentation true "User"
// @Success 201 {object} dto.CreatedResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Router /users [post]
func CreateUserHandler(c *gin.Context) {
	service := keycloak.NewUserService(RealmName)
	var user keycloakadminclient.UserRepresentation
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, dto.ErrorResponse{Message: err.Error()})
		return
	}
	userId, statusCode, err := service.CreateUser(&user)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(statusCode, dto.CreatedResponse{Id: userId})
}

// UpdateUserHandler update user
// @Summary Update user
// @Description Update user
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body keycloakadminclient.UserRepresentation true "User"
// @Success 200 {object} keycloakadminclient.UserRepresentation
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /users/{id} [put]
func UpdateUserHandler(c *gin.Context) {
	service := keycloak.NewUserService(RealmName)
	var user keycloakadminclient.UserRepresentation
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, dto.ErrorResponse{Message: err.Error()})
		return
	}
	user.Id = str.Ptr(c.Param("id"))
	u, statusCode, err := service.UpdateUser(&user)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(statusCode, u)
}

// DeleteUserHandler delete user
// @Summary Delete user
// @Description Delete user
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Failure 400 {object} dto.ErrorResponse
// @Router /users/{id} [delete]
func DeleteUserHandler(c *gin.Context) {
	service := keycloak.NewUserService(RealmName)
	userID := c.Param("id")
	statusCode, err := service.DeleteUser(userID)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.Status(statusCode)
}

// ListUsersHandler list users
// @Summary List users
// @Description List users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {array} keycloakadminclient.UserRepresentation
// @Failure 400 {object} dto.ErrorResponse
// @Router /users [get]
func ListUsersHandler(c *gin.Context) {
	service := keycloak.NewUserService(RealmName)
	users, statusCode, err := service.ListUsers()
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(statusCode, users)
}

// JoinGroupHandler join group
// @Summary Join group
// @Description Join group
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param groupId path string true "Group ID"
// @Success 204
// @Failure 400 {object} dto.ErrorResponse
// @Router /users/{id}/groups/{groupId} [post]
func JoinGroupHandler(c *gin.Context) {
	service := keycloak.NewUserService(RealmName)
	userID := c.Param("id")
	groupID := c.Param("groupId")
	statusCode, err := service.JoinGroup(userID, groupID)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.Status(statusCode)
}

// LeaveGroupHandler leave group
// @Summary Leave group
// @Description Leave group
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param groupId path string true "Group ID"
// @Success 204
// @Failure 400 {object} dto.ErrorResponse
// @Router /users/{id}/groups/{groupId} [delete]
func LeaveGroupHandler(c *gin.Context) {
	service := keycloak.NewUserService(RealmName)
	userID := c.Param("id")
	groupID := c.Param("groupId")
	statusCode, err := service.LeaveGroup(userID, groupID)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.Status(statusCode)
}

// ListGroupsByUserHandler list groups by user
// @Summary List groups by user
// @Description List groups by user
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} keycloakadminclient.GroupRepresentation
// @Failure 400 {object} dto.ErrorResponse
func ListGroupsByUserHandler(c *gin.Context) {
	service := keycloak.NewUserService(RealmName)
	userID := c.Param("id")
	groups, statusCode, err := service.ListGroups(userID)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(statusCode, groups)
}

// CheckUserHandler check user
// @Summary Check user
// @Description Check user
// @Tags user
// @Accept json
// @Produce json
// @Param username query string true "Username"
// @Success 200
// @Failure 404
// @Router /users [head]
func CheckUserHandler(c *gin.Context) {
	service := keycloak.NewUserService(RealmName)
	username := c.Query("username")
	user, statusCode, err := service.GetUserByUsername(username)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	if user == nil {
		c.Status(404)
		return
	}
	c.Status(200)
}
