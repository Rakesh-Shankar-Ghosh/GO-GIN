package routes

import (
	"backend/controllers"
    "backend/filters"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
    ProductGroup := router.Group("/api/v2",filters.Test_Middleware())

    // Define the DELETE route and link it to the DeleteProduct function in the controller
    ProductGroup.GET("/Testt", controllers.Testt);
}
