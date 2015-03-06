package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, _ := sql.Open("mysql", "root:123456@tcp(192.168.1.102:3306)/test?charset=utf8")

	stmt, _ := db.Prepare("DELETE FROM user where uid = ?")

	res, _ := stmt.Exec("2")

	affect, _ := res.RowsAffected()

	log.Println(affect)
	log.Println(last)
}
