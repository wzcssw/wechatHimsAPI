package main

import (
	"flag"
	"wechatHimsAPI/api"
	"wechatHimsAPI/config"
	"wechatHimsAPI/lib"
)

var ENV *string = flag.String("d", "development", "Enviorment development staging production")

func init() {
	config.LoadConfig(*ENV)
	lib.InitDB()
	lib.InitRedisClient()
}

func main() {
	flag.Parse()
	api.Router.Run(config.C["port"])
}
