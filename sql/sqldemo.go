package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// 1.打开连接
	db, _ := sql.Open("mysql", "root:123456@tcp(192.168.1.102:3306)/test?charset=utf8")

	// 2.创建预处理语句(更新)
	stmt, _ := db.Prepare("UPDATE user SET username = ?, password = ? WHERE uid = ?")

	// 3.执行预处理语句返回结果
	res, _ := stmt.Exec("哈哈哈", "qwe@ewq.hotmail.com", 1)

	last, _ := res.LastInsertId()
	affect, _ := res.RowsAffected()

	log.Println(last)
	log.Println(affect)
}
