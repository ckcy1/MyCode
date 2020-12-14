package utils

import (
	"fmt"
)

//Student 结构体
type Student struct {
	id   int    //学生唯一ID
	name string //姓名
	age  int    //年龄
	mark int    // 分数
}

//NewStudent  的构造函数
func NewStudent(id int, name string, age int, mark int) *Student {
	return &Student{
		id:   id,
		name: name,
		age:  age,
		mark: mark,
	}

}
//AllStudent 切片结构体
type AllStudent struct {
	AllStudent []*Student //系统学生集合
}

//NewAllStudent 是一个构造函数，初始化
func NewAllStudent() *AllStudent {
	return &AllStudent{
		AllStudent: make([]*Student, 0, 100),
	}
}

//AddStudent 添加学生信息
func (a *AllStudent) AddStudent(stu *Student) {
	a.AllStudent = append(a.AllStudent, stu)

}

//EditStudent 编辑学生信息
func (a *AllStudent) EditStudent(stu *Student) {
	for i, v := range a.AllStudent {
		if v.id == stu.id {
			a.AllStudent[i] = stu
			return
		}
	}
	fmt.Printf("没有这人")
}

// //输入学生信息
// func inputStudent() *Student {
// 	var (
// 		id   int    //学生唯一ID
// 		name string //姓名
// 		age  int    //年龄
// 		mark int    // 分数
// 	)
// 	fmt.Printf("请你输入学号：")
// 	fmt.Scan(&id)
// 	fmt.Printf("请你输入姓名：")
// 	fmt.Scan(&name)
// 	fmt.Printf("请你输入年龄：")
// 	fmt.Scan(&age)
// 	fmt.Printf("请你输入分数：")
// 	fmt.Scan(&mark)
// 	stu := NewStudent(id, name, age, mark)
// 	return stu
// }

//ShowStudent 展示学生列表
func (a *AllStudent) ShowStudent() {
	for _, v := range a.AllStudent {
		fmt.Println(v.name)
		fmt.Printf("学号：%d,姓名：%s，年龄：%d，分数：%d\n", v.id, v.name, v.age, v.mark)

	}
}
