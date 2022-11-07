package main
import(
	"fmt"
	"net/http"
	"time"
)
//创建处理器的第一种方法：使用HandleFunc
func handler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"hello world", r.URL.Path)
}
//创建处理器的第二种方法
//(1)创建结构体，不使用HandleFunc，自己实现Handler接口里的ServceHTTP方法
type MyHandler struct {
}
//(2)实现Handler接口里的ServceHTTP方法
func (m *MyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){
	//fmt.Fprintf(w,"通过自己创建的处理器去处理请求")
	fmt.Fprintf(w,"通过详细配置服务器信息来处理器去处理请求")

}
func main(){
	//（1）
	 // http.HandleFunc("/" ,handler)

	  //(2)
     myHandler:=MyHandler{}
	 //使用Handle时，处理器是myHandler,实现了Handler里面的ServceHTTP方法接口的结构
	/* http.Handle("/myHandler",&myHandler)//浏览器访问/myHandler的时候才能调用我们的处理器处理请求*/
	//创建Server结构,并详细配置里面的字段
	server:=http.Server{
		Addr: "8080",
		Handler: &myHandler,
		ReadTimeout: 2*time.Second,
	}
	// http.ListenAndServe(":8080",nil)
	 server.ListenAndServe()//不用这样”http.ListenAndServe(":8080",nil)”配置路由，因为server接口里面已经有ListenAndServer方法了
}