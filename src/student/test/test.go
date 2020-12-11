// 使用“面向对象”的思维方式编写一个学生信息管理系统。
// 学生有id、姓名、年龄、分数等信息
// 程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能
package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int    //学生唯一ID
	name string //姓名
	age  int    //年龄
	mark int    // 分数
}

//student  的构造函数
func newstudent(id int, name string, age int, mark int) *student {
	return &student{
		id:   id,
		name: name,
		age:  age,
		mark: mark,
	}

}

type allstudent struct {
	allstudent []*student //系统学生集合
}

func newallstudent() *allstudent {
	return &allstudent{
		allstudent: make([]*student, 0, 100),
	}
}

//添加学生信息
func (a *allstudent) addstudent(stu *student) {
	a.allstudent = append(a.allstudent, stu)

}

//编辑学生信息
func (a *allstudent) editstudent(stu *student) {
	for i, v := range a.allstudent {
		if v.id == stu.id {
			a.allstudent[i] = stu
			return
		}
	}
	fmt.Printf("没有这人")
}

//输入学生信息
func inputstudent() *student {
	var (
		id   int    //学生唯一ID
		name string //姓名
		age  int    //年龄
		mark int    // 分数
	)
	fmt.Printf("请你输入学号：")
	fmt.Scan(&id)
	fmt.Printf("请你输入姓名：")
	fmt.Scan(&name)
	fmt.Printf("请你输入年龄：")
	fmt.Scan(&age)
	fmt.Printf("请你输入分数：")
	fmt.Scan(&mark)
	stu := newstudent(id, name, age, mark)
	return stu
}

//展示学生列表
func (a *allstudent) showstudent() {
	for _, v := range a.allstudent {
		fmt.Println(v.name)
		fmt.Printf("学号：%d,姓名：%s，年龄：%d，分数：%d\n",v.id, v.name, v.age, v.mark)

	}
}

func main() {
	sk := newallstudent()
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
		fmt.Scan(&k)
		switch k {
		case "1":
			sk.showstudent()
			//展示学生列表
		case "2":
			stu := inputstudent()
			sk.addstudent(stu)
			//添加学生信息
		case "3":
			stu := inputstudent()
			sk.editstudent(stu)
			//编辑学生信息
		case "4":
			inputstudent()
			//删除学生信息
		case "5":
			//退出学生系统
			os.Exit(0)
		default:

		}
	}

}
