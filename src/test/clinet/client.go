package main

import (
	"fmt"
	"net"
)

func main() {
	// 客户端开始连接
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//发送数据

	conn.Write([]byte("hello"))
	defer conn.Close() //延时关闭连接
}
