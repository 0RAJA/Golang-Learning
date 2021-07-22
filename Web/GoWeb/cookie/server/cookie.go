package main

import (
	"fmt"
	"net/http"
)

/*
type Cookie struct {
    Name       string
    Value      string
    Path       string
		//设置当前的 Cookie 值只有在访问指定路径时才能被服务器程序读取。
		//默认为服务端应用程序上的任何路径，但是您可以使用它限制为特定的子目录
    Domain     string
		//默认值是当前正在访问的Host的域名，假设我们现在正在访问的是www.example.com，如果需要其他子域名也能够访问到正在设置的Cookie值的话，
		//将它设置为example.com 。
		//注意，只有正在被设置的Cookie需要被其他子域名的服务访问到时才这么设置
    Expires    time.Time
		//持续时间
    RawExpires string
    	// MaxAge=0表示未设置Max-Age属性
    	// MaxAge<0表示立刻删除该cookie，等价于"Max-Age: 0"
    	// MaxAge>0表示存在Max-Age属性，单位是秒
    MaxAge   int
    Secure   bool
    HttpOnly bool
		//为避免跨域脚本 (XSS) 攻击，通过JavaScript的API无法访问带有 HttpOnly 标记的Cookie，它们只应该发送给服务端。如果包含服务端Session 信息的Cookie 不想被客户端JavaScript 脚本调用，那么就应该为其设置 HttpOnly 标记。
    Raw      string
    Unparsed []string // 未解析的“属性-值”对的原始文本
}

Cookie代表一个出现在HTTP回复的头域中Set-Cookie头的值里或者HTTP请求的头域中Cookie头的值里的HTTP cookie。
运行原理:
1)	第一次向服务器发送请求时在服务器端创建 Cookie
2)	将在服务器端创建的 Cookie 以响应头的方式发送给浏览器
3)	以后再发送请求浏览器就会携带着该 Cookie
4)	服务器得到 Cookie 之后根据 Cookie 的信息来区分不同的用户

func SetCookie(w ResponseWriter, cookie *Cookie)
	SetCookie在w的头域中添加Set-Cookie头，该HTTP头的值为cookie。
*/

//创建cookie并发送给浏览器
func setCookie(w http.ResponseWriter, r *http.Request) {
	//创建Cookie
	cookie := http.Cookie{
		Name:     "user",
		Value:    "admin", //携带:user=admin
		HttpOnly: true,
		MaxAge:   100, //有效时间,不设置默认关闭浏览器之后失效,设置后按照时间失效
	}
	//添加第二个Cookie
	cookie2 := http.Cookie{
		Name:     "user2",
		Value:    "ama",
		HttpOnly: true,
	}
	//将Cookie发送给浏览器
	//第一种:
	//w.Header().Set("Set-Cookie", cookie.String())
	//w.Header().Add("set", cookie2.String())
	//第二种:(推荐)
	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie2)
}

//
func getCookies(w http.ResponseWriter, r *http.Request) {
	//获取所有cookie
	/*
		由于我们在发送请求时 Cookie 在请求头中，所以我们可以通过 Request 结构中的
		Header 字段来获取 Cookie

	*/
	//得到所有的cookie
	cookie := r.Header["Cookie"]
	fmt.Println(w, cookie)
	//得到指定名字的cookie
	fmt.Println(r.Cookie("user"))
}

func main() {
	http.HandleFunc("/set-Cookie", setCookie)
	http.HandleFunc("/get-Cookie", getCookies)
	http.ListenAndServe(":8080", nil)
}
