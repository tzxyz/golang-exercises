package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	Uid      int64
	Username string
	Password string
}

func main() {
	db, _ := sql.Open("mysql", "root:123456@tcp(192.168.1.102:3306)/test?charset=utf8")

	// 查看所有，question：无法传参
	// rows, err := db.Query("SELECT uid,username,password FROM user")

	// 查询一行，第一个匹配到的
	// question：查询多行匹配的怎么办
	row := db.QueryRow("SELECT * FROM user where username = ?", "卓尼玛")

	log.Println(row)

	var uid int64
	var username string
	var password string
	row.Scan(&uid, &username, &password)

	log.Println(uid)
	log.Println(username)
	log.Println(password)

	// for rows.Next() {
	// 	log.Println("--- query start ---")
	// 	var uid int
	// 	var username string
	// 	var password string
	// 	err = rows.Scan(&uid, &username, &password) //这里的实现应该是个光标
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	log.Println(uid)
	// 	log.Println(username)
	// 	log.Println(password)
	// 	log.Println("--- one row ---")
	// }
}
