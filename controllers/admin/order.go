package admin

import (
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["Website"] = "beego.me"
	val, _ := config.String("email")
	c.Data["Email"] = val
	logs.Info("sss", val)
	c.TplName = "index.tpl"
}
func (c *UserController) GetUserById() {
	c.Ctx.WriteString("admin-getUserId")
}
