package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
)

const indent = "	"

type Result struct {
	XMLName  xml.Name `xml:"result"`
	Property string   `xml:"property,attr"`
	Column   string   `xml:"column,attr"`
}

type ReTest struct {
	raw    string
	effect string
	actual string
}

func main() {
	re := regexp.MustCompile("private\\s([a-zA-Z]\\w*)\\s([\\_\\$a-z]\\w*)")
	// fmt.Println(re)

	// result := &Result{
	// 	Property: "aaa",
	// }
	// line, _ := xml.Marshal(result)
	// fmt.Println(string(line))

	// strtest := "private String name;"

	// isMatch := re.MatchString(strtest)

	// res := re.FindStringSubmatch(strtest)

	// fmt.Println(isMatch)

	// fmt.Println(len(res))

	// fmt.Println(res)
	// fmt.Println(res[0])
	// fmt.Println(res[2])

	reTests := []ReTest{
		ReTest{
			raw:    "private Long id",
			effect: "id",
		},
		ReTest{
			raw:    "private String birthday",
			effect: "birthday",
		},
		ReTest{
			raw:    "private Date now",
			effect: "now",
		},
		ReTest{
			raw:    "private String _name",
			effect: "_name",
		},
		ReTest{
			raw:    "private Person $name",
			effect: "$name",
		},
	}

	// results := make([]Result, 0, 32)

	for i := 0; i < len(reTests); i++ {
		actual := re.FindStringSubmatch(reTests[i].raw)[2]
		if actual == reTests[i].effect {
			// fmt.Printf("success! effect data=(%s), actual data=(%s)\n", reTests[i].effect, actual)
			res := &Result{
				Property: actual,
			}
			// fmt.Println(res)
			data, _ := xml.MarshalIndent(res, indent, indent) //字段必须是导出的
			// data, err := xml.Marshal(res)
			// fmt.Println(err)
			// fmt.Println(string(data))
			os.Stdout.Write(data)
		}
	}
	fmt.Println()
}
