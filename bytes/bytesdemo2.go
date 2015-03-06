package main

import (
	"bytes"
	"fmt"
)

func main() {
	// NewBufferDemo()
	// NewBufferStringDemo()
	// StringDemo()
	// TruncateDemo()
	// LenDemo()
	// NextDemo()
	// BytesDemo()
	ResetDemo()
}

// 8.func (b *Buffer)
// 重置b.buf,清空buf里面的字节
// 实现b.Trancate(0)
func ResetDemo() {
	str := "reset demo"
	buf := bytes.NewBufferString(str)

	fmt.Println("reset befor:", buf) //Output: "reset demo"

	buf.Reset()

	fmt.Println("reset after:", buf) //Output: ""
}

// 7.func (b *Buffer) []byte
// 返回未读的字节切片
// 实现: return b.buf[b.off:]
func BytesDemo() {
	str := "(#￣▽￣#)"
	buf := bytes.NewBufferString(str)
	fmt.Println("bytes:", buf) //Output: "(#￣▽￣#)"
	buf.Next(2)
	fmt.Println("bytes:", buf) //Output: ￣▽￣#)"
}

// 6.func (b *Buffer) Next(n int)
func NextDemo() {

}

// 5.func (b *Buffer) Len() int
// 实现:return len(b.buf) - b.off
// 返回的是未读的字节数
func LenDemo() {
	b := []byte("~~~ ○ (*￣︶￣*）○ ~~~")
	buf := bytes.NewBuffer(b)

	fmt.Printf("read before: len = %d, buf = %s\n", buf.Len(), buf)
	//Output: "read before: len = 30, buf = ~~~ ○ (*￣︶￣*）○ ~~~"
	buf.Next(3)
	fmt.Printf("read once: len = %d, buf = %s\n", buf.Len(), buf)
	//Output: "read before: len = 27, buf =  ○ (*￣︶￣*）○ ~~~"

	buf.Next(23)
	fmt.Printf("read two: len = %d, buf = %s\n", buf.Len(), buf)
	//Output: "read before: len = 3, buf = ~~~"
}

// 4.func (b *Buffer) Truncate(n int)
// 没有返回值
// 底层只是修改了b.off b.buf = b.buf[0 : b.off+n](仅保留未读的部分+n个字节)
func TruncateDemo() {
	b := []byte("uuuucccc")
	buf := bytes.NewBuffer(b)

	fmt.Printf("before truncate, buf = %s\n", buf)
	//Output: "before truncate, buf = uuuucccc"

	buf.Truncate(4)

	fmt.Printf("after truncate, buf = %s\n", buf)
	//Output: "after truncate buf = uuuu"
}

// 3.func (b *Buffer) String() string
// 如果b==nil,返回字符串<nil>
// 如果b!=nil,返回字符串string(b.buf[0:b.off])
// ps:即返回b已经读过的字节转换成的字符串
func StringDemo() {
	b := []byte("ffffffff")
	buf := bytes.NewBuffer(b)

	s := "kkkkkkkk"
	bufs := bytes.NewBufferString(s)

	fmt.Printf("buf还有%d个字节未读,buf=%s\n", buf.Len(), buf.String())
	fmt.Printf("bufs还有%d个字节未读,bufs=%s\n", bufs.Len(), bufs.String())

	buf.Truncate(3) //保留bytes[0:b.off+n]
	bufs.Truncate(3)

	fmt.Printf("buf还有%d个字节未读,buf=%s\n", buf.Len(), buf.String())
	fmt.Printf("bufs还有%d个字节未读,bufs=%s\n", bufs.Len(), bufs.String())
}

// 2.bytes.NewBufferString(s string) *Buffer
// 创建和初始化一个Buffer,使用s作为初始化内容
// 返回一个*Buffer对象
// 实际上和NewBuffer一样
func NewBufferStringDemo() {
	s1 := "fuck my life"
	bufs1 := bytes.NewBufferString(s1)

	s2 := "法克买赖扶"
	bufs2 := bytes.NewBufferString(s2)

	fmt.Println(bufs1) //Output: "fuck my life"
	fmt.Println(bufs2) //Output: "法克买赖扶"
}

// 1.bytes.NewBuffer(buf []byte) *Buffer
// 创建和初始化一个Buffer,使用buf作为初始化内容
// return &Buffer{buf:buf}
// 返回一个*Buffer对象
func NewBufferDemo() {
	b1 := []byte("abcdefg")
	buf1 := bytes.NewBuffer(b1)

	b2 := []byte{}
	buf2 := bytes.NewBuffer(b2)

	fmt.Println("buf1=", buf1) //Output: "abcdefg"
	fmt.Println("buf2=", buf2) //Output: ""
}
