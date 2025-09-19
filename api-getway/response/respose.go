// response.go
package response

import (
	

	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, status int, data interface{}, err error) {
	if err != nil {
		c.JSON(status, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"success": true,
		"data":    data,
	})
}
