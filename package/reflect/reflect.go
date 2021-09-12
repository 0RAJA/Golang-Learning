package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//反射是指在程序运行期对程序本身进行访问和修改的能力
/*
	reflect.TypeOf 能获取类型信息；
	reflect.ValueOf 能获取数据的运行时表示；
*/
func main() {
	//反射三大法则
	if false {
		//第一法则:反射可以将接口类型变量转换为反射对象
		/*
				有了变量的类型之后，我们可以通过 Method 方法获得类型实现的方法，通过 Field 获取类型包含的全部字段。
			对于不同的类型，我们也可以调用不同的方法获取相关信息：

			结构体：获取字段的数量并通过下标和字段名获取字段 StructField；
			哈希表：获取哈希表的 Key 类型；
			函数或方法：获取入参和返回值的类型；
			…
		*/
		var x int64 = 4
		fmt.Println(reflect.TypeOf(x), reflect.ValueOf(x))
		//第二法则:反射可以把反射对象还原为接口对象
		/*
				不过调用 reflect.Value.Interface 方法只能获得 interface{} 类型的变量，
					如果想要将其还原成最原始的状态还需要经过如下所示的显式类型转换：

				v := reflect.ValueOf(1)
				v.Interface().(int)

			当然不是所有的变量都需要类型转换这一过程。
			如果变量本身就是 interface{} 类型的，那么它不需要类型转换，
			因为类型转换这一过程一般都是隐式的，所以我不太需要关心它，只有在我们需要将反射对象转换回基本类型时才需要显式的转换操作。
		*/
		var a interface{} = 4.0
		v := reflect.ValueOf(a) //反射对象
		b := v.Interface()      //接口对象
		fmt.Println(a == b)
		fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))
		//第三法则:反射对象可修改,value值必须是可设置的
		/*
			func main() {
				i := 1
				v := reflect.ValueOf(i)
				v.SetInt(10)
				fmt.Println(i)
			}

			$ go run reflect.go
			panic: reflect: reflect.flag.mustBeAssignable using unaddressable value

			由于 Go 语言的函数调用都是传值的，所以我们得到的反射对象跟最开始的变量没有任何关系，那么直接修改反射对象无法改变原始变量，程序为了防止错误就会崩溃。
			想要修改原变量只能使用如下的方法：
		*/
		i := 1
		p := reflect.ValueOf(&i)
		p.Elem().SetInt(20)
		fmt.Println(i)
	}
	//类型与值
	if false {
		/*
				当我们想要将一个变量转换成反射对象时，Go 语言会在编译期间完成类型转换，将变量的类型和值转换成了 interface{} 并等待运行期间使用 reflect 包获取接口中存储的信息。

			Go 语言的 interface{} 类型在语言内部是通过 reflect.emptyInterface 结体表示的

			//emptyInterface
			type emptyInterface struct {
				typ  *rtype	//表示变量的类型
				word unsafe.Pointer //指向内部封装的数据
			}

			//Typeof
			func TypeOf(i interface{}) Type {
				eface := *(*emptyInterface)(unsafe.Pointer(&i))
				return toType(eface.typ)
			}

			func toType(t *rtype) Type {
				if t == nil {
					return nil
				}
				return t
			}

			//ValueOf
			func ValueOf(i interface{}) Value {
				if i == nil {
					return Value{}
				}

				escapes(i)

				return unpackEface(i)
			}

			func unpackEface(i interface{}) Value {
				e := (*emptyInterface)(unsafe.Pointer(&i))
				t := e.typ
				if t == nil {
					return Value{}
				}
				f := flag(t.Kind())
				if ifaceIndir(t) {
					f |= flagIndir
				}
				return Value{t, e.word, f}
			}
		*/
	}
	//更新变量
	if false {
		/*
			当我们想要更新 reflect.Value 时，就需要调用 reflect.Value.Set 更新反射对象，
			该方法会调用 reflect.flag.mustBeAssignable 和 reflect.flag.mustBeExported
			分别检查当前反射对象是否是可以被设置的以及字段是否是对外公开的：

			func (v Value) Set(x Value) {
				v.mustBeAssignable()	//是否可以被设置
				x.mustBeExported()		//是否对外公开
				var target unsafe.Pointer
				if v.kind() == Interface {
					target = v.ptr
				}
				x = x.assignTo("reflect.Set", v.typ, target)
				typedmemmove(v.typ, v.ptr, x.ptr)
			}
		*/
	}
	//方法调用
	if false {
		/*
			1.通过 reflect.ValueOf 获取函数 Add 对应的反射对象；
			2.调用 reflect.rtype.NumIn 获取函数的入参个数；
			3.多次调用 reflect.ValueOf 函数逐一设置 argv 数组中的各个参数；
			4.调用反射对象 Add 的 reflect.Value.Call 方法并传入参数列表；
			5.获取返回值数组、验证数组的长度以及类型并打印其中的数据；
		*/
		v := reflect.ValueOf(Add) //反射对象
		if v.Kind() != reflect.Func {
			return
		}
		t := v.Type()
		argv := make([]reflect.Value, t.NumIn()) //入参个数
		for i := range argv {
			if t.In(i).Kind() != reflect.Int {
				return
			}
			argv[i] = reflect.ValueOf(i) //填充参数
		}
		result := v.Call(argv) //调用方法
		if len(result) != 1 || result[0].Kind() != reflect.Int {
			return
		}
		fmt.Println(result[0].Int()) // #=> 1
	}
	//结构体与反射
	if true {
		user := User{
			Id:   1,
			Name: "Raja",
			Age:  19,
		}
		Poni(user)
		//匿名字段
		fmt.Println("匿名字段:===================")
		NiM(Boy{
			User: user,
			Addr: "111",
		})
		//修改结构体值
		fmt.Println("修改结构体值:================")
		fmt.Println(user)
		SetValue(&user)
		fmt.Println(user)
		//调用方法
		fmt.Println("调用方法:==================")
		CallFunc(user)
		//获取字段
		fmt.Println("获取字段:==================")
		GetTag(user)
	}
}

