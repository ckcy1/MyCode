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
//Add 将新的客户加入到切片
func (c *CustomerService) Add(cus model.Customer){
	// CustomerService.num 成为计算器（***）
	c.num++
	cus.ID =c.num
	// cus.ID =1
	// cus.ID++
	c.customers = append(c.customers,cus)

}
