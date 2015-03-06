package main
import "fmt"

type Connecter interface{
	Connect()
}

type UsbConnecter interface{
	GetUsbName() string
	Connecter
}

type Usb struct{
	Name string
	Connected bool
}

func (u Usb) Connect() {
	u.Connected = true
	fmt.Println(u.GetUsbName()," is connect ... :", u.Connected)
}

func (u Usb) GetUsbName() string {
	return u.Name
}

/**
 *	接口查询：
 *	接口引用的结构体类型 是/否 := 对象.(接口引用的结构体类型)
 */
func Disconnect(usbConnecter UsbConnecter) {
	//判断usbConnecter是不是Usb类型，是的话返回两个值，一个是Usb，一个是true
	if usb,ok:=usbConnecter.(Usb); ok {
		fmt.Println(usb.Name," is disconnect...")
		return
	}
	fmt.Println("Unknown device")
}

/**
 *	根据接口引用的类型的不同，选择执行的方法
 */
func DisconnectByType(usbConnecter UsbConnecter) {
	switch v:=usbConnecter.(type){
	case Usb:
		fmt.Println(v.Name," is disconnect ...")
	default:
		fmt.Println("Unknown device")
	} 
}

func main() {
	var usb Usb = Usb{Name:"Usb",Connected:false}
	var usbConnecter UsbConnecter = usb
	var connecter Connecter = usbConnecter

	connecter.Connect()
	fmt.Println(usb.Connected)	//需要通过什么方式传进来

	Disconnect(usbConnecter)
	DisconnectByType(usbConnecter)
}