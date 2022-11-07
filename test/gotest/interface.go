package gotest

import "fmt"
type Phone interface {
	Call()
	Start()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) Call() {
	fmt.Println("I am Nokia, I can call you!")
}

func (nokiaPhone NokiaPhone) Start() {
	fmt.Println("I am Nokia, I can start you!")
}

type IPhone struct {
}

func (iPhone IPhone) Call() {
	fmt.Println("I am iPhone, I can call you!")
}

func (iPhone IPhone) Start() {
	fmt.Println("I am iPhone, I can start you!")
}

func main1()  {


	//NokiaPhone对象给Phone类型变量赋值的两种方法：
	//(1)new方法
	var phone Phone
	phone = new(NokiaPhone)
	phone.Call()
	//(2)赋值法
	var Nokia2 NokiaPhone
	phone=Nokia2
	phone.Call()

	phone=new(IPhone)
	phone.Call()




}
