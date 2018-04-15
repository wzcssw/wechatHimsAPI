package api

import (
	"log"
	"strings"
	"wechatHimsAPI/model"

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

	// Interceptor (access-token)
	Router.Use(func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "/users/authenticate") {
			c.Next()
		} else {
			token := c.Request.Header.Get("access-token")
			log.Println("[Token] access-token:", token, " userID:", model.GetUserIDByAccessToken(token)) // test
			if userID := model.GetUserIDByAccessToken(token); userID == "" {
				c.JSON(200, gin.H{"success": false, "msg": "access-token invalid"})
				c.Abort()
			} else {
				c.Next()
			}
		}
	})

	// index
	Router.GET("/", func(c *gin.Context) {
		c.String(200, cow)
	})
}
