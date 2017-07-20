package main

import (
	_ "lattecake.com/www/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"lattecake.com/www/helper"
)

func main() {
	orm.Debug = true

	beego.AddFuncMap("replace_image_url", helper.ReplaceImageUrl)
	beego.AddFuncMap("replace_image_src", helper.ReplaceImageSrc)

	beego.Run()
}

