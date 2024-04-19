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
// @Success 200 {object} dto.GroupList
// @Failure 500 {object} dto.ErrorResponse
// @Router /groups [get]
func ListGroupsHandler(c *gin.Context) {
	admin := keycloak.GetAdminClient()
	groupRepresentations, h, err := admin.GroupsAPI.AdminRealmsRealmGroupsGet(c, RealmName).Execute()

	if h != nil {
		defer h.Body.Close()
	}

	statusCode, err := keycloak.CheckResponse(h, err)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}

	groups := make([]dto.Group, len(groupRepresentations))
	for i, group := range groupRepresentations {
		groups[i] = dto.Group{
			Id:   *group.Id,
			Name: *group.Name,
			Path: *group.Path,
		}
	}
	c.JSON(http.StatusOK, dto.GroupList{Items: groups})
}

// GetGroupHandler Get group
// @Summary Get group
// @Description Get group
// @Tags group
// @Accept  json
// @Produce  json
// @Param id path string true "Group ID"
// @Success 200 {object} dto.Group
// @Failure 500 {object} dto.ErrorResponse
// @Router /groups/{id} [get]
func GetGroupHandler(c *gin.Context) {
	admin := keycloak.GetAdminClient()
	groupId := c.Param("id")
	group, h, err := admin.GroupsAPI.AdminRealmsRealmGroupsGroupIdGet(c, RealmName, groupId).Execute()

	if h != nil {
		defer h.Body.Close()
	}

	statusCode, err := keycloak.CheckResponse(h, err)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Group{
		Id:   groupId,
		Name: *group.Name,
		Path: *group.Path,
	})
}

// CreateGroupHandler Create group
// @Summary Create group
// @Description Create group
// @Tags group
// @Accept  json
// @Produce  json
// @Param group body dto.Group true "Group"
// @Success 200 {object} string
// @Failure 500 {object} dto.ErrorResponse
// @Router /groups [post]
func CreateGroupHandler(c *gin.Context) {

	admin := keycloak.GetAdminClient()
	group := dto.Group{}
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(400, dto.ErrorResponse{Message: err.Error()})
		return
	}

	groupRepresentation := keycloakadminclient.GroupRepresentation{
		Name: &group.Name,
		Path: &group.Path,
	}
	h, err := admin.GroupsAPI.AdminRealmsRealmGroupsPost(c, RealmName).
		GroupRepresentation(groupRepresentation).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := keycloak.CheckResponse(h, err)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(200, "")
}

// RenameGroupHandler Rename group
// @Summary Rename group
// @Description Rename group
// @Tags group
// @Accept  json
// @Produce  json
// @Param id path string true "Group ID"
// @Param group body dto.Group true "Group"
// @Success 200 {object} string
// @Failure 500 {object} dto.ErrorResponse
// @Router /groups/{id} [put]
func RenameGroupHandler(c *gin.Context) {
	admin := keycloak.GetAdminClient()
	groupId := c.Param("id")
	group := dto.Group{}
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(400, dto.ErrorResponse{Message: err.Error()})
		return
	}

	groupRepresentation := keycloakadminclient.GroupRepresentation{
		Name: &group.Name,
		Path: &group.Path,
	}

	h, err := admin.GroupsAPI.AdminRealmsRealmGroupsGroupIdPut(c, RealmName, groupId).
		GroupRepresentation(groupRepresentation).
		Execute()
	if h != nil {
		defer h.Body.Close()
	}
	statusCode, err := keycloak.CheckResponse(h, err)
	if err != nil {
		c.JSON(statusCode, dto.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(200, "")
}

func MoveGroupHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "MoveGroupHandler"})
}

func DeleteGroupHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteGroupHandler"})
}

func CheckGroupHandler(c *gin.Context) {

}