type User struct {
	Id   int `json:"id"`
	Name string
	Age  int
}

func (u User) Hello(s string) {
	fmt.Println("Hello", u.Name, s)
}

func Poni(o interface{}) {
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	// 获取结构体字段个数：t.NumField()
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
		//获取字段的值信息
		//Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val:", val)
	}
	fmt.Println("方法:==================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}
}

func Add(a, b int) int {
	return a + b
}

type Boy struct {
	User //匿名字段
	Addr string
}

func NiM(boy Boy) {
	t := reflect.TypeOf(boy)
	fmt.Println(t)
	// Anonymous：匿名
	fmt.Println(t.Field(0))
	//值信息
	fmt.Println(reflect.ValueOf(boy).Field(0))
}

func SetValue(o interface{}) {
	v := reflect.ValueOf(o)
	//获取指针指向的元素
	v = v.Elem()
	//取字段
	f := v.FieldByName("Name")
	if f.Kind() == reflect.String {
		f.SetString("Name")
	}
}

func CallFunc(o interface{}) {
	v := reflect.ValueOf(o)
	//获取方法
	m := v.MethodByName("Hello")
	t := m.Type()
	//构建参数
	args := make([]reflect.Value, t.NumIn())
	for i := 0; i < len(args); i++ {
		if t.In(i).Kind() != reflect.String {
			fmt.Println("Kind Err!")
			return
		}
		args[i] = reflect.ValueOf(strconv.Itoa(i))
	}
	m.Call(args)
}

func GetTag(o interface{}) {
	v := reflect.ValueOf(o)
	t := v.Type()
	f := t.Field(0)
	fmt.Println(f.Tag.Get("json"))
}
