package api

import (
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

	// index
	Router.GET("/", func(c *gin.Context) {
		c.String(200, cow)
	})
}
