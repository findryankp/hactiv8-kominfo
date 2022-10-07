package routes

import (
	"orders/controllers"

	"github.com/gin-gonic/gin"
)

func orderRoutes(r *gin.Engine) {
	orders := r.Group("/orders")
	{
		orders.GET("/", controllers.GetOrder)
		orders.POST("/", controllers.CreateOrder)
		orders.PUT("/:id", controllers.UpdateOrder)
		orders.DELETE("/:id", controllers.DeleteOrder)
	}
}
