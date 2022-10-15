package admin

import (
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type ProductController struct {
	beego.Controller
}

func (c *ProductController) Get() {
	c.Data["Website"] = "admin-product"
	val, _ := config.String("email")
	c.Data["Email"] = val
	logs.Info("admin-product", val)
	c.TplName = "index.tpl"
}
func (c *ProductController) GetUserById() {
	c.Ctx.WriteString("admin-productController-getUserId")
}
