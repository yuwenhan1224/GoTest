package gotest

import "fmt"

func menumm() {
	//普通枚举
	const (
		cpp    = 0
		java   = 1
		python = 2
	)
	fmt.Printf("cpp=%d  java=%d  python=%d\n", cpp, java, python) //a=0  b=1  c=2
	//1.iota只能在常量的表达式中使用
	//fmt.Println(iota)  //undefined: iota
	//2.它默认开始值是0，const中每增加一行加1
	const (
		a = iota //0
		b        //1
		c        //2
	)
	fmt.Printf("a=%d  b=%d  c=%d\n", a, b, c) //a=0  b=1  c=2
	//3.每次 const 出现时，都会让 iota 初始化为0
	const d = iota // a=0
	const (
		e = iota //b=0
		f        //c=1
	)
	fmt.Printf("d=%d  e=%d  f=%d\n", d, e, f) //d=0  e=0  f=1
	//4.如果中断iota，必须显式恢复！！
	const (
		Low    = iota //0
		Medium        //1
		High   = 100  //100，这里本来是2的，但是给High赋值100了，中断了iota的自增，所以在Band显式的恢复了
		Super         //100
		Band   = iota //4
	)
	//Low=0  Medium=1  High=100  Super=100  Band=4
	fmt.Printf("Low=%d  Medium=%d  High=%d  Super=%d  Band=%d\n", Low, Medium, High, Super, Band)
	//5.如果是同一行,值都一样
	const (
		i          = iota
		j1, j2, j3 = iota, iota, iota
		k          = iota
	)
	//i=0  j1=1  j2=1  j3=1  k=2
	fmt.Printf("i=%d  j1=%d  j2=%d  j3=%d  k=%d\n", i, j1, j2, j3, k)
	//6.可跳过的值
	const (
		k1 = iota // 0
		k2        // 1
		_         //2
		_         //3
		k3        // 4
	)
	//	k1=0  k2=1  k3=4
	fmt.Printf("k1=%d  k2=%d  k3=%d \n", k1, k2, k3)
	//7.中间插入一个值
	const (
		Sun = iota //Sun = 0
		Mon        // Mon = 1
		Tue = 7    //7
		Thu = iota // 3
		Fri        //4
	)
	//Sun=0  Mon=1  Tue=7  Thu=3  Fri=4
	fmt.Printf("Sun=%d  Mon=%d  Tue=%d  Thu=%d  Fri=%d\n", Sun, Mon, Tue, Thu, Fri)

}
