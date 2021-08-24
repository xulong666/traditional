package middleware

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )


// func Cors(c *gin.Context) {
	
// 		c.Header("Access-Control-Allow-Origin", "*")
// 		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
// 		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
// 		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
// 		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
// 		c.Header("Content-Type", "application/json")
// 		c.Header("Access-Control-Allow-Credentials", "true")
// 		if c.Request.Method != "OPTIONS" {
// 			c.AbortWithStatus(http.StatusNoContent)
// 		} 
// 		c.Next()
	
// }
