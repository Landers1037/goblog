/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package models

type Uv struct {
	Model

	Count int `json:"count"`
	Name string `json:"name"`

}

func GetuvFromdb() (uv Uv) {
	db.First(&uv,1)
	return
}

func AddUv(num int) {
	var uv Uv
	//每5个请求为一个缓存
	db.Select("count").First(&uv)
	num = num+uv.Count
	db.First(&uv,1).Update("count",num)
	return
}