package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	//监听端口
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Print("监听出问题：", err)
	}
	for {
		//连接等待
		fmt.Println("等待连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept,err=", err)
		}
		//(********)
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Println("reader.Read err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client发来的数据", recvStr)
		conn.Close()
	}
}
