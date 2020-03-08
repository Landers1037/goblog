/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package routers

import (
	"cloudp/middleware"
	"cloudp/pkg/settings"
	"cloudp/routers/api/article"
	"cloudp/routers/api/sys"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(middleware.Cors())
	r.Use(middleware.Uv())

	r.Use(gin.Recovery())
	r.Use()

	gin.SetMode(settings.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	//仅供测试
	//文章api
	api_article := r.Group("/api/article")
	{
		api_article.GET("/tags", article.Gettags)         //获取全部标签
		api_article.GET("/tag", article.Getarticle_bytag) //获取对应标签下的文章
		api_article.GET("/posts", article.Getarticles)    //全部文章列表
		api_article.GET("/post", article.Getarticle)      //指定文章
		api_article.GET("/brother", article.Getbrother)      //指定文章
	}
	//系统参数api
	api_sys := r.Group("/api/sys")
	{
		api_sys.GET("/routines", sys.Getroutines)
		api_sys.GET("/mem", sys.Getmem)
		api_sys.GET("/uv", sys.Getuv)
		api_sys.GET("/count", sys.Getcount)
	}
	return r
}

