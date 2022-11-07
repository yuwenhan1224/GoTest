package main

import (
	"fmt"
	"log"
	"net/http"
)

//创建处理器函数
func HandlerHttp(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"测试Http协议")
}
func main(){
	//调用处理器请求
  http.HandleFunc("/http",HandlerHttp)
	//路由
  err :=http.ListenAndServe(":8088",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
