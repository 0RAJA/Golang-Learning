package main

import (
	"fmt"
	"package_init/tool"
)

/*
关于包的使用:
	1.一个目录下的统计文件归属一个包。package的声明要一致
	2.package声明的包和对应的目录名可以不一致。但习惯上还是写成一致的
	3.包可以嵌套
	4.同包下的函数不需要导入包，可以直接使用
	5.main包:main() 函数所在的包，其他的包不能使用
	6.导入包的时候，路径要从src下开始写

init()函数和main()函数
	1.这两个函数都是go语言中的保留函数。init()用于初始化信息，main()用于作为程序的入口
	2.这两个函数定义的时候:不能有参数,返回值。只能由go程序自动调用,不能被引用。
	3.init()函数可以定义在任意的包中，可以有多个.main()函数只能在main包下,并且只能有一个。
	4.执行顺序
		A:先执行init()函数，后执行main()函数
		B:对于同一个go文件中，init()调用顺序是从上到下的，也就是说，先写的先被执行，后写的后被执行
		C:对于同一个包下，将文件名按照字符串进行排序，之后顺序调用各个文件中init()函数
		D:对于不同包下:
			如果不存在依赖，按照main包中import的顺序来调用对应包中init()函数
			如果存在依赖，最后被依赖的最先被初始化
				导入顺序: main->A-->B-->C
				执行顺序: C-->B-->A-->main
	5.存在依赖的包之间，不能循环导入:A –> B –> C –> A。
	6.一个包可以被其他多个包import,但是只能被初始化一次。
	7.点操作
		import(
			. "fmt"
		)
		这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，
		也就是前面你调用的`fmt.Println("hello world")`可以省略的写成`Println("hello world")`
	8.起别名
		import (
  		p1 "package1"
  		p2 "package2"
  		)
		// 使用时：别名操作，调用包函数时前缀变成了我们的前缀
		p1.Method()
	9. _ 操作
		如果仅仅需要导入包时执行初始化操作，并不需要使用包内的其他函数，常量等资源。则可以在导入包时，匿名导入。
		import (
   		"database/sql"
   		_ "github.com/ziutek/mymysql/godrv"
 		)
		_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。
			也就是说，使用下划线作为包的别名，会仅仅执行init()。
*/
func main() {
	tool.MyTime()
	tool.MyTime2()
	s := tool.Student{Name: "张三", Age: 20}
	fmt.Println(s)
}
