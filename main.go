package main

import (
	"fmt"
	"time"
	"workwx/bootstrap"
	"workwx/config"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

var ch = make(chan int)

func main() {
	// 初始化 SQL
	bootstrap.SetupDB()
	// externalContact.CustomerDownloads()

	ticker := time.NewTicker(time.Second * 10)

	go func() {
		for range ticker.C {
			fmt.Println("ticked  at ", time.Now())
		}
	}()

	t2 := time.NewTicker(time.Second * 30)
	go func() {
		for range t2.C {
			fmt.Println("t2  at ", time.Now())
		}
	}()

	<-ch //阻塞主线程

	// gin.SetMode("release")
	// r := gin.New()

	// //fmt.Println(conf.GetString("app.storage"))
	// r.Use(logger.Logger(conf.GetString("app.storage")))
	// r.GET("/ping", func(c *gin.Context) {

	// })
	// r.Run(":9999") // 监听并在 0.0.0.0:8080 上启动服务
}
