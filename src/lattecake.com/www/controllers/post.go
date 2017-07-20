package controllers

import (
	"strconv"
	"lattecake.com/www/models"
	"strings"
	"log"
	"github.com/russross/blackfriday"
)

type PostController struct {
	BaseController
}

func (c *PostController) Get() {

	c.TplName = "post/list.tpl"
}

func (c *PostController) Detail() {

	id := c.Ctx.Input.Param(":id")

	postId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.Abort("404")
	}

	result, err := (&models.Post{}).Find(postId)
	if err != nil {
		c.Abort("404")
	}

	//md := markdown.New(markdown.HTML(true), markdown.Nofollow(true))
	//content := string(md.RenderToString([]byte(result.Content)))

	//(&helper.GoMarkdown{}).Text(result.Content)

	content := string(blackfriday.MarkdownCommon([]byte(result.Content)))

	con := strings.Split(content, "\n")
	topic := strings.Join(con[1:len(con)-1], "")


	result.ReadNum = result.ReadNum + 1

	(&models.Post{}).Update(result)

	c.Data["Post"] = result
	c.Data["Content"] = topic

	c.TplName = "post/detail.tpl"
}

func (c *PostController) Popular(limit int) {

	popular, err := (&models.Post{}).Popular(limit)
	if err != nil {
		log.Println("POPULAR NOT DATA.")
	}

	c.Data["Popular"] = popular
	c.LayoutSections["PopularPosts"] = "post/popular.tpl"
}
