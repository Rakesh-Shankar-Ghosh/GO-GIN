package routes

import (
	"backend/controllers"
    "backend/filters"

	"github.com/gin-gonic/gin"
)

func PersonRoutes(router *gin.Engine) {
    PersonGroup := router.Group("/api/person")

    // Define the DELETE route and link it to the DeleteProduct function in the controller
    PersonGroup.GET("/Test", filters.Test_Middleware(), controllers.Test);
    PersonGroup.POST("/addPerson", controllers.AddPerson);
    PersonGroup.GET("/getPersons", controllers.GetPersons);
    PersonGroup.GET("/getPerson/:id", controllers.GetPersonByID);
    PersonGroup.PUT("/updatePerson/:id", controllers.UpdatePersonByID);
    PersonGroup.DELETE("/deletePerson/:id", controllers.DeletePersonByID);

    
}
