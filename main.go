package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"qiqi-go/conf"
	"qiqi-go/route"
)

func main() {
	fmt.Println("OSS Go SDK Version: ", oss.Version)
	// 初始化配置
	conf.Init()
	// 路由连接
	router := route.CollectRoute()
	//端口号启动在9000端口
	panic(router.Run(":9000"))
}
