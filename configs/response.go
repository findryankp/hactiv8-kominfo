package configs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseJson(c *gin.Context, code int, status, msg, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": msg,
		"data":    data,
	})
}
