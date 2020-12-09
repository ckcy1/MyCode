package main

import "fmt"

type account struct {
	k     string  //输入数字
	note  string  //说明
	pay   float64 //收支金额
	money float64 // 账户金额
	str   string  // \n收支\t账户金额\t收支金额\t说明 “拼接字符”
	flag  bool    // 定义一个变量 ，记录是否有收支
}

func newaccount(k string, note string, pay float64,
	money float64, str string, flag bool) *account {
	return &account{
		k:     k,
		note:  note,
		pay:   pay,
		money: money,
		str:   str,
		flag:  flag}
}

func (a account) menu() {

	for {
		fmt.Println("\n---------------------家庭收支记账软件-----------------------------")
		fmt.Println("                       1.收支明细")
		fmt.Println("                       2.登记收入")
		fmt.Println("                       3.登记支出")
		fmt.Println("                       4.退    出")
		fmt.Printf("请选择（1-4）：")
		fmt.Scan(&a.k)
		switch a.k {
		case "1":

		case "2":
			a.shouru()
		case "3":
			a.zhichu()
		case "4":
a.exit()
		default:
			fmt.Println("你输入无效！")
		}
	}
	// breakTag:
	// 	fmt.Println("退出")

}
func (a *account) shouru() {
	fmt.Println("本次收入金额：")
	fmt.Scan(&a.pay)
	a.money += a.pay
	fmt.Printf("本次收入说明：")
	fmt.Scan(&a.note)
	// 字符串的拼接
	a.str += fmt.Sprintf("\n收入\t\t%v\t\t%v\t\t%v", a.money, a.pay, a.note)
	a.flag = false

}
func (a *account) zhichu() {
	fmt.Println("本次支出金额：")
	fmt.Scan(&a.pay)
	if a.pay > a.money {
		fmt.Println("余额不足")
		//break
	} else {
		a.money -= a.pay
		fmt.Printf("本次支出说明：")
		fmt.Scan(&a.note)
		// 字符串的拼接
		a.str += fmt.Sprintf("\n支出\t\t%v\t\t%v\t\t%v", a.money, a.pay, a.note)
		a.flag = false
	}

}
func (a *account) exit() {
	for {
		fmt.Printf("你确定要退出y/n?")
		fmt.Scan(&a.k)
		if a.k == "y" {
			goto breakTag
		} else if a.k == "n" {
			break
		}
	breakTag:
		fmt.Println("退出")
	}
}
func main() {
	a := newaccount("", "", 3.0, 100.0, "", false)
	a.menu()

}
