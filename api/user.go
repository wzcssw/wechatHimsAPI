package api

import (
	"wechatHimsAPI/lib"
	"wechatHimsAPI/model"

	"github.com/gin-gonic/gin"
)

func init() {
	user := Router.Group("/users")

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
		user := &model.User{}
		lib.DB.Where("name = ?", username).First(user)
		user = user.Auth(password)
		if user == nil {
			result["success"] = false
			result["msg"] = "the username or password is incorrect"
		} else {
			result["success"] = true
			result["msg"] = "OK"
			result["user"] = user
			result["access_token"] = user.SaveAccessToken()
		}
		c.JSON(200, result)
	})
}
