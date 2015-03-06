package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

type User struct {
	uid      int
	name     string
	deptname string
	created  string
}

func NewUser() *User {
	return &User{}
}

func main() {
	db, err := sql.Open("mysql", "root:pass@tcp(192.168.0.209:3306)/test?charset=utf8")
	ChkErr(err)
	defer db.Close()
	Query(db)
}

func Query(db *sql.DB) {
	rows, err := db.Query("select * from user")
	ChkErr(err)
	for rows.Next() {
		u := NewUser()
		err = rows.Scan(&u.uid, &u.name, &u.deptname, &u.created)	//居然真的封装进去了
		ChkErr(err)
		fmt.Println(u)
	}
}

func ChkErr(err error) {
	if err != nil {
		panic(err)
	}
}
