package service

import (
	"04customerManage/model"
)

//CustomerService 构造函数
type CustomerService struct {
	customers []model.Customer
	num       int
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
