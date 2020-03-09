package main

import (
	"cloudp/routers"
	"cloudp/utils/settings"
	"fmt"
	"net/http"
)

func main()  {
	//setting初始化的设置项导入

	//main api
	router := routers.InitRouter()
	//app.Run(":5000")
	//使用gin自带的web服务器启动
	//或者使用httpserver启动
	s := &http.Server{
		Addr: fmt.Sprintf(":%d",settings.HTTPPort),
		Handler: router,
		ReadTimeout: settings.ReadTimeout,
		WriteTimeout: settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err!=nil{
		fmt.Println(err)
	}
}

