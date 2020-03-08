/*
Author: Landers
Github: Landers1037
Date: 2020-02
Name: cloudp
*/
package settings

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File
	RunMode string

	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	PageSize int
	JwtSecret string
) //一次定义多个变量

func init()  {
	var err error
	Cfg,err = ini.Load("conf/app.ini")
	if err!=nil{
		log.Fatal(2, "Fail to parse 'conf/app.ini': %v", err)
	}//出现错误时执行
	loadMode()
	loadServer()
	loadApp()
	
}

func loadMode()  {
	//决定debug模式的开关
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	//muststring 当配置不存在时所返回的默认值
}

func loadServer()  {
	//服务器相关配置
	server, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal(2, "Fail to get section 'server': %v", err)
	}
	HTTPPort = server.Key("HTTP_PORT").MustInt(6000)
	ReadTimeout = time.Duration(server.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout =  time.Duration(server.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadApp() {
	//应用的配置
	app, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatal(2, "Fail to get section 'app': %v", err)
	}

	JwtSecret = app.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = app.Key("PAGE_SIZE").MustInt(10)
}