package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Testt(c *gin.Context) {
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
