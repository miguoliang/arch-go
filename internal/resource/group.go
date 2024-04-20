package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/arch-go/internal/dto"
	"github.com/miguoliang/arch-go/internal/keycloak"
	"github.com/miguoliang/keycloakadminclient"
	"net/http"
)

const (
	RealmName = "custom"
)

// ListGroupsHandler List groups
// @Summary List groups
// @Description List groups
// @Tags group
// @Accept  json
// @Produce  json
// @Success 200 {array} keycloakadminclient.GroupRepresentation
// @Failure 500 {object} dto.ErrorResponse
// @Router /groups [get]
func ListGroupsHandler(c *gin.Context) {
	service := keycloak.NewGroupService(RealmName)
	groups, statusCode, err := service.ListGroups()
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, groups)
}

// GetGroupHandler Get group
// @Summary Get group
// @Description Get group
// @Tags group
// @Accept  json
// @Produce  json
// @Param id path string true "Group ID"
// @Success 200 {object} keycloakadminclient.GroupRepresentation
// @Failure 500 {object} dto.ErrorResponse
// @Router /groups/{id} [get]
func GetGroupHandler(c *gin.Context) {
	service := keycloak.NewGroupService(RealmName)
	groupId := c.Param("id")
	group, statusCode, err := service.GetGroup(groupId)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, group)
}

// CreateGroupHandler Create group
// @Summary Create group
// @Description Create group
// @Tags group
// @Accept  json
// @Produce  json
// @Param group body dto.Group true "Group"
// @Failure 500 {object} dto.ErrorResponse
// @Router /groups [post]
func CreateGroupHandler(c *gin.Context) {
	group := keycloakadminclient.GroupRepresentation{}
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(400, dto.ErrorResponse{Message: err.Error()})
		return
	}
	service := keycloak.NewGroupService(RealmName)
	groupId, statusCode, err := service.CreateGroup(&group)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(statusCode, dto.CreatedResponse{Id: groupId})
}

// UpdateGroupHandler Update group
// @Summary Update group
// @Description Update group
// @Tags group
// @Accept  json
// @Produce  json
// @Param id path string true "Group ID"
// @Param group body dto.Group true "Group"
// @Success 200
// @Failure 500 {object} dto.ErrorResponse
// @Router /groups/{id} [put]
func UpdateGroupHandler(c *gin.Context) {
	group := keycloakadminclient.GroupRepresentation{}
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(400, dto.ErrorResponse{Message: err.Error()})
		return
	}
	service := keycloak.NewGroupService(RealmName)
	groupId := c.Param("id")
	statusCode, err := service.UpdateGroup(groupId, &group)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.Status(statusCode)
}

// DeleteGroupHandler Delete group
// @Summary Delete group
// @Description Delete group
// @Tags group
// @Accept  json
// @Produce  json
// @Param id path string true "Group ID"
// @Success 200
// @Failure 500 {object} dto.ErrorResponse
// @Router /groups/{id} [delete]
func DeleteGroupHandler(c *gin.Context) {
	service := keycloak.NewGroupService(RealmName)
	groupId := c.Param("id")
	statusCode, err := service.DeleteGroup(groupId)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}
	c.Status(statusCode)
}
