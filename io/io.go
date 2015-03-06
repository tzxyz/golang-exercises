package main
import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
)
func main() {
	f1, _ := os.Open("fuck.txt")		//打开一个文件
	f2, _ := os.Create("fuck2.txt")		//创建一个文件

	num, err := io.Copy(f2,f1)			//io.Copy(dst, src)拷贝
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)

// ***********************************************************************************
	f3, _ := os.Open("fuck.txt")
	f4, _ := os.Create("fuck2.txt")
	num2, err := io.CopyN(f4, f3, 10)	//io.CopyN(dst, src, n int64) (written int64, err error) n:是读取限制，只能n个？
	if err != nil {
		fmt.Println(err)
	}
	defer f3.Close()
	rb, _ := ioutil.ReadFile("fuck2.txt")
	fmt.Println(string(rb))	//将字节数组转换成字符串
	fmt.Println(num2)	

//  ***********************************************************************************
//  io.ReadAtLeast(r Reader, b [] byte, min int)
	fmt.Println("*****************************************************************")
	r, _ := os.Open("fuck.txt")
	defer r.Close()
	b := make([]byte, 16)
	var total int 
	for {
		n, err := io.ReadAtLeast(r, b ,5)	//io.ReadAtLeast(r Reader, b byte[], min int)：将r中的内容读到b中，至少读取min个
		if err == nil {
			fmt.Println(string(b[:n]))	//使用切片的方式 将b中数据显示出来
		}
		if err == io.ErrUnexpectedEOF {		//读取的数据小于min
			fmt.Println("read to short", string(b[:n]))
		}
		if err == io.ErrShortBuffer {		//如果设置的buffer的length小于8
			fmt.Println("buffer to short")
		}
		if err == io.EOF {
			fmt.Println("bingo", total)
			break
		}
		total += n
	}

//  ***********************************************************************************
//  io.ReadFull(r Reader, b []byte)
	fmt.Println("****************************************************************")
	r2, err := os.Open("fuck.txt")
	defer r2.Close()
	b2 := make([]byte, 16)
	total = 0 
	for {
		n, err := io.ReadFull(r2, b2)
		if err == nil {
			fmt.Println(string(b2[:n]),"读取中")		//要把换行也算进去 一次读取buf.length个
		}
		if err == io.ErrUnexpectedEOF {
			fmt.Println(string(b2[:n]),"太短了")
		}
		if err == io.EOF {
			fmt.Println(string(b2[:n]), "一共读取了", total, "个字符")
			break
		}
		total += n
	}

//  **********************************************************************************
//  io.WriteString(s string)
	fmt.Println("****************************************************************")
	r3, err := os.OpenFile("d://fuck2.txt",os.O_RDWR,os.ModePerm)	//打开文件，拥有读写权限
	defer r3.Close()
	l, err := r3.WriteString("what do u want")
	if err != nil {
		fmt.Println(err)
	}	
	b3, _ := ioutil.ReadFile("d://fuck2.txt")		//并不是绝对路径的问题，是没有权限写入文件
	fmt.Println(string(b3[:l]))

// *******************************************************************************

}