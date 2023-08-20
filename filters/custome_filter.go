package filters

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Custom middleware function
func Test_Middleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Add your middleware logic here
        if 5 == 5 {
            fmt.Println("Middleware executed")
            c.Next() // If the condition is met, proceed to the next handler
        } else {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
                "message": "Access denied",
            })
        }
    }
}

