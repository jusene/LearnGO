package main

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/client/orm/filter/bean"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID            int    `orm:"column(id)"`
	Name          string `orm:"column(name)"`
	Age           int    `default:"12"`
	AgeInOldStyle int    `orm:"default(13);bee()"`
}

func init() {
	// need to register models
	orm.RegisterModel(new(User))
	// need to register db driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// need to register default database
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/beego?charset=utf8")
}

func main() {
	builder := bean.NewDefaultValueFilterChainBuilder(nil, true, true)
	orm.AddGlobalFilterChain(builder.FilterChain)
	o := orm.NewOrm()
	_, _ = o.Insert(&User{
		ID:   1,
		Name: "Tom",
	})
}
