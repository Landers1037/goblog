/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package models

type Tags struct {
	Model
	Tag  string `json:"tag"`
	Posturl string `json:"posturl"`
	Title string `json:"title"`
	Content string `json:"content"`

}

func Gettags() (tags []Tags) {
	//返回全部的tags
	db.Table("tags").Select("tag").Where("tag != ?","暂时没有标签").Find(&tags)
	//第一个元素
	//db.First(&Tag{})
	return
}

func Getarticle_bytag(tag string) (posts []Tags)  {
	db.Model(&Tags{}).Where("tag = ?",tag).Find(&posts)
	return
}
