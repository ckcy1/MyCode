package main

import (
	"fmt"
	"04customerManage/service"
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
		key:  "",
		flag: false,
		customerService:service.NewCustomerService(),
	}
}
//(*********)
func(c *customerView)list(){
	fmt.Println("----------------客户列表---------------")
	fmt.Printf("编号\t姓名\t性别\t年龄\t电话\t邮箱\n")
	cus :=c.customerService.List()
	for _, v := range cus {
		fmt.Printf("%d\t%s\t%s\t%d\t%s\t%s\n",v.ID,v.Name,v.Gender,
	v.Age,v.Phone,v.Emain)
	}
	fmt.Println("---------------客户列表完成--------------")
}

func(c *customerView)menu(){
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
			fmt.Println("1.添加客户")
		case "2":
			fmt.Println("2.修改客户")
		case "3":
			fmt.Println("3.删除客户")
		case "4":
			c.list()
			fmt.Println("4.客户列表")
		case "5":
			c.flag = true
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
