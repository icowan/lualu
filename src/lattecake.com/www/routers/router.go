package routers

import (
	"lattecake.com/www/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/about", &controllers.DefaultController{}, "GET:About")
	beego.Router("/contact", &controllers.DefaultController{}, "GET:Contact")
	beego.Router("/post/:id", &controllers.PostController{}, "GET:Detail")
}
