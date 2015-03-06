package main

import "fmt"

type PersonInfo struct {
	Pid   string
	Name  string
	Phone string
}

func main() {
	/*
		1.声明一个map，并且对他赋值
		var personDB map[string]PersonInfo
		personDB = make(map[string]PersonInfo)
	*/

	personDB := make(map[string]PersonInfo)

	/*
		往map中插入两条数据
		也可以在创建的时候初始化
		personDB := make(map[string]PersonInfo){
			k1 : v1 ,
			k2 : v2 ...
		}
	*/
	personDB["001"] = PersonInfo{"001", "vincent", "xxx-xxxx-xxxx"}
	personDB["002"] = PersonInfo{"002", "focus", "xxx-xxxx-xxxx"}

	/*
		从map中查找出key为001的PersonInfo
	*/
	person, ok := personDB["001"]
	if ok {
		fmt.Println("person :", person, "ok :", ok)
	} else {
		fmt.Println("not found person")
	}

	/*
		遍历出map的元素
		map的range返回的两个值分别是k，v
	*/
	fmt.Println("删除前")
	for k, v := range personDB {
		fmt.Println("key: ", k, "--> value: ", v)
	}

	/*
		删除容器中指定的元素
	*/
	delete(personDB, "001")
	fmt.Println("删除后")
	for k, v := range personDB {
		fmt.Println("key: ", k, "--> value: ", v)
	}
}
