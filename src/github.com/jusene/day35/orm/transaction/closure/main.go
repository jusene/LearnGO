package main

import (
	"context"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `orm:"column(id)"`
	Name string `orm:"column(name)"`
}

func init() {
	// need to register models in init
	orm.RegisterModel(new(User))

	// need to register db driver
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// need to register default database
	orm.RegisterDataBase("default", "mysql", "beego:test@tcp(192.168.0.105:13306)/orm_test?charset=utf8")

	orm.RunSyncdb("default", true, false)
}

// Demo using closure to execute transactions
func main() {
	o := orm.NewOrm()

	// Beego will manage the transaction's lifecycle
	// if the @param task return error, the transaction will be rollback
	// or the transaction will be committed
	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		// data
		user := new(User)
		user.Name = "test_transaction"

		// insert data
		// Using txOrm to execute SQL
		_, e := txOrm.Insert(user)
		// if e != nil the transaction will be rollback
		// or it will be committed
		return e
	})

	if err != nil {
		logs.Error("execute insert failed", err)
	} else {
		logs.Info("execute insert success")
	}
}
