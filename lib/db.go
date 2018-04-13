package lib

import (
	"wechatHimsAPI/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 数据库连接
var DB *gorm.DB

func InitDB() {
	connectString := config.C["username"] + `:` + config.C["password"] + `@tcp(` + config.C["host"] + `:` + config.C["dbport"] + `)/` + config.C["database"] + `?charset=utf8&parseTime=true`
	DB, _ = gorm.Open("mysql", connectString)
}
