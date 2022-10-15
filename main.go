package main

import (
	_ "bee1/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	beego.BConfig.CopyRequestBody = true
	beego.Run()
}
