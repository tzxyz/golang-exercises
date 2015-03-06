package main
import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go Go(i)
	}
	time.Sleep(3*time.Second)
}

func Go(i int){
	fmt.Println(i," : go go go ...")
}