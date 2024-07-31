package main

import (
	"github.com/Linxhhh/easy-doc/core"
	"github.com/Linxhhh/easy-doc/flags"
	"github.com/Linxhhh/easy-doc/global"
	"github.com/Linxhhh/easy-doc/routers"
)

func main() {

	/*
		global.Log = core.InitLogger(core.LogRequest{
			LogPath: "logs",
			AppName: "gvd",
		})
	*/
	global.Log = core.InitLogger()
	global.Config = core.InitConfig()
	global.DB = core.InitMysql()
	global.Redis = core.InitRedis(0)
	/*
		name, err := global.Reids.Get("name").Result()
	*/

	// 运行程序时，是否携带参数
	option := flags.Parse()
	if option.Run() {
		return
	}

	router := routers.Routers()

	addr := global.Config.System.Addr()
	router.Run(addr)
}
