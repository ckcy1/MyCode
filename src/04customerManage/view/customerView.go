package main

import (
	"04customerManage/model"
	"04customerManage/service"
	"fmt"
)

type customerView struct {
	key  string
	flag bool
	//(*********)customerView为什么要加字段customerService
	//调用customerService 里头List方法
	customerService *service.CustomerService
}

//newcustomerView 构造函数初始化
func newcustomerView() *customerView {
	return &customerView{
		key:             "",
		flag:            false,
		customerService: service.NewCustomerService(),
	}
}

//(*********)
func (c *customerView) list() {
	fmt.Println("----------------客户列表---------------")
	fmt.Printf("编号\t姓名\t性别\t年龄\t电话\t邮箱\n")
	cus := c.customerService.List()
	for _, v := range cus {
		// 	fmt.Printf("%d\t%s\t%s\t%d\t%s\t%s\n",v.ID,v.Name,v.Gender,
		// v.Age,v.Phone,v.Emain)
		fmt.Printf(v.GetInfo())
	}
	fmt.Println("---------------客户列表完成--------------")
}
func (c *customerView) add() {
	var (
		name   string
		gender string
		age    int
		phone  string
		emain  string
	)
	fmt.Println("----------------添加客户---------------")
	fmt.Println("姓名：")
	_, _ = fmt.Scan(&name)
	fmt.Println("性别：")
	_, _ = fmt.Scan(&gender)
	fmt.Println("年龄：")
	_, _ = fmt.Scan(&age)
	fmt.Println("电话：")
	_, _ = fmt.Scan(&phone)
	fmt.Println("邮箱：")
	_, _ = fmt.Scan(&emain)
	cus := model.NewCustomer2(name, gender, age, phone, emain)
	if c.customerService.Add(cus) {
		fmt.Println("----------------添加完成---------------")
	} else {
		fmt.Println("----------------添加失败---------------")
	}
}
func (c *customerView) delete() {
	fmt.Println("----------------删除客户---------------")
	fmt.Println("请选择待删除客户编号（-1退出）：")
	id := -1
	fmt.Scan(&id)
	fmt.Println("确认是否删除（Y/N）:")
	choice := ""
	fmt.Scan(&choice)
	if choice == "y" || choice == "Y" {
		if c.customerService.Delete(id) {
			fmt.Println("----------------删除成功---------------")
		} else {
			fmt.Println("----------------删除失败，输入ID 不存在-")
		}
	}

}
func (c *customerView) exit() {
	fmt.Println("确定是否退出（Y/N）:")
	for {
		fmt.Scan(&c.key)
		if c.key == "Y" || c.key == "y" || c.key == "n" || c.key == "N" {
			break
		}
		fmt.Println("你的输入有误，确认是否退出（Y/N）")

	}
	if c.key == "Y" || c.key == "y" {
		c.flag = true
	}
}

func (c *customerView) menu() {
	for {
		fmt.Println("----------------客户信息管理软件---------------")
		fmt.Println("                  1.添加客户")
		fmt.Println("                  2.修改客户")
		fmt.Println("                  3.删除客户")
		fmt.Println("                  4.客户列表")
		fmt.Println("                  5.退    出")
		fmt.Println("请选择（1-5）：")
		fmt.Scan(&c.key)
		switch c.key {
		case "1":
			c.add()
		case "2":
			fmt.Println("2.修改客户")
		case "3":
			c.delete()
		case "4":
			c.list()
		case "5":
			c.exit()
			fmt.Println("你已经退出了软件！")
		default:
			fmt.Printf("你的输入有误！！\n")
		}
		if c.flag {
			break
		}
	}

}
func main() {
	cus := newcustomerView()
	cus.menu()

}
