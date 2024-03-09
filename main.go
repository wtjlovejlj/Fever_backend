package main

import (
	"Fever_backend/dao/mysql"
	"Fever_backend/logger"
	"Fever_backend/pkg/snowflake"
	"Fever_backend/routers"
	"Fever_backend/settings"
	"fmt"
)

func main() {

	//加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:#{err}\n")
		return
	}

	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	//注册路由
	r := routers.SetupRouter()
	err := r.Run(":8085")
	if err != nil {
		fmt.Printf("run server failed,err:%v\n", err)
		return
	}

}
