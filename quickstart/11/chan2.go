package main
import "fmt"
import "runtime"
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c:= make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i," : go go go")
			c <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<- c
	}
}