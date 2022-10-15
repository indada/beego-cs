package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id      int `orm:"auto"`
	Age     int
	Name    string `orm:"column(user_name)"`
	Address string
}

func init() {
	//数据库
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/beeone?charset=utf8mb4")
	orm.RegisterModel(new(User))
}
func (this *User) TableName() string {
	return "user"
}
