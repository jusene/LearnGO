## Go语言 数据库

Go语言没有内置的驱动支持任何的数据库，但是Go语言定义了database/sql接口。

### Mysql

```go
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func checERR(err error)  {
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
```

### SQLite

```go
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
```

#### postgresql

```go
package main

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/lib/pq"
)

func checkErr(err error) {
	panic(err)
}

func main() {
	db, err :=sql.Open("postgres", "user=test password=test dbname=test sslmode=disable")
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
```