package main

import (
	"GoCode/GoWeb/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//ResponseWriter接口，*http.Request指针
func handler(w http.ResponseWriter,r *http.Request){
	//处理请求，获取请求参数的第一种方法/////////////////////
	fmt.Fprintln(w,"你发送的请求地址是：",r.URL.Path)
	fmt.Fprintln(w,"你发送的请求地址后的查询字符串是：",r.URL.RawQuery)
	fmt.Fprintln(w,"请求头中的所有信息：",r.Header)
	fmt.Fprintln(w,"请求头中的Accept-Encoding的信息：",r.Header["Accept-Encoding"])//结果返回切片---[gzip, deflate, br]
	fmt.Fprintln(w,"请求头中的Accept-Encoding的属性值信息：",r.Header.Get("Accept-Encoding"))//结构直接返回切片里的值---gzip, deflate, br
	//获取请求体内容长度
	//len:=r.ContentLength
	//创建byte切片
	//body:=make([]byte,len)
	//将请求体Body中的内容读取到body的内容中
	//r.Body.Read(body)
	//在浏览器中显示请求体中的内容
	//fmt.Fprintln(w,"请求体中的内容有：",string(body))

	//处理请求，获取请求参数的第二种方法//////////////////////
	err:=r.ParseForm()
	fmt.Println("错误信息是：",err)
	//解析表单，在调用r.Form之前必须执行该操作
	//r.ParseForm()
	//获取请求参数
	//如果from表单的action属性的地址值中也有与from表单参数名相同的请求参数，
	//<form action="http://localhost:8080/hello?username=yu&password=666" method="POST">
	//那么参数值都可以得到，并且from表单中的参数值在URL的参数值的前面
	//fmt.Fprintln(w,"Fprintln::请求参数有：",r.Form)//可以直接在网页上打印出参数
  // fmt.Fprintln(w,"Fprintln::POST请求的form表单的请求参数有：",r.PostForm)

	//处理请求，获取请求参数的第三种方法/////////////////////
	//通过直接调用FormValue方法和PostFormValue方法直接获取请求参数
   //FormValue方法可查询既可以获取URL里面的请求参数也可以获取Form表单的
   //PostFormValue只可以获取Form表单里面的参数
    fmt.Fprintln(w,"URL中的user请求参数的值是：",r.FormValue("user"))
	//PostForm字段只支持"application/x-www-form-urlencoded"编码，如果form表单的entype属性值为multipart/form-data，那么使用Postform字段无法获取表单内容
	fmt.Fprintln(w,"Form表单中的username请求参数的值是：",r.PostFormValue("username"))
}

func testJsonRes(w http.ResponseWriter,r *http.Request){
	//设置响应内容的类型
	w.Header().Set("Content-Type","application/json")
	//创建User
	user:=model.User{
		ID:1,
		Username: "admin",
		Password: "123456",
		Email: "admin@222.com",
	}

	//将User转换为Json个数
	json,_:=json.Marshal(user)
	//将json格式的数据相应给客户端
	w.Write(json)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	html := `<html>
<head>
<title>测试响应内容为网页</title>
<meta charset="utf-8"/>
</head>
<body>
我是以网页的形式响应过来的！
</body>
</html>`
	w.Write([]byte(html))
}
func handler3(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("你的请求我已经收到"))
}
func testRedire(w http.ResponseWriter,r *http.Request){
	//设置响应头的Location属性
	w.Header().Set("Location","https://www.baidu.com")
	//设置响应的状态码
    w.WriteHeader(302)//重定向的状态码是302

}

func main(){
	http.HandleFunc("/hello",handler)//”映射的路径，设置访问的路由
	http.HandleFunc("/hello2",handler2)//”映射的路径，设置访问的路由
	http.HandleFunc("/hello3",handler3)//”映射的路径，设置访问的路由
	http.HandleFunc("/testJson",testJsonRes)
	http.HandleFunc("/testRedirect",testRedire)
	err:=http.ListenAndServe(":8888",nil)////创建路由 设置监听的端口9090,nil用默认的多路复用器
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}












