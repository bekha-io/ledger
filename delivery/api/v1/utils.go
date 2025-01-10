package v1

import "github.com/gin-gonic/gin"

func handleError(c *gin.Context, err error) {
	c.JSON(400, gin.H{"error": err.Error()})
}
