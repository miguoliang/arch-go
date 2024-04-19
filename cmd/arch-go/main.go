package main

import (
	"github.com/gin-gonic/gin"
	"github.com/miguoliang/arch-go/internal/resource"
	"gopkg.in/Graylog2/go-gelf.v2/gelf"
	"io"
	"log"
	"os"
)

const (
	graylogAddr = "localhost:12201"
)

// @title Arch-Go API
// @description This is the API for Arch-Go
// @version 1.0
// @host localhost:8080
// @BasePath /api
// @schemes http
// @schemes https
// @contact.name Guoliang Mi
// @contact.email boymgl@qq.com
// @contact.url https://miguoliang.com
func routes() *gin.Engine {

	r := gin.Default()

	api := r.Group("/api/v1")

	api.Group("/users").
		DELETE("/:id", resource.DeleteUserHandler).
		DELETE("/:id/group/:groupId", resource.LeaveGroupHandler).
		DELETE("/:id/role/:roleId", resource.RevokeRoleHandler).
		GET("/", resource.ListUsersHandler).
		GET("/:id", resource.GetUserHandler).
		GET("/:id/groups", resource.ListGroupsByUserHandler).
		GET("/:id/roles", resource.ListRolesByUserHandler).
		HEAD("/:username", resource.CheckUserHandler).
		POST("/", resource.CreateUserHandler).
		POST("/:id/group/:groupId", resource.JoinGroupHandler).
		POST("/:id/role/:roleId", resource.AssignRoleHandler).
		PUT("/:id", resource.UpdateUserHandler)

	api.Group("/groups").
		DELETE("/:id", resource.DeleteGroupHandler).
		GET("/", resource.ListGroupsHandler).
		GET("/:id", resource.GetGroupHandler).
		HEAD("/", resource.CheckGroupHandler).
		POST("/", resource.CreateGroupHandler).
		POST("/:id/rename", resource.RenameGroupHandler).
		POST("/:id/move", resource.MoveGroupHandler)

	api.Group("/roles").
		DELETE("/:id", resource.DeleteRoleHandler).
		GET("/", resource.ListRolesHandler).
		GET("/:id", resource.GetRoleHandler).
		HEAD("/", resource.CheckRoleHandler).
		POST("/", resource.CreateRoleHandler).
		POST("/:id/rename", resource.RenameRoleHandler).
		POST("/:id/permissions", resource.AssignPermissionsToRoleHandler)

	api.POST("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"error": 0,
		})
	})

	return r
}

func main() {

	setupLog()

	r := routes()

	err := r.Run("0.0.0.0:8081")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	log.Println("Started!")
}

func setupLog() {

	gelfWriter, err := gelf.NewUDPWriter(graylogAddr)
	if err != nil {
		log.Fatalf("gelf.NewWriter: %s", err)
	}
	// log to both stderr and graylog2
	log.SetOutput(io.MultiWriter(os.Stderr, gelfWriter))
}
