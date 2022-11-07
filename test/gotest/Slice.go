package gotest

import"fmt"

func Slice55(){
	var number=make([]int,3,5)
	fmt.Printf("len=%d cap=%d slice=%v\n",len(number),cap(number),number)
}