package api

import (
	"wechatHimsAPI/model"

	"github.com/gin-gonic/gin"
)

func init() {
	user := Router.Group("/users")

	user.GET("/current", func(c *gin.Context) {
		result := gin.H{}
		if user := model.CurrentUser(c); user == nil {
			result["success"] = false
		} else {
			result["success"] = true
			result["user"] = user
		}
		c.JSON(200, result)
	})

	user.POST("/authenticate", func(c *gin.Context) {
		result := gin.H{}
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "" || password == "" {
			result["success"] = false
			result["msg"] = "username or password is empty"
			c.JSON(200, result)
			return
		}
		user := model.GetUserByAuth(username, password)
		if user == nil {
			result["success"] = false
			result["msg"] = "the username or password is incorrect"
		} else {
			result["success"] = true
			result["msg"] = "OK"
			result["user"] = user
			result["access-token"] = user.SaveAccessToken()
		}
		c.JSON(200, result)
	})
}
