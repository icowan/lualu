package controllers


type DefaultController struct {
	BaseController
}

func (c *DefaultController) About() {
	c.TplName = "default/about.tpl"
}

func (c *DefaultController) Contact() {
	c.TplName = "default/contact.tpl"
}
