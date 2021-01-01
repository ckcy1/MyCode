// 使用“面向对象”的思维方式编写一个学生信息管理系统。
// 学生有id、姓名、年龄、分数等信息
// 程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能
package main

import (
	"fmt"
	"os"
)

type Student struct {
	id   int    //学生唯一ID
	name string //姓名
	age  int    //年龄
	mark int    // 分数
}

//Student  的构造函数
func NewStudent(id int, name string, age int, mark int) *Student {
	return &Student{
		id:   id,
		name: name,
		age:  age,
		mark: mark,
	}

}

type AllStudent struct {
	AllStudent []*Student //系统学生集合
}

func newAllStudent() *AllStudent {
	return &AllStudent{
		AllStudent: make([]*Student, 0, 100),
	}
}

//添加学生信息
func (a *AllStudent) addStudent(stu *Student) {
	a.AllStudent = append(a.AllStudent, stu)

}

//编辑学生信息
func (a *AllStudent) editStudent(stu *Student) {
	for i, v := range a.AllStudent {
		if v.id == stu.id {
			a.AllStudent[i] = stu
			return
		}
	}
	fmt.Printf("没有这人")
}

//展示学生列表
func (a *AllStudent) showStudent() {
	for _, v := range a.AllStudent {
		fmt.Println(v.name)
		fmt.Printf("学号：%d,姓名：%s，年龄：%d，分数：%d\n", v.id, v.name, v.age, v.mark)

	}
}

func main() {
	sk := newAllStudent()
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
			sk.showStudent()
			//展示学生列表
		case "2":
			stu := inputStudent()
			sk.addStudent(stu)
			//添加学生信息
		case "3":
			stu := inputStudent()
			sk.editStudent(stu)
			//编辑学生信息
		case "4":
			inputStudent()
			//删除学生信息
		case "5":
			//退出学生系统
			os.Exit(0)
		default:

		}
	}

}
