package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

const cow = ` 
____________________ 
< Hello,C >
 -------------------- 
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
`

func init() {
	Router = gin.Default()

	// Interceptor (access_token)
	Router.Use(func(c *gin.Context) {
		token := c.GetHeader("access_token")
		fmt.Println("======= access_token: " + token + " =======")
	})

	// index
	Router.GET("/", func(c *gin.Context) {
		c.String(200, cow)
	})
}
