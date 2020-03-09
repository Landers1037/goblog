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
	"cloudp/routers/api/robotTXT"
	"cloudp/routers/api/sys"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//中间件
	r.Use(gin.Logger())
	r.Use(middleware.Cors())
	r.Use(middleware.Uv())
	r.Use(middleware.SimpleAuth())
	r.Use(gin.Recovery())

	gin.SetMode(settings.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	//仅供测试

	robot := r.Group("/")
	{
		robot.GET("/robot.txt",robotTXT.GetRobot)
	}
	//文章api
	apiArticle := r.Group("/api/article")
	{
		apiArticle.GET("/tags", article.Gettags)         //获取全部标签
		apiArticle.GET("/tag", article.Getarticle_bytag) //获取对应标签下的文章
		apiArticle.GET("/posts", article.Getarticles)    //全部文章列表
		apiArticle.GET("/post", article.Getarticle)      //指定文章
		apiArticle.GET("/brother", article.Getbrother)   //指定文章前后篇
	}
	//系统参数api
	apiSys := r.Group("/api/sys")
	{
		apiSys.GET("/routines", sys.GetRoutines) //routine数目
		apiSys.GET("/mem", sys.GetMem) //占用内存大小
		apiSys.GET("/uv", sys.GetUv) // uv访问数
		apiSys.GET("/count", sys.GetCount) //ip当日访问总数，基于nginx日志
	}
	return r
}

