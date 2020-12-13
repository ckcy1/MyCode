package service

import (
	"04customerManage/model"
	"fmt"
)

//CustomerService 构造函数
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段后面，还可以作为新客户的ID +1
	num int
}

//NewCustomerService 构造函数初始化 返回一个实例（*******）
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.num = 1
	customer := model.NewCustomer(1, "张三", "男", 33, "112", "zhang@qq.com")
	customerService.customers = append(customerService.customers, customer)
	// c:=model.NewCustomer(1,"张三","男", 33 ,"112","zhang@qq.com")
	// return &CustomerService{
	// 	Customers : []model.Customer{1,"张三","男", 33 ,"112","zhang@qq.com"}
	// }
	return customerService
}

//List 返回客户信息，其实就是切片
func (c *CustomerService) List() []model.Customer {
	return c.customers
}

//Add 将新的客户加入到切片
func (c *CustomerService) Add(cus model.Customer) bool {
	// CustomerService.num 成为计算器（***）
	c.num++
	cus.ID = c.num
	// cus.ID =1
	// cus.ID++
	c.customers = append(c.customers, cus)
	return true
}

//Findbyid 返回ID号对应的切片的下标,如果没有客户 返回-1
func (c *CustomerService) Findbyid(id int) int {
	index := -1 //(************)
	for i := 0; i < len(c.customers); i++ {
		if c.customers[i].ID == id {
			// index = id 错误（*****） 找下标
			index = i
		}
	}
	return index
}

//Delete 删除一个客户
func (c *CustomerService) Delete(id int) bool {
	index := c.Findbyid(id)
	if index == -1 {
		fmt.Println("删除失败")
		return false
	}
	// 删除一个切片的元素  相当于删除客户
	c.customers = append(c.customers[:index], c.customers[index+1:]...)
	return true
}
