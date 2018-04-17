package model

import (
	"time"
	"wechatHimsAPI/lib"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID             uint
	Name           string
	Realname       string
	Phone          string
	Email          string
	PasswordDigest string `json:"-"`
	RoleID         string
	HospitalID     uint
	Hospital       Hospital
}

// 使用Bcrypt比对password
func (user *User) Auth(pwd string) *User {
	if lib.CheckPasswordHash(pwd, user.PasswordDigest) {
		return user
	} else {
		return nil
	}
}

func (user *User) SaveAccessToken() string {
	accessToken := lib.GetMd5(user.Name)
	// ps: key: wechat_hims_api:f0a185d9c948178ec108f2d50bed48c5  value: 4
	// 24*30*12: one year
	lib.RedisClient.Set(lib.KeyHead+accessToken, user.ID, time.Duration(24*30*12)*time.Hour)
	return accessToken
}

func GetUserByAccessToken(accessToken string) *User {
	stringCmd := lib.RedisClient.Get(lib.KeyHead + accessToken)
	if result := stringCmd.Val(); result == "" {
		return nil
	} else {
		user := User{}
		lib.DB.First(&user, result).Related(&user.Hospital)
		return &user
	}
}

func GetUserIDByAccessToken(accessToken string) string {
	stringCmd := lib.RedisClient.Get(lib.KeyHead + accessToken)
	return stringCmd.Val()
}

func GetUserByAuth(username, password string) *User {
	user := User{}
	lib.DB.Where("name = ?", username).First(&user)
	lib.DB.Model(&user).Related(&user.Hospital)
	return user.Auth(password)
}

func CurrentUser(c *gin.Context) *User {
	token := c.Request.Header.Get("access-token")
	return GetUserByAccessToken(token)
}
