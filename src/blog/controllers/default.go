package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "songgl"
	c.TplName = "index.tpl"
	c.Data["Desc"]="this is beego framework!"
}
