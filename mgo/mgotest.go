package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("persons").C("persons")

	fmt.Println(c)

	err = c.Insert(&Person{"fuck", "185-1166-9684"}, &Person{"ttttt", "158-5900-9096"})

	if err != nil {
		fmt.Println(err)
	}
}
