package main

import (
	"backend/config"
	"backend/migrations"
	
	"backend/routes"
    "backend/middlewares"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	"log"

	"os"
)

func main() {
	
	db, err := config.InitDB()
    if err != nil {
        log.Fatalf("Error initializing database: %v", err)
    }
    
    // Defer the closing of the database connection
    defer db.Close()

	migrations.Migrate(db)
    


	// person := models.Person{
	// 	Username: "Rakesh Shankar",
	// 	Email:    "rakesh@example.com",
	// 	Address: "Boyra Khulna",
	// }
	
	// db.Create(&person)
	

    router := gin.Default()
    router.Use(middlewares.CorsMiddleware())
    
    routes.PersonRoutes(router)
	routes.ProductRoutes(router)

    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Start the server using the SERVER_PORT environment variable
    serverPort := os.Getenv("SERVER_PORT")

    if err := router.Run(":" + serverPort); err != nil {
        panic(err)
	}
}


    // // Create a new Gin router
    // router := gin.Default()

    // // Define a route
    // router.GET("/", func(c *gin.Context) {
    //     c.JSON(http.StatusOK, gin.H{
    //         "message": "Hello, Gin!",
    //     })
    // })

    // // Start the server
    // router.Run(":8080")