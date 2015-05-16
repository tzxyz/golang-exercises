package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type Config struct {
	Location string `json:"location"`
	Output   string `json:"output"`
}

func main() {
	// testMarshal()
	// testUnmarshal()
	testReadFile()
}

func testReadFile() {
	data, err := ioutil.ReadFile("User.java")
	if err != nil {
		fmt.Println(err)
	}

	// const idrule = "@Id\\s+private\\s[a-zA-Z]\\w*\\s([\\_\\$a-z]\\w*)"
	const idrule = "@Id"

	// 匹配出带@Column注解的所有字段
	const pattern = "@Column\\(.*\\)\\s+private\\s[a-zA-Z]\\w*\\s([\\_\\$a-z]\\w*);"

	re := regexp.MustCompile(pattern)

	text := string(data)

	// cols := re.FindAllString(text, -1)

	fields := re.FindAllStringSubmatch(text, -1)

	for _, field := range fields {
		fmt.Println(field[1])
	}

	re = regexp.MustCompile(idrule)

	ids := re.FindAllStringSubmatch(text, -1)
	fmt.Println(ids)

	// fmt.Println(len(cols))
	// fmt.Println(cols)

	// for _, col := range cols {
	// 	fmt.Println(col)
	// }
}

func testMarshal() {
	config := &Config{3
		Location: "f:/javaworkspaces/src/main/java/model/User.java",
		Output:   "f:/",
	}
	data, err := json.Marshal(config)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	// 格式
	var out bytes.Buffer
	json.Indent(&out, data, "", "\t")
	fmt.Println(string(out.Bytes()))

	file, err := os.Create("config.json")
	out.WriteTo(file)
}

func testUnmarshal() {
	data, err := ioutil.ReadFile("config.json")

	if err != nil {
		fmt.Println(err)
	}
	// testJson := `{"location":"f:/javaworkspaces/src/main/java/model/User.java","Output":"f:/"}`

	config := &Config{}

	err = json.Unmarshal(data, config)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config)
}
