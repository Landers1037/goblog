/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package sys

import (
	"cloudp/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"runtime"
)

func Getroutines(c * gin.Context)  {
	n := getGID()

	c.JSON(http.StatusOK,n)
}

func Getmem(c *gin.Context)  {
	o := sh("free")

	c.JSON(http.StatusOK,o)
}

func Getuv(c *gin.Context)  {
	uv :=  models.GetuvFromdb()
	c.JSON(http.StatusOK,uv)
}

func Getcount(c *gin.Context)  {
	count := count()
	c.JSON(http.StatusOK,count)
}
//内部函数
func getGID() int {
	n := runtime.NumGoroutine()
	return n
}

func sh(cmd string) string {
	cmdout := exec.Command(cmd)
	op,err:= cmdout.Output()
	if err!=nil{
		fmt.Println(err)
		return "getmem failed"
	}
	return string(op)
}


func count() string{

	return "null"
}