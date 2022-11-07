package model

import (
	"fmt"
	"testing"
)

//testing 提供对 Go 包的自动化测试的支持。通过 `go test` 命令，能够自动执行如下形式的任何函数：func TestXxx(*testing.T)
//其中 Xxx 可以是任何字母数字字符串（但第一个字母不能是 [a-z]），用于识别测试例程。
//命令行输入go test 即可测试
//命令行输入go test -v 可查看测试详细情况
//////////////////////////////////////////////测试方法一：
//func TestAddUser(t *testing.T){
//	fmt.Println("测试添加用户")
//	user:=&User{}
//	//调用添加用户方法
//	//user.AddUser()
//	//user.AddUser2()
//	user.SelectUse()
//}
//////////////////////////////////////////////测试方法二：

/////如果该函数不是以Test开头，那么该函数默认不执行，解决办法：我们可以将她设置成一个子测试函数
//TestMain函数可以在测试函数执行之前做一些其他操作
func TestMain(m *testing.M){
	fmt.Println("测试开始：")
	//通过m.run来执行测试函数
	m.Run()
}

func TestUser(t *testing.T){
	fmt.Println("开始测试User中的相关方法")
	//通过t.Run()来测试子函数
	//t.Run("测试添加用户：",testAddUser)
	t.Run("测试获取用户：",testGetUserById)
	t.Run("测试获取用户：",testGetUsers)
}
func testAddUser(t *testing.T){
	fmt.Println("子测试函数执行：")
	user:=&User{}
	//调用添加用户方法
	user.AddUser()
	user.AddUser2()
	user.GetUsers2()
}
//测试获取一个User
 func testGetUserById(t *testing.T) {
	 fmt.Println("测试查询一条记录:")
	 user:=User{
		 ID:1,
	 }
	 //调用获取User的方法
	 u,_:=user.GetUserById()
	 fmt.Println("得到的User的信息",u)
 }

 func testGetUsers(t *testing.T) {
	 fmt.Println("测试查询所有记录：")
	 user:=&User{}
	 //调用获取所有的User的方法
	 us,_:=user.GetUsers()
	 //遍历切片
	 for k,v :=range us{
		 fmt.Printf("第%v一个用户是%v：\n",k+1,v)//因为key从0开始
	 }
 }