package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// 1.打开连接
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.1.102:3306)/test?charset=utf8")

	CheckErr(err)

	// log.Println(db)

	// 2.准备预处理语句
	stmt, err := db.Prepare("INSERT user SET username = ?, password = ?")

	CheckErr(err)

	// 3.执行语句
	res, err := stmt.Exec("卓尼玛", "123456")

	CheckErr(err)

	// 4.driver.Result定义了两个接口，返回自增最后的id，和影响的条数
	affect, _ := res.RowsAffected()
	lastid, _ := res.LastInsertId()

	log.Println(affect, lastid)
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
