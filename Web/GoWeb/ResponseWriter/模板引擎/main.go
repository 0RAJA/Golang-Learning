package main

import (
	"net/http"
	"text/template"
)

/*
Template类型是text/template包的Template类型的特化版本，用于生成安全的HTML文本片段。

func New(name string) *Template
	创建一个名为name的模板。
func ParseFiles(filenames ...string) (*Template, error)
	ParseFiles函数创建一个模板并解析filenames指定的文件里的模板定义。
	返回的模板的名字是第一个文件的文件名（不含扩展名），内容为解析后的第一个文件的内容。至少要提供一个文件。如果发生错误，会停止解析并返回nil。
	t, _ := template.ParseFiles("hello.html")
		相当于:t := template.New("hello.html");t, _ = t.ParseFiles("hello.html")
func ParseGlob(pattern string) (*Template, error)
	ParseGlob创建一个模板并解析匹配pattern的文件（参见glob规则）里的模板定义。
	返回的模板的名字是第一个匹配的文件的文件名（不含扩展名），内容为解析后的第一个文件的内容。
	至少要存在一个匹配的文件。
	如果发生错误，会停止解析并返回nil。ParseGlob等价于使用匹配pattern的文件的列表为参数调用ParseFiles。
	t, _ := template.ParseGlob("*.html")
func (t *Template) Parse(src string) (*Template, error)
	Parse方法将字符串text解析为模板。嵌套定义的模板会关联到最顶层的t。
	Parse可以多次调用，但只有第一次调用可以包含空格、注释和模板定义之外的文本。
	如果后面的调用在解析后仍剩余文本会引发错误、返回nil且丢弃剩余文本；
	如果解析得到的模板已有相关联的同名模板，会覆盖掉原模板。
func Must(t *Template, err error) *Template
	Must函数用于包装返回(*Template, error)的函数/方法调用，它会在err非nil时panic，一般用于变量初始化：
	var t = t := template.Must(template.ParseFiles("D:\\WorkSpace\\Go\\goTest\\src\\Web\\GoWeb\\ResponseWriter\\模板引擎\\hello.html"))

执行模板->
func (t *Template) Execute(wr io.Writer, data interface{}) error
	Execute方法将解析好的模板应用到data上，并将输出写入wr。
	如果执行时出现错误，会停止执行，但有可能已经写入wr部分数据。模板可以安全的并发执行。
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	ExecuteTemplate方法类似Execute，但是使用名为name的t关联的模板产生输出。
	例如:
		t, _ := template.ParseFiles("hello.html", "hello2.html")
		t.ExecuteTemplate(w, "hello2.html", "我要在 hello2.html 中显示")
*/
func handleConn(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("D:\\WorkSpace\\Go\\goTest\\src\\Web\\GoWeb\\ResponseWriter\\模板引擎\\hello.html"))
	//t, err := template.ParseFiles("D:\\WorkSpace\\Go\\goTest\\src\\Web\\GoWeb\\ResponseWriter\\模板引擎\\hello.html")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	t.Execute(w, "HelloWorld")
}

func main() {
	http.HandleFunc("/go", handleConn)
	http.ListenAndServe(":8080", nil)
}
