package main

import "fmt"

var k string

func main() {
	for {
		fmt.Println(".........................主菜单.........................")
		fmt.Println("                       第一级菜单")
		fmt.Println("                       第二级菜单")
		fmt.Println("                       第三级菜单")
		fmt.Println("                       第四级菜单")
		fmt.Println("请你输入菜单：")
		fmt.Scan(&k)
	}
}
