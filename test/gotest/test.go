package gotest

import "fmt"

var A string = "DDDDD"

/* 定义相互交换值的函数 */
func swap2(x, y int) int {
	var temp int

	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/

	return temp;
}
func swap3( x *int,y *int){
	var temp int
	temp=*x
	*y=temp
}

func variable(){
	//var x [10]float32
	//var q=[5]float32{1,2,4,2356,6}
	//qq:=[...]float32{1,2,3,45,3}
	//balance:=[5]float32{1:2,3:5}

	var n [10]int
	var i,j int
	for i=0;i<10;i++{
		n[i]=i+100
	}
	for j=0;j<10;j++{
		fmt.Println("Element[%d]=%d\n",j,n[j])
	}


	//var threedim [5][10][4]int
	//var ip *int
	//var fp *float32
	//
	//var a int =20
	//var ip *int
	//ip=&a
	type Books struct {
		title string
		author string
		subject string
		book_id int
	}
	fmt.Println(Books{})
}