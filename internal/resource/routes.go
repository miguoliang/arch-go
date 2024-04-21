package resource

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var CustomRealmName = viper.GetString("keycloak.custom.realm")

func SetupRoutes() *gin.Engine {

	r := gin.Default()

	api := r.Group("/api/v1")

	api.Group("/users").
		DELETE("/:id", DeleteUserHandler).
		DELETE("/:id/groups/:groupId", LeaveGroupHandler).
		GET("", ListUsersHandler).
		GET("/:id", GetUserHandler).
		GET("/:id/groups", ListGroupsByUserHandler).
		HEAD("", CheckUserHandler).
		POST("", CreateUserHandler).
		POST("/:id/groups/:groupId", JoinGroupHandler).
		PUT("/:id", UpdateUserHandler)

	api.Group("/groups").
		DELETE("/:id", DeleteGroupHandler).
		GET("", ListGroupsHandler).
		GET("/:id", GetGroupHandler).
		POST("", CreateGroupHandler).
		PUT("/:id", UpdateGroupHandler)

	api.Group("/roles").
		DELETE("/:id", DeleteRoleHandler).
		GET("", ListRolesHandler).
		GET("/:id", GetRoleHandler).
		HEAD("", CheckRoleHandler).
		POST("", CreateRoleHandler).
		PUT("/:id", UpdateRoleHandler)

	api.POST("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error": 0,
		})
	})

	return r
}
