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
		fmt.Println("                       5.退出菜单")
		fmt.Println("请你输入菜单：")
		fmt.Scan(&k)
		switch k {
		case "1":
			fmt.Println("菜单一1")
		case "2":
			fmt.Println("菜单一2")
		case "3":
			fmt.Println("菜单一3")
		case "4":
			fmt.Println("菜单一4")
		case "5":
			goto breakTag
		default:
			fmt.Println("你输入无效！")
		}
	}
breakTag:
	fmt.Println("退出")

}
