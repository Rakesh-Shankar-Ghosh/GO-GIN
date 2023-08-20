package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)



func Test(c *gin.Context) {
    data := 5
    if data == 5 {
        c.JSON(http.StatusOK, gin.H{
            "success": true,
            "message": "perfectly done",
            "data":    data,
        })
    } else {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "message": "failed",
        })
    }
}





func UpdatePersonByID(c *gin.Context) {
    // Get the person ID from the URL parameters
    personID := c.Param("id")

    // Initialize the database connection
    db, _ := config.InitDB()

    // Check if the person with the specified ID exists
    var existingPerson models.Person
    if err := db.Where("id = ?", personID).First(&existingPerson).Error; err != nil {
        c.JSON(404, gin.H{
            "success": false,
            "message": "Person not found",
        })
        return
    }

    // Define a struct to bind the request body data to
    var request models.Person

    // Bind the request body to the request struct
    if err := c.BindJSON(&request); err != nil {
        c.JSON(400, gin.H{
            "success": false,
            "message": "Invalid request data",
        })
        return
    }

    // Update the existing person's information
    existingPerson.Username = request.Username
    existingPerson.Email = request.Email
    existingPerson.Address = request.Address

    if err := db.Save(&existingPerson).Error; err != nil {
        c.JSON(500, gin.H{
            "success": false,
            "message": "Error while updating the person",
            "error":   err.Error(),
        })
        return
    }

    c.JSON(200, gin.H{
        "success": true,
        "message": "Person updated successfully",
        "data":    existingPerson,
    })
}


func AddPerson(c *gin.Context) {
    // Define a struct to bind the request body data to
    var request models.Person

    // Bind the request body to the request struct
    if err := c.BindJSON(&request); err != nil {
        c.JSON(400, gin.H{
            "success": false,
            "message": "Invalid request data",
        })
        return
    }

    // Initialize the database connection
    db, _ := config.InitDB()

    // Insert the new person into the database
    if err := db.Create(&request).Error; err != nil {
        c.JSON(500, gin.H{
            "success": false,
            "message": "Error while adding the person",
            "error":   err.Error(),
        })
        return
    }

    c.JSON(200, gin.H{
        "success": true,
        "message": "Person added successfully",
        "data":    request,
    })
}


func DeletePersonByID(c *gin.Context) {
    // Get the person ID from the URL parameters
    personID := c.Param("id")

    // Initialize the database connection
    db, _ := config.InitDB()

    // Delete the person by ID
    if err := db.Where("id = ?", personID).Delete(&models.Person{}).Error; err != nil {
        c.JSON(500, gin.H{
            "success": false,
            "message": "Error while deleting the person",
            "error":   err.Error(),
        })
        return
    }

    c.JSON(200, gin.H{
        "success": true,
        "message": "Person deleted successfully",
    })
}


func GetPersonByID(c *gin.Context) {
    // Get the person ID from the URL parameters
    personID := c.Param("id")

    // Query the database to find the person by ID
    var person models.Person
    db, _ := config.InitDB()
    if err := db.Where("id = ?", personID).First(&person).Error; err != nil {
        c.JSON(404, gin.H{
            "success": false,
            "message": "Person not found",
        })
        return
    }

    c.JSON(200, gin.H{
        "success": true,
        "message": "Person found",
        "data":    person,
    })
}


func GetPersons(c *gin.Context) {
    // Query the database to fetch all persons
    var persons []models.Person // Assuming you have a Person model
    db, _ := config.InitDB()
    if err := db.Find(&persons).Error; err != nil {
        c.JSON(500, gin.H{
            "success": false,
            "message": "Error while fetching data",
            "error":   err.Error(),
        })
        return
    }

    c.JSON(200, gin.H{
        "success": true,
        "message": "Data fetched successfully",
        "data":    persons,
    })
}
