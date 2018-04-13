package model

import (
	"wechatHimsAPI/lib"
)

const keyHead = "wechat_hims_api:"

type User struct {
	ID             uint
	Name           string
	RealName       string
	Phone          string
	Email          string
	PasswordDigest string
	RoleID         string
	HospitalID     uint
}

func (user *User) Auth(pwd string) *User {
	if lib.CheckPasswordHash(pwd, user.PasswordDigest) {
		return user
	} else {
		return nil
	}
}

func (user *User) SaveAccessToken() string {
	accessToken := lib.GetMd5(user.Name)
	// ps: key: wechat_hims_api:f0a185d9c948178ec108f2d50bed48c5   value: 4
	lib.RedisClient.Set(keyHead+accessToken, user.ID, 0)
	return accessToken
}

func GetUserByAccessToken(accessToken string) *User {
	stringCmd := lib.RedisClient.Get(keyHead + accessToken)
	if result := stringCmd.Val(); result == "" {
		return nil
	} else {
		user := User{}
		lib.DB.Where("id=" + result).First(&user)
		return &user
	}
}
