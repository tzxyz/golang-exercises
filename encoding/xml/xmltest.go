package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Names   []Name   `xml:"name"`
	Id      string   `xml:"name,attr"`
	Comment string   `xml:",comment"`
}

type Name string

func main() {
	names := make([]Name, 0, 8)

	var name1 Name = "111"
	var name2 Name = "222"

	names = append(names, name1)
	names = append(names, name2)

	person := &Person{
		Names:   names,
		Id:      "11",
		Comment: "注释",
	}

	indent := "    "

	data, err := xml.MarshalIndent(person, indent, indent)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
