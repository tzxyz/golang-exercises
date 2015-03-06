package main

import "fmt"

type Parent struct {
	Name string
}

type Child struct {
	Parent
}

func (p *Parent) Foo() {
	fmt.Println("parent foo")
}
func (p *Parent) Bar() {
	fmt.Println("parent bar")
}
func (c *Child) Bar() {
	c.Par.Bar()
	fmt.Println("child bar")
}

func main() {
	p := &Parent{Name: "vincent"}
	c := &Child{Par: p}
	c.Bar()
	c.Par.Bar()
	c.Foo()
}
