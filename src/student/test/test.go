// 使用“面向对象”的思维方式编写一个学生信息管理系统。
// 学生有id、姓名、年龄、分数等信息
// 程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能
package main

import "fmt"

type student struct {
	id   int    //学生唯一ID
	name string //姓名
	age  int    //年龄
	mark int    // 分数
}

func main() {
	// 展示学生信息管理系统
	for {
		fmt.Println("学生信息管理系统")
		fmt.Println("1.展示学生列表")
		fmt.Println("2.添加学生信息")
		fmt.Println("3.编辑学生信息")
		fmt.Println("4.删除学生信息")
		fmt.Println("5.退出学生系统")
		//输入学生信息
		//执行输入信息
		k := ""
		switch k {
		case "1":

		case "2":

		case "3":

		case "4":

		case "5":

		default:

		}
	}

}
