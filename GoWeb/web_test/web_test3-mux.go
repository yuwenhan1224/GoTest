package main
import(
	"fmt"
	"net/http"

)
//创建处理器的第一种方法：使用HandleFunc
func handler2(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"通过自己创建多路复用器处理请求", r.URL.Path)
}


func main() {
	//创建多路复用器
	mux:=http.NewServeMux()
	//用自己的复用器
    mux.HandleFunc("/" ,handler2)

	http.ListenAndServe(":8088", mux)
}