package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)
//ResponseWriter接口，*http.Request指针
func testTemplate(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"你发送的请求地址是：",r.URL.Path)
     //解析模板
	//t,err:=template.ParseFiles("D:\\GO\\GoCode\\GoWeb\\web_template\\index.html")
   //must函数可专门处理解析模板的错误,自动处理异常
	t:=template.Must(template.ParseFiles("D:\\GO\\GoCode\\GoWeb\\web_template\\index.html","D:\\GO\\GoCode\\GoWeb\\web_template\\index2.html"))

	//t,err:=template.ParseFiles("index.html")
	//real_dir,err:= filepath.Abs("./index.html")go get -u github.com/gin-gonic/gin
	//t,err2:=template.ParseFiles(real_dir)
	//fmt.Println("index.htm的绝对路径：",real_dir, err)
	//fmt.Println(err)
	//执行
	//t.Execute(w,"Hello Template")//t.Execute(输出到哪,"传入数据")
	//t是一个包含看两个模板的模板集合，如果调用Execute方法在，则只有第一个模板被执行
	//将响应数据在index2.html文件中显示
	t.ExecuteTemplate(w,"D:\\GO\\GoCode\\GoWeb\\web_template\\index2.html","我要在index2.html中显示")

}

func main(){
	http.HandleFunc("/testTemplate",testTemplate)//”映射的路径，设置访问的路由
	err:=http.ListenAndServe(":8888",nil)////创建路由 设置监听的端口9090,nil用默认的多路复用器
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}












