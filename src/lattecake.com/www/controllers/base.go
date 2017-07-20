package controllers

import (
	"github.com/astaxie/beego"
	"lattecake.com/www/models"
	"log"
)

type BaseController struct {
	beego.Controller
}

func init() {

}

func (c *BaseController) Prepare() {


	popularCh := make(chan models.Posts)
	var popularPosts models.Posts

	go func() {
		posts, err := (&models.Post{}).Popular(5)
		if err != nil {
			log.Println(err)
		}
		popularCh <- posts
	}()

	popularPosts = <-popularCh

	close(popularCh)

	c.Data["PopularPosts"] = popularPosts
	c.Layout = "default_layout.tpl"
	c.LayoutSections = make(map[string]string)
	//c.LayoutSections["PopularPosts"] = "post/popular.tpl"
}