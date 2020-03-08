/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package article
////文章列表api
import (
	"cloudp/models"
	"cloudp/pkg/err"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func Getarticle(c *gin.Context){
	//单个文章的获取，因为使用前端渲染的方式所以数据库里只保存文章的文件名称
	//文件名称唯一且等于本地的文件名
	name := c.Query("name")
	if name != ""{
		data := models.Getarticle(name)
		code := err.SUCCESS
		c.JSON(http.StatusOK,gin.H{
			"code": code,
			"msg": err.GetMsg(code),
			"data": data,
		})
	}else {
		cwd,_ := os.Getwd()
		c.File(cwd+"/posts/404.md ")
	}
	//文章的获取应该是传输md文件流

}

func Getarticles(c *gin.Context)  {
	//获取文章的列表
	data := models.Getarticles()
	var length int = len(data)
	page := c.Query("p")
	if page!=""{
		p ,_:= strconv.Atoi(page)
		data = data[(p-1)*8:(p-1)*8+8]
		code := err.SUCCESS
		c.JSON(http.StatusOK,gin.H{
			"code" : code,
			"msg": err.GetMsg(code),
			"data": data,
			"len": length,

		})
	}else {
		code := err.SUCCESS
		c.JSON(http.StatusOK,gin.H{
			"code" : code,
			"msg": err.GetMsg(code),
			"data": data,

		})
	}

}

func Getarticle_bytag(c *gin.Context)  {
	//获取对应tag的文章
	tag := c.Query("tag")
	data := models.Getarticle_bytag(tag)
	code := err.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"code" : code,
		"msg": err.GetMsg(code),
		"data": data,

	})

}
