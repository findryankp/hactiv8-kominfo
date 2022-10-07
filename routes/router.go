package routes

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	r := gin.Default()
	orderRoutes(r)
	return r
}
