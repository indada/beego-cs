package routers

import (
	"bee1/controllers"
	"bee1/controllers/admin"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	user := &controllers.UserController{}
	product := &admin.ProductController{}
	admin := &admin.UserController{}
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(user)
	/*beego.Router("/user", user)
	beego.Router("/users", user, "get,post:GetUserById")*/
	beego.Router("/admin/users", admin, "get,post:GetUserById")
	beego.AutoRouter(product)
}
