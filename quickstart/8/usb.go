package main

import "fmt"

type Usb interface {
	GetName() string
	Connect()
}

type PhoneConnecter struct {
	Name      string
	Connected bool
}

func (pc PhoneConnecter) GetName() string {
	return pc.Name
}

func (pc PhoneConnecter) Connect() {
	pc.Connected = true
	fmt.Println("PhoneConnecter is connect... :", pc.Connected)
}

func main() {
	var p PhoneConnecter = PhoneConnecter{Name: "Phone", Connected: false}
	var u Usb = p
	fmt.Println(u.GetName())
	u.Connect()
	fmt.Println(p.Connected)	//这里要怎么解决，Connect不会变为true
}
