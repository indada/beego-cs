package controllers

import (
	"bee1/models"
	"bee1/server"
	"encoding/json"
	"fmt"
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
	c.Ctx.WriteString("getUserId")
}
func (c *UserController) He() {
	name := c.GetString("name", "default")
	if name == "" {
		c.Ctx.WriteString("Hello World")
		return
	}
	c.Ctx.WriteString("Hello " + name)
}
func (c *UserController) Hjson() {
	var use models.User
	body := c.Ctx.Input.RequestBody
	bb := string(body)
	fmt.Println("name:", c.GetString("name"), "  bb:", bb)
	err := json.Unmarshal(body, &use)
	if err != nil {
		fmt.Println(err)
	}
	dd := make(map[string]interface{})
	fmt.Println(use.Name)
	dd["name"] = use.Name
	dd["age"] = use.Age
	c.Data["json"] = dd
	c.ServeJSON()
}

type Request struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

func (c *UserController) Add() {
	request := Request{Code: 201, Msg: "请求失败！", Data: map[string]interface{}{}}
	var use models.User
	body := c.Ctx.Input.RequestBody
	bb := string(body)
	fmt.Println("name:", c.GetString("name"), "  bb:", bb)
	err := json.Unmarshal(body, &use)
	if err != nil {
		request.Code = 201
		request.Msg = err.Error()
		fmt.Println(err)
		c.Data["json"] = request
		c.ServeJSON()
		return
	}
	userServer := server.NewUserServer(&use)
	id, err := userServer.Add()
	fmt.Println("id:", id, "err:", err)
	if err != nil {
		fmt.Println("err:", err)
		request.Code = 201
		request.Msg = err.Error()
	} else {
		request.Data["id"] = id
		request.Code = 200
		request.Msg = "add数据成功！"
	}

	fmt.Println("request:", request)
	c.Data["json"] = request
	c.ServeJSON()
}

// 批量添加
func (c *UserController) MultiAdd() {
	request := Request{Code: 201, Msg: "请求失败！", Data: map[string]interface{}{}}
	usrList := make([]models.User, 6)
	body := c.Ctx.Input.RequestBody
	bb := string(body)
	fmt.Println("MultiAdd:", "  bb:", bb)
	err := json.Unmarshal(body, &usrList)
	if err != nil {
		request.Code = 201
		request.Msg = err.Error()
		fmt.Println(err)
		c.Data["json"] = request
		c.ServeJSON()
		return
	}
	n, err := server.MultiAdd(len(usrList), usrList)
	fmt.Println("添加了:", n, "err:", err)
	if err != nil {
		fmt.Println("err:", err)
		request.Code = 201
		request.Msg = err.Error()
	} else {
		request.Data["number"] = n
		request.Code = 200
		request.Msg = "adds数据成功！"
	}

	fmt.Println("request:", request)
	c.Data["json"] = request
	c.ServeJSON()
}

func (c *UserController) Update() {
	request := Request{Code: 201, Msg: "请求失败！", Data: map[string]interface{}{}}
	body := c.Ctx.Input.RequestBody
	use := models.User{}
	err := json.Unmarshal(body, &use)
	if err != nil {
		request.Code = 201
		request.Msg = err.Error()
		fmt.Println(err)
		c.Data["json"] = request
		c.ServeJSON()
		return
	}
	fmt.Println("use:", use)
	n, err := server.Update(&use)
	fmt.Println("修改n:", n, "err:", err)
	if err != nil {
		fmt.Println("err:", err)
		request.Code = 201
		request.Msg = err.Error()
	} else {
		request.Data["number"] = n
		request.Code = 200
		request.Msg = "up数据成功！"
	}

	fmt.Println("request:", request)
	c.Data["json"] = request
	c.ServeJSON()
}
func (c *UserController) Delete() {
	request := Request{Code: 201, Msg: "请求失败！", Data: map[string]interface{}{}}
	body := c.Ctx.Input.RequestBody
	use := models.User{}
	err := json.Unmarshal(body, &use)
	if err != nil {
		request.Code = 201
		request.Msg = err.Error()
		fmt.Println(err)
		c.Data["json"] = request
		c.ServeJSON()
		return
	}
	fmt.Println("use:", use)
	n, err := server.Delete(&use)
	fmt.Println("删除n:", n, "err:", err)
	if err != nil {
		fmt.Println("err:", err)
		request.Code = 201
		request.Msg = err.Error()
	} else {
		request.Data["number"] = n
		request.Code = 200
		request.Msg = "De数据成功！"
	}

	fmt.Println("request:", request)
	c.Data["json"] = request
	c.ServeJSON()
}
func (c *UserController) Read() {
	request := Request{Code: 201, Msg: "请求失败！", Data: map[string]interface{}{}}
	body := c.Ctx.Input.RequestBody
	use := models.User{}
	err := json.Unmarshal(body, &use)
	if err != nil {
		request.Code = 201
		request.Msg = err.Error()
		fmt.Println(err)
		c.Data["json"] = request
		c.ServeJSON()
		return
	}
	fmt.Println("use:", use)
	userServer := server.NewUserServer(&use)

	/*
		u = &User{Id: user.Id}
		err = Ormer.Read(u)
		// 只读取用户名这一个列
		u = &User{}
		err = Ormer.Read(u, "UserName")*/

	//id查询
	/*err = userServer.Read()
	fmt.Println("Read-err:", err)*/
	err = userServer.ReadAge()
	fmt.Println("ReadAge-err:", err)
	if err != nil {
		fmt.Println("err:", err)
		request.Code = 201
		request.Msg = err.Error()
	} else {
		request.Code = 200
		request.Msg = "Read数据成功！"
		request.Data["userInfo"] = use
	}

	fmt.Println("request:", request)
	c.Data["json"] = request
	c.ServeJSON()
}
func (c *UserController) WorkAdd() {
	//事务添加
	request := Request{Code: 201, Msg: "请求失败！", Data: map[string]interface{}{}}
	var use models.User
	body := c.Ctx.Input.RequestBody
	bb := string(body)
	fmt.Println("name:", c.GetString("name"), "  bb:", bb)
	err := json.Unmarshal(body, &use)
	if err != nil {
		request.Code = 201
		request.Msg = err.Error()
		fmt.Println(err)
		c.Data["json"] = request
		c.ServeJSON()
		return
	}
	userServer := server.NewUserServer(&use)
	err = userServer.WorkAdds()
	fmt.Println("err:", err)
	if err != nil {
		fmt.Println("err:", err)
		request.Code = 201
		request.Msg = err.Error()
	} else {
		request.Code = 200
		request.Msg = "add数据成功！"
	}

	fmt.Println("request:", request)
	c.Data["json"] = request
	c.ServeJSON()
}

func (c *UserController) Builder() {
	request := Request{Code: 201, Msg: "请求失败！", Data: map[string]interface{}{}}
	request.Data["list"] = server.Builder()
	request.Code = 200
	request.Msg = "获取数据成功！"
	c.Data["json"] = request
	c.ServeJSON()
}
