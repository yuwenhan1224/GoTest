
package main

import (
	"fmt"      //格式化I/O
	"log"      //简单的日志服务,把日志输出，只要是实现了io.writer接口的write方法就可以输出到任意地方，默认输出到控制台
	"net/http" //创建 Web 服务器的标准库,提供了HTTP客户端和服务端的实现
)
/**
sayHello函数需要2个参数，一个是http.ResponseWriter，另一个是http.Request
往http.ResponseWriter写入什么内容，浏览器的网页源码就是什么内容；
http.Request里面封装了浏览器发过来的请求（包含路径、浏览器类型等等）。
*/
/*
scheme://host[:port#]/path/.../[?query-string][#anchor]
scheme         指定底层使用的协议(例如：http, https, ftp)
host           HTTP 服务器的 IP 地址或者域名
port#          HTTP 服务器的默认端口是 80，这种情况下端口号可以省略。如果使用了别的端口，必须指明，例如 http://www.cnblogs.com:8080/
path           访问资源的路径
query-string   发送给 http 服务器的数据
anchor         锚

 */
//创建处理器函数，一个处理器就是一个拥有ServeHTTP方法的接口，里面的参数必须是这两个
//Request：用户请求的信息,post、get、url等这些信息
//Response: 返回给客户端的信息
//Conn: 用户每次的连接请求
//Handler：处理请求和返回信息的逻辑处理

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w,"------------------")
	//r.ParseForm()  // 解析参数，默认是不会解析的
	//
	//fmt.Println(r.Form)  // 这些信息是输出到服务器端的打印信息
	//fmt.Println("path", r.URL.Path)//打印访问资源路径
	//fmt.Println("scheme", r.URL.Scheme)//打印指定底层协议：http\https\ftp
	//fmt.Println(r.Form["url_long"])
	//for k, v := range r.Form {
	//	fmt.Println("key:", k)
	//	fmt.Println("val:", strings.Join(v, ""))
	//}
	fmt.Fprintf(w, "Hello 123!",r.URL.Path) // 这个写入到 w 的是输出到客户端的
}

func main() {
	//http包里面还定义了一个类型HandlerFunc,我们定义的函数sayhelloName就是这个HandlerFunc调用之后的结果，
	//这个类型默认就实现了ServeHTTP这个接口，即我们调用了HandlerFunc(f),强制类型转换f成为HandlerFunc类型，这样f就拥有了ServeHTTP方法
	//HandlerFunc是一个适配器，通过类型转换可以把普通函数作为HTTP处理器使用，如果函数是一个具有适当签名的函数
	//（也就是这个函数的参数不能乱写，必须是(w http.ResponseWriter, r *http.Request），就像ServeHTTP方法中的形参列表一样，HandlerFunc才可以转化
	//HandlerFunc(f)通过调用f实现了Handler接口
	//如果要创建处理器，必须实现ServeHTTP方法。
	//我们也可以不使用HandleFunc，自己创建一个结构体，通过使用结构体实现ServeHTTP方法,从而来创建一个处理器
	http.HandleFunc("/", sayhelloName) // “/”映射的路径，设置访问的路由，“ sayhelloName”调用处理器
	//ListenAndServe 函数并传入网络地址以及负责处理请求的处理器( handler )作为参数就可以了
	//如果网络地址参数为空字符串，那 么服务器默认使用 80 端口进行网络连接；如果处理器参数为 nil，那么服务器将使用默认的多路复用器 DefaultServeMux
	//创建路由，http1.0,http1.1,http2.0就是https
	err := http.ListenAndServe(":8080", nil) //创建路由 设置监听的端口9090,nil用默认的多路复用器
//如果ListenAndServe里面的网络地址参数为空字符串，那么服务器默认使用80端口进行网络连接
	//http.ListenAndServe:底层其实这样处理的：初始化一个 server 对象，然后调用了 net.Listen("tcp", addr)，也就是底层用 TCP 协议搭建了一个服务，然后监控我们设置的端口。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}