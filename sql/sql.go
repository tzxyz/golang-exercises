package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

func main() {
	db, err := sql.Open("mysql", "root:pass@tcp(192.168.0.209:3306)/test?charset=utf8")
	defer db.Close()
	// fmt.Println(reflect.TypeOf(db))
	chkErr(err)
	
	// fmt.Println(db)

	// Insert(db)
	// Update(db)
	// Delete(db)
	Query(db)
}

func Query(db *sql.DB) {
	rows, err := db.Query("select * from user")
	chkErr(err)
	for rows.Next() {
		var uid int
		var name string
		var deptname string
		var created string
		err = rows.Scan(&uid, &name, &deptname, &created)
		fmt.Println("###################")
		fmt.Println("uid：", uid)
		fmt.Println("name：", name)
		fmt.Println("deptname：", deptname)
		fmt.Println("created：", created)
	}

}

func Delete(db *sql.DB){
	stmt, err := db.Prepare("delete from user where uid = ?")
	chkErr(err)
	res, err := stmt.Exec("3")
	chkErr(err)
	affect, err := res.RowsAffected()
	fmt.Println(affect)
}

func Update(db *sql.DB){
	stmt, err := db.Prepare("update user set name = ? where uid = ?")
	// defer stmt.Close()
	chkErr(err)
	res, err := stmt.Exec("fuck", "2")
	chkErr(err)
	affect, err := res.RowsAffected()
	chkErr(err)
	fmt.Println(affect)		//影响的行数
}

func Insert(db *sql.DB){
	//插入数据			为什么不能插入中文
	stmt, err := db.Prepare("insert into user (name, deptname, created) values(?, ?, ?)")
	defer stmt.Close()
	chkErr(err)
	res, err := stmt.Exec("zhuonima", "yanfa", "2013-01-12")

	chkErr(err)

	id, err := res.LastInsertId()
	chkErr(err)

	fmt.Println(id)
}

func chkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	//panic(err.String())
}
