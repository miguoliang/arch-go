package resource

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGroupsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetGroupsHandler"})
}

func GetGroupHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetGroupHandler"})
}

func CreateGroupHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "CreateGroupHandler"})
}

func RenameGroupHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateGroupHandler"})
}

func MoveGroupHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "MoveGroupHandler"})
}

func DeleteGroupHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "DeleteGroupHandler"})
}

func HasGroupHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "existsGroupHandler"})
}
