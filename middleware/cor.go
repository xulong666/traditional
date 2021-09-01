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


import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")  // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}