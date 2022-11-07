package gotest

import "fmt"

//接口的作用就是定义一些对象的共同特征，特征以接口内的方法来定义
type People interface {
	Speak(string) string
}
//定义类型student2,是一种struct类型
type Student2 struct {

}
// 类型Student实现People特征 (也就是让它实现People中的方法)
func (stu Student2) Speak(think string) (talk string){
	if think=="hello world"{
		talk="你好，世界"
	}else{
		talk="您好"
	}
	return
}
type Doctor struct {

}
func (stu *Doctor) Speak(think string) (talk string){
	if think=="check"{
		talk="检查"
	}else{
		talk="治疗"
	}
	return
}
func main3()  {
	//student对象给people类型变量赋值
	//值类型实现接口
	var peo People=Student2{}
	think:="hello"
	fmt.Println(peo.Speak(think))
	//指针类型实现接口
	var peo2 People=&Student2{}
	fmt.Println(peo2.Speak("hello world"))
	var peo3 People=&Doctor{}
	fmt.Println(peo3.Speak("check"))


}
