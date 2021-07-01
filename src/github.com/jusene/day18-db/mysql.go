package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func checERR(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "test:test@tcp(192.168.88.137:3306)/go_mysql?charset=utf8")
	checERR(err)

	// 插入数据
	stmt, err := db.Prepare("INSERT userinfo set username=?, departname=?, created=?")
	checERR(err)

	res, err := stmt.Exec("JUSENE", "DR", "2020-5-11")
	checERR(err)

	id, err := res.LastInsertId()
	checERR(err)

	fmt.Println(id)

	// 更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checERR(err)

	res, err = stmt.Exec("zgx", id)
	checERR(err)

	affect, err := res.RowsAffected()
	checERR(err)

	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checERR(err)

	for rows.Next() {
		var id int
		var username string
		var departname string
		var created string
		err = rows.Scan(&id, &username, &departname, &created)
		checERR(err)
		fmt.Println(id)
		fmt.Println(username)
		fmt.Println(departname)
		fmt.Println(created)
	}

	// 删除数据
	stmt, err = db.Prepare("DELETE FROM userinfo where uid=?")
	checERR(err)

	res, err = stmt.Exec(id)
	checERR(err)

	affect, err = res.RowsAffected()
	checERR(err)

	fmt.Println(affect)

	// row sql
	res, err = db.Exec("DELETE FROM userinfo where uid=?", "5")
	checERR(err)

	affect, err = res.RowsAffected()
	checERR(err)

	fmt.Println(affect)

	db.Close()
}
