package model

//Customer 是客户信息结构体
type Customer struct {
	ID     int    // 客户ID 唯一
	Name   string // 姓名
	Gender string //性别
	Age    int    // 年龄
	Phone  string // 电话
	Emain  string // 电子邮件
}

//NewCustomer 构造函数
func NewCustomer(id int, name string, gender string,
	age int, phone string, emain string) Customer {
	return Customer{
		ID:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Emain:  emain,
	}
}

//NewCustomer2 不带ID的构造函数
func NewCustomer2(name string, gender string,
	age int, phone string, emain string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Emain:  emain,
	}
}
