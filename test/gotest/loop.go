package gotest

import "fmt"

func goloop() {
	/* 定义局部变量 */
	var a int = 10

	/* 循环 */
LOOP: for a < 20 {
	if a == 15 {
		/* 跳过迭代 */
		a = a + 1
		goto LOOP
	}
	fmt.Printf("a的值为 : %d\n", a)
	a++
}
}
func goloop2() {

	fmt.Printf("goloop2")
}
//func cycle11() {
//	forCycle()
//	gotocycle()
//}

// forCycle
//
//	@Description:for循环
//	@return int
func forCycle() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	/////////
	sum2 := 1
	for sum2 <= 10 {
		sum2 += sum2
	}
	fmt.Println(sum2)
	//while语句
	sum3 := 1
	for sum3 <= 10 {
		sum3 += sum3
	}
	fmt.Println(sum3)
	//无限循环
	//sum4 := 0
	//for {
	//	sum4++
	//}
	//fmt.Println(sum4)
	//func main() {
	//	for true  {
	//		fmt.Printf("这是无限循环。\n");
	//	}
	//}
}

