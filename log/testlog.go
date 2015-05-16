package main

import (
	"log"
	"os"
)

func main() {
	testPrint()
}

func testPrint() {
	logger := log.New(os.Stdout, "[info]-->", log.LstdFlags)
	logger.Print("ahduiashdu")
}

func testPrintf() {
	log.Printf("%s", "aaa")
}
