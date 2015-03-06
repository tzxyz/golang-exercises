package main
import (
	"reflect"
	"fmt"
)

type Man struct {
	name string
	email string
}

func (m Man) Ins() string {
	return "My name is" + m.name + ",and My email is" + m.email
}

func main() {
	p := Man{"卓尼玛", "vincent.fml@hotmail.com"}

	// 反射一个类型的值，必须先将它转换成reflect.Type或reflect.Value的一个对象
	t := reflect.TypeOf(p)		//得到类型的元数据
	// v := reflect.ValueOf(p)		//得到类型的值

	//结构体的反射
	fmt.Println(t.Name())	//获取元数据类型的名字（Man）
	fmt.Println(t.Kind())	//获取元数据类型的分类
	fmt.Println(t.PkgPath())	//返回短包名
	fmt.Println(t.String())		//元数据类型全名？
	fmt.Println(t.NumField())	//返回结构体字段的总数
	fmt.Println(t.Field(0))		//返回结构体的第i个字段
	// fmt.Println(t.Elem())

	//方法的反射
	fmt.Println(t.NumMethod()) 	//返回该struct方法的个数(receiver不能使*T类型，不然不会统计)
	fmt.Println(t.Method(0))	//返回改struct的第n个方法
	fmt.Println(t.MethodByName("Ins"))	//根据方法名查找该struct的方法，返回method和是否找到方法
	
	//函数的反射
	f := func(str string) string{
		return str + "fuck"
	}

	t = reflect.TypeOf(f)
	fmt.Println(t.Name())	//匿名函数没有Name
	fmt.Println(t.Kind())	//返回的是func类型
	fmt.Println(t.NumIn())	//函数参数的个数
	fmt.Println(t.NumOut())	//函数返回值的个数
	fmt.Println(t.In(0))	//函数第n个参数的参数类型
	fmt.Println(t.Out(0))	//函数第n个返回值的返回值类型
}
