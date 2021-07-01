package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func checkErr(err error) {
	panic(err)
}

func main() {
	db, err := sql.Open("postgres", "user=test password=test dbname=test sslmode=disable")
	checkErr(err)

	// 插入数据
	stmt, err := db.Prepare("INSET INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)

	res, err := stmt.Exec("JUSENE", "DR", "2020-05-12")
	checkErr(err)

	// pg不支持LastInsertId, 没有类似mysql的自增id
	// id, err := res.LastInsetId()
	// checkErr(err)

	var LastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES ($1,$2,$3) returning uid;",
		"jusene", "dr", "2020-05-12").Scan(&LastInsertId)
	checkErr(err)
	fmt.Println(LastInsertId)

	// 更新数据
	stmt, err = db.Prepare("UPDATE userinfo SET username=$1 where uid=$2")
	checkErr(err)

	res, err = stmt.Exec("zgx", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("SELECT * from userinfo")
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
	stmt, err = db.Prepare("DELETE FROM userinfo where uid=$1")
	checkErr(err)

	res, err = stmt.Exec(1)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	db.Close()

}
