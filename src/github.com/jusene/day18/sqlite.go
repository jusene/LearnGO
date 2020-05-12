package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func checkErr(err error) {
	panic(err)
}

func main() {
	db, err := sql.Open("sqlite3", "./demo.db")
	checkErr(err)

	// 插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) VALUES (?,?,?)")
	checkErr(err)

	id, err := stmt.Exec("jusene", "dr", "2020-05-12")
	checkErr(err)

	fmt.Println(id)

	// 更新数据
	stmt, err = db.Prepare("UPDATE userinfo set username=? where uid=?")
	checkErr(err)

	res, err := stmt.Exec("zgx", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// 删除数据
	stmt, err = db.Prepare("DELETE FROM userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}
