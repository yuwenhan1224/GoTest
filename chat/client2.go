package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// 用户名
var loginName2 string

// 本机连接
var selfconnect22 *net.TCPConn

// 读取行文本
var reader2 =bufio.NewReader(os.Stdin)

// 建立连接
func connect2(addr string) {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", addr) // 使用tcp
	con, err := net.DialTCP("tcp", nil, tcpAddr)  // 拨号：主动向server建立连接
	selfconnect22 = con
	if err != nil {
		fmt.Println("连接服务器失败")
		os.Exit(1)
	}
	go msgSender2()
	go msgReceiver2()
}

// 消息接收器
func msgReceiver2() {
	buff := make([]byte, 2048)
	for {
		len, _ := selfconnect22.Read(buff) // 从建立连接的缓冲区读消息
		fmt.Println(string(buff[:len]))
	}
}

// 消息发送器
func msgSender2() {
	for {
		bMsg, _, _ := reader2.ReadLine()
		bMsg = []byte(loginName2 + " : " + string(bMsg))
		selfconnect22.Write(bMsg) // 发消息
	}
}

// 初始化
func initGroupChatClient2() {
	fmt.Println("请问您怎么称呼？")
	bName, _, _ := reader2.ReadLine()
	loginName2 = string(bName)
	connect2("127.0.0.1:1801")
	select {}
}

func main() {
	initGroupChatClient2()
}