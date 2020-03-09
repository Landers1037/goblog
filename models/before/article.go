/*
Author: Landers
Github: Landers1037
Date: 2020-03
Name: cloudp
*/
package models

type Posts struct {
	Model
	Title string `json:"title"`
	Content string `json:"content"`
	Url string `json:"url"`
}
type Post struct {
	Model
	
	Title string `json:"title"`
	Date string `json:"date"`
	Content string `json:"content"`
	Url string `json:"url"`
}

func Getarticles() (articles []Posts)  {
	db.Model(&Posts{}).Find(&articles)
	return
}

func Getarticle(name string) (article Post)  {
	db.Model(&Post{}).Where("url = ?",name).First(&article)
	return
}

func Getbrother(name string) (p ,n string)  {
	var a Posts
	var prev Posts
	var next Posts
	db.Where("url = ?",name).First(&a)
	db.First(&prev,a.ID-1)
	p = prev.Url
	db.First(&next,a.ID+1)
	n = next.Url

	return p,n
}