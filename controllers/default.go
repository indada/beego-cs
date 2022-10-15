package controllers

import (
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	val, _ := config.String("email")
	c.Data["Email"] = val
	logs.Info("auto load config name is", val)
	c.TplName = "index.tpl"
}
