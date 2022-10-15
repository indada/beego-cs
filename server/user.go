package server

import (
	"bee1/models"
	"context"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type UserServer interface {
	Add() (int64, error)
	WorkAdd() error
	WorkAdds() error
	Read() error
	ReadAge() error
}
type users struct {
	UserModel *models.User
}
type product struct {
	models.User
	ProductName string `orm:"column(product_name)"`
}

func (u *users) Add() (int64, error) {
	o := orm.NewOrm()
	fmt.Println("AddUserModel:", u.UserModel)
	id, err := o.Insert(u.UserModel)
	return id, err
}
func (u *users) WorkAdd() (err error) {
	//InnoDB
	o := orm.NewOrm()
	to, err := o.Begin()
	if err != nil {
		logs.Error("start the transaction failed")
		return
	}
	fmt.Println("WorkAdd:", u.UserModel)
	id, err := to.Insert(u.UserModel)
	err = to.Rollback()
	if err != nil {
		logs.Error("roll back transaction failed", err)
	}
	fmt.Println("id:", id)
	return
	if err != nil {
		logs.Error("execute transaction's sql fail, rollback.", err)
		err = to.Rollback()
		if err != nil {
			logs.Error("roll back transaction failed", err)
		}
		return
	}
	err = to.Commit()
	if err != nil {
		logs.Error("commit transaction failed.", err)
	}
	return
}
func (u *users) WorkAdds() (err error) {
	//闭包方式
	o := orm.NewOrm()
	err = o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		id, e := txOrm.Insert(u.UserModel)
		fmt.Println("id:", id)
		//e = errors.New("主动错误！")
		return e
	})
	return
}
func (u *users) Read() error {
	o := orm.NewOrm()
	fmt.Println("ReadUserModel-qian:", u.UserModel)
	err := o.Read(u.UserModel)
	fmt.Println("ReadUserModel-hou:", u.UserModel)
	return err
}

// 指定字段查询
func (u *users) ReadAge() error {
	o := orm.NewOrm()
	fmt.Println("ReadAgeUserModel:", u.UserModel)
	err := o.Read(u.UserModel, "name")
	return err
}
func MultiAdd(bulk int, userList []models.User) (n int64, err error) {
	o := orm.NewOrm()
	n, err = o.InsertMulti(bulk, userList)
	return
}
func Update(user *models.User) (n int64, err error) {
	o := orm.NewOrm()
	n, err = o.Update(user)
	return
}

// 主键删除
func Delete(user *models.User) (n int64, err error) {
	o := orm.NewOrm()
	n, err = o.Delete(user)
	return
}

func Builder() []product {
	qb, _ := orm.NewQueryBuilder("mysql")
	users := make([]product, 10)
	qb.Select("user.user_name", "user.id", "user.age",
		"user_product.product_name").
		From("user").
		InnerJoin("user_product").On("user.id = user_product.user_id").
		Where("age > ?").
		OrderBy("user_name").Desc().
		Limit(10).Offset(0)

	// 导出 SQL 语句
	sql := qb.String()
	println("sql:", sql)
	// 执行 SQL 语句
	o := orm.NewOrm()
	o.Raw(sql, 120).QueryRows(&users)
	println(users)
	return users
}
func Seter() {
	o := orm.NewOrm()

	// 获取 QuerySeter 对象，user 为表名
	qs := o.QueryTable("user")

	// 也可以直接使用 Model 结构体作为表名
	qs = o.QueryTable(&models.User{})

	// 也可以直接使用对象作为表名
	user := new(models.User)
	qs = o.QueryTable(user)              // 返回 QuerySeter
	qs.Filter("id", 55)                  // WHERE id = 1
	qs.Filter("user_product__id", 4)     // WHERE user_product.id = 4
	qs.Filter("user_product__id__gt", 4) // WHERE user_product.id > 4 gte>=,lt<,lte<=
	/*exact / iexact 等于
	contains / icontains 包含
	gt / gte 大于 / 大于等于
	lt / lte 小于 / 小于等于
	startswith / istartswith 以...起始
	endswith / iendswith 以...结束
	in
	isnull
	后面以 i 开头的表示：大小写不敏感*/
	// 后面可以调用qs上的方法，执行复杂查询
	qs.Filter("user_name__icontains", "55")
	// WHERE user_name LIKE '%5%'

	// 原子操作增加字段
	num, err := o.QueryTable("user").Filter("id", 55).Update(orm.Params{
		"age": orm.ColValue(orm.ColAdd, 100),
	})
	println(num, err)
	// SET nums = nums + 100
	/*ColAdd      // 加
	ColMinus    // 减
	ColMultiply // 乘
	ColExcept   // 除*/
}

var _ UserServer = &users{}

func NewUserServer(ser *models.User) UserServer {
	return &users{
		UserModel: ser,
	}
}
