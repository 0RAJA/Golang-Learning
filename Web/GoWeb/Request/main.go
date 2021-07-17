package main

import (
	"fmt"
	"net/http"
)

/*
	Go 语言的 net/http 包提供了一系列用于表示 HTTP 报文的结构，我们可以使用它处理请求和发送相应，
	其中 Request 结构代表了客户端发送的请求报文，
	Request类型代表一个服务端接受到的或者客户端发送出去的HTTP请求。
	Request各字段的意义和用途在服务端和客户端是不同的。除了字段本身上方文档，还可参见Request.Write方法和RoundTripper接口的文档。
type Request struct {
	// Method指定HTTP方法（GET、POST、PUT等）。对客户端，""代表GET。
	Method string
	// URL在服务端表示被请求的URI，在客户端表示要访问的URL。
	//
	// 在服务端，URL字段是解析请求行的URI（保存在RequestURI字段）得到的，
	// 对大多数请求来说，除了Path和RawQuery之外的字段都是空字符串。
	// （参见RFC 2616, Section 5.1.2）
	//
	// 在客户端，URL的Host字段指定了要连接的服务器，
	// 而Request的Host字段（可选地）指定要发送的HTTP请求的Host头的值。
	URL *url.URL
	// 接收到的请求的协议版本。本包生产的Request总是使用HTTP/1.1
	Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0
	// Header字段用来表示HTTP请求的头域。如果头域（多行键值对格式）为：
	//	accept-encoding: gzip, deflate
	//	Accept-Language: en-us
	//	Connection: keep-alive
	// 则：
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Connection": {"keep-alive"},
	//	}
	// HTTP规定头域的键名（头名）是大小写敏感的，请求的解析器通过规范化头域的键名来实现这点。
	// 在客户端的请求，可能会被自动添加或重写Header中的特定的头，参见Request.Write方法。
	Header Header
	// Body是请求的主体。
	//
	// 在客户端，如果Body是nil表示该请求没有主体买入GET请求。
	// Client的Transport字段会负责调用Body的Close方法。
	//
	// 在服务端，Body字段总是非nil的；但在没有主体时，读取Body会立刻返回EOF。
	// Server会关闭请求的主体，ServeHTTP处理器不需要关闭Body字段。
	Body io.ReadCloser
	// ContentLength记录相关内容的长度。
	// 如果为-1，表示长度未知，如果>=0，表示可以从Body字段读取ContentLength字节数据。
	// 在客户端，如果Body非nil而该字段为0，表示不知道Body的长度。
	ContentLength int64
	// TransferEncoding按从最外到最里的顺序列出传输编码，空切片表示"identity"编码。
	// 本字段一般会被忽略。当发送或接受请求时，会自动添加或移除"chunked"传输编码。
	TransferEncoding []string
	// Close在服务端指定是否在回复请求后关闭连接，在客户端指定是否在发送请求后关闭连接。
	Close bool
	// 在服务端，Host指定URL会在其上寻找资源的主机。
	// 根据RFC 2616，该值可以是Host头的值，或者URL自身提供的主机名。
	// Host的格式可以是"host:port"。
	//
	// 在客户端，请求的Host字段（可选地）用来重写请求的Host头。
	// 如过该字段为""，Request.Write方法会使用URL字段的Host。
	Host string
	// Form是解析好的表单数据，包括URL字段的query参数和POST或PUT的表单数据。
	// 本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代。
	Form url.Values
	// PostForm是解析好的POST或PUT的表单数据。
	// 本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代。
	PostForm url.Values
	// MultipartForm是解析好的多部件表单，包括上传的文件。
	// 本字段只有在调用ParseMultipartForm后才有效。
	// 在客户端，会忽略请求中的本字段而使用Body替代。
	MultipartForm *multipart.Form
	// Trailer指定了会在请求主体之后发送的额外的头域。
	//
	// 在服务端，Trailer字段必须初始化为只有trailer键，所有键都对应nil值。
	// （客户端会声明哪些trailer会发送）
	// 在处理器从Body读取时，不能使用本字段。
	// 在从Body的读取返回EOF后，Trailer字段会被更新完毕并包含非nil的值。
	// （如果客户端发送了这些键值对），此时才可以访问本字段。
	//
	// 在客户端，Trail必须初始化为一个包含将要发送的键值对的映射。（值可以是nil或其终值）
	// ContentLength字段必须是0或-1，以启用"chunked"传输编码发送请求。
	// 在开始发送请求后，Trailer可以在读取请求主体期间被修改，
	// 一旦请求主体返回EOF，调用者就不可再修改Trailer。
	//
	// 很少有HTTP客户端、服务端或代理支持HTTP trailer。
	Trailer Header
	// RemoteAddr允许HTTP服务器和其他软件记录该请求的来源地址，一般用于日志。
	// 本字段不是ReadRequest函数填写的，也没有定义格式。
	// 本包的HTTP服务器会在调用处理器之前设置RemoteAddr为"IP:port"格式的地址。
	// 客户端会忽略请求中的RemoteAddr字段。
	RemoteAddr string
	// RequestURI是被客户端发送到服务端的请求的请求行中未修改的请求URI
	// （参见RFC 2616, Section 5.1）
	// 一般应使用URI字段，在客户端设置请求的本字段会导致错误。
	RequestURI string
	// TLS字段允许HTTP服务器和其他软件记录接收到该请求的TLS连接的信息
	// 本字段不是ReadRequest函数填写的。
	// 对启用了TLS的连接，本包的HTTP服务器会在调用处理器之前设置TLS字段，否则将设TLS为nil。
	// 客户端会忽略请求中的TLS字段。
	TLS *tls.ConnectionState
}
*/
/*
请求行:URL字段
获取请求URL
	Request 结构中的 URL 字段用于表示请求行中包含的 URL，该字段是一个指向url.URL 结构的指针，让我们来看一下 URL 结构
type URL struct {
	Scheme   string
	Opaque   string    // 编码后的不透明数据
	User     *Userinfo //用户名和密码信息
	Host     string    // host或host:port
	Path     string
	RawQuery string //编码后的查询字符串，没有'?'
	Fragment string //引用的片段(文档位置)，没有'#'
}
1)Path 字段、
	获取请求的 URL
	例如：http://localhost:8080/hello?username=admin&password=123456
		i.通过 r.URL.Path 只能得到 /hello
2)RawQuery 字段
	获取请求的 URL 后面?后面的查询字符串
	例如：http://localhost:8080/hello?username=admin&password=123456
		i.通过 r.URL.RawQuery 得到的是 username=admin&password=123456
*/
/*
请求头:Header字段
type Header map[string][]string
	Header代表HTTP头域的键值对。
func (h Header) Get(key string) string
	Get返回键对应的第一个值，如果键不存在会返回""。如要获取该键对应的值切片，请直接用规范格式的键访问map。
func (h Header) Set(key, value string)
	Set添加键值对到h，如键已存在则会用只有新值一个元素的切片取代旧值切片。
func (h Header) Add(key, value string)
	Add添加键值对到h，如键已存在则会将新的值附加到旧值切片后面。
func (h Header) Del(key string)
	Del删除键值对。
func (h Header) Write(w io.Writer) error
	Write以有线格式将头域写入w。
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error
	WriteSubset以有线格式将头域写入w。当exclude不为nil时，如果h的键值对的键在exclude中存在且其对应值为真，该键值对就不会被写入w。
*/
/*
请求体:Body字段
	请求和响应的主体都是由 Request 结构中的 Body 字段表示，这个字段的类型是io.ReadCloser 接口，
	该接口包含了 Reader 接口和 Closer 接口，Reader 接口拥有 Read方法，Closer 接口拥有 Close 方法
type ReadCloser interface {
	Reader
	Closer
}
	ReadCloser接口聚合了基本的读取和关闭操作。
type Reader interface {
	Read(p []byte) (n int, err error)
}
Reader接口用于包装基本的读取方法。
	Read方法读取len(p)字节数据写入p。它返回写入的字节数和遇到的任何错误。即使Read方法返回值n < len(p),
	本方法在被调用时仍可能使用p的全部长度作为暂存空间。如果有部分可用数据，但不够len(p)字节，Read按惯例
	会返回可以读取到的数据，而不是等待更多数据。
	当Read在读取n > 0个字节后遭遇错误或者到达文件结尾时，会返回读取的字节数。它可能会在该次调用返回一个
	非nil的错误，或者在下一次调用时返回0和该错误。一个常见的例子， Reader接口会在输入流的结尾返回非0的字
	节数，返回值err == EOF或err== nil,但不管怎样，下一次Read调用必然返回(0, EOF)。调用者应该总是先处理
	读取的n > 0字节再处理错误值。这么做可以正确的处理发生在读取部分数据后的IO错误，也能正确处理EOF事件。
	如果Read的某个实现返回0字节数和nil错误值，表示被阻碍;调用者应该将这种情况视为未进行操作。
type Closer interface {
	Close() error
}
	Closer接口用于包装基本的关闭方法。在第一次调用之后再次被调用时，Close方法的的行为是未定义的。某些实现可能会说明他们自己的行为。
*/
/*
请求参数:Form 字段
下面我们就通过 net/http 库中的 Request 结构的字段以及方法获取请求 URL 后面的请求参数以及 form 表单中提交的请求参数
1.类型是 url.Values 类型，Form 是解析好的表单数据，包括 URL 字段的 query 参数和 POST 或 PUT 的表单数据
	type Values map[string][]string
		Values将键映射到值的列表。它一般用于查询的参数和表单的属性。不同于http.Header这个字典类型，Values的键是大小写敏感的。
2.Form 字段只有在调用 Request 的 ParseForm 方法后才有效。在客户端，会忽略请求中的本字段而使用 Body 替代
	func (r *Request) ParseForm() error
		ParseForm解析URL中的查询字符串，并将解析结果更新到r.Form字段。
		对于POST或PUT请求，ParseForm还会将body当作表单解析,并将结果既更新到r.PostForm也更新到r.Form.解析结果中，
			POST或PUT请求主体要优先于URL查询字符串(同名变量,主体的值在查询字符串的值前面)。
		如果请求的主体的大小没有被MaxBytesReader函数设定限制，其大小默认限制为开头10MB.
	ParseMultipartForm会自动调用ParseForm.重复调用本方法是无意义的。
3.获取表单中提交的请求参数
*/
func handler(w http.ResponseWriter, r *http.Request) {
	//获取请求行中的信息
	fmt.Fprintln(w, "你请求地址为:", r.URL.Path)
	fmt.Fprintln(w, "你请求地址后的查询字符串为:", r.URL.RawQuery)
	/*
		URL:http://localhost:8080/go?user=admin&pwd=123
			r.URL.Path = "/go"
			r.URL.RawQuery = "user=admin&pwd=123"
	*/
	//获取请求头中的所有信息
	fmt.Fprintln(w, "请求头信息:", r.Header)
	//获取请求头中指定信息
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息:", r.Header["Accept-Encoding"])       //字符串切片
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性信息:", r.Header.Get("Accept-Encoding")) //字符串
	//获取请求体中的信息
	//获取请求体长度
	//length := r.ContentLength
	//创建byte切片
	//body := make([]byte, length)
	//读内容
	//r.Body.Read(body)//注意!!!这里对body内容的读取会导致之后对body内容的使用为空,因为读取到最后了
	//fmt.Fprintln(w, "请求体中的内容:", string(body))
	//获取请求参数
	//如果form表单的action属性的URL地址中也有与form表单参数名相同的请求参数，
	//那么参数值都可以得到，并且form表单中的参数值在ULR的参数值的前面
	//解析表单
	r.ParseForm()
	//获取参数
	fmt.Fprintln(w, "请求参数为:", r.Form)
	fmt.Fprintln(w, "POST请求参数为:", r.PostForm) //只获取post表单的值
	/*
		但是 PostForm 字段只支持 application/x-www-form-urlencoded 编码，
		如果form 表单的 enctype 属性值为 multipart/form-data，
		那么使用 PostForm 字段无法获取表单中的数据，此时需要使用 MultipartForm 字段
		说明：form 表单的 enctype 属性的默认值是 application/x-www-form-urlencoded 编码
		实现上传文件时需要讲该属性的值设置为 multipart/form-data 编码格式
	*/
	//直接获取参数
	/*
		func (r *Request) FormValue(key string) string
			FormValue返回key为键查询r.Form字段得到结果[]string切片的第一个值。
			POST和PUT主体中的同名参数优先于URL查询字符串。如果必要，本函数会隐式调用ParseMultipartForm和ParseForm.
		func (r *Request) PostFormValue(key string) string
			PostFormValue返回key为键查询r. PostForm字段得到结果[]string切片的第一个值。
			如果必要，本函数会隐式调用ParseMultipartForm和ParseForm.
	*/
	fmt.Fprintln(w, "直接获取username参数:", r.FormValue("username"))
}

func main() {
	http.HandleFunc("/go", handler)
	http.ListenAndServe(":8080", nil)
}
