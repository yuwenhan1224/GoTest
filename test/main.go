package main //定义包名
import (
	"fmt"
	"test/gotest"
	"time"
)
func say(s string){
	for i:=0;i<5;i++{
		time.Sleep(100*time.Microsecond)
		fmt.Printf("hello:%d\n",i)
	}

}
func main()  {
	//student对象给people类型变量赋值
	//值类型实现接口
	var peo gotest.People=gotest.Student2{}
	think:="hello"
	fmt.Println(peo.Speak(think))
	//指针类型实现接口
	var peo2 gotest.People=&gotest.Student2{}
	fmt.Println(peo2.Speak("hello world"))
	var peo3 gotest.People=&gotest.Doctor{}
	fmt.Println(peo3.Speak("check"))


    //NokiaPhone对象给Phone类型变量赋值的两种方法：
	//(1)new方法
	var phone gotest.Phone
	phone = new(gotest.NokiaPhone)
	phone.Call()
   //(2)赋值法
	var Nokia2 gotest.NokiaPhone
	phone=Nokia2
	phone.Call()

	phone=new(gotest.IPhone)
	phone.Call()


	var skill  gotest.Skills
	var stu1  gotest.Student
	stu1.Name = "interface"
	stu1.Age = 22
	skill = stu1//
	skill.Running()  //调用接口

	go say("world")
	say("hello")
}











//fmt.Printf("hello")

//// %d 表示整型数字，%s 表示字符串

////fmt.Println(fff.A)
//selectt()
//goloop()
//goloop2()
//ffunc()
//fmt.Println(A)
//variable()
//selectt()
//gotest.Slice55()