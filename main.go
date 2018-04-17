package main

import (
	"flag"
	"fmt"
	"wechatHimsAPI/api"
	"wechatHimsAPI/config"
	"wechatHimsAPI/lib"
)

var ENV *string = flag.String("d", "development", "Enviorment development staging production")

func main() {
	flag.Parse()
	config.LoadConfig(*ENV)
	lib.InitDB()
	lib.InitRedisClient()

	fmt.Println("["+*ENV+"] server running  in "+*ENV+"", " enviorment.")
	api.Router.Run(config.C["port"])
}
