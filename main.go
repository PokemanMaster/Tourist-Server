package main

import (
	"qiqi-go/api"
	"qiqi-go/conf"
)

func main() {
	// 初始化配置
	conf.Init()
	// 路由连接
	router := api.CollectRoute()
	//端口号启动在9000端口
	panic(router.Run(":9000"))
}
