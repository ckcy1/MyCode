/*
引用包包 import  方法 跟 函数明要
第一个字母要大写，同时要加// 注释
注释 要 方法名 跟函数名 上面  加空格
*/

package utils

import "fmt"

// Account 结构体
type Account struct {
	k         string  //输入数字
	note      string  //说明
	pay       float64 //收支金额
	money     float64 // 账户金额
	str       string  // \n收支\t账户金额\t收支金额\t说明 “拼接字符”
	flag      bool    // 定义一个变量 ，记录是否有收支
	breakflag bool    // 跳出循环
}

//NewAccount GONGC
// func newAccount(k string, note string, pay float64,
// 	money float64, str string, flag bool, breakflag bool) *Account {
// 	return &Account{
// 		k:         k,
// 		note:      note,
// 		pay:       pay,
// 		money:     money,
// 		str:       str,
// 		flag:      flag,
// 		breakflag: breakflag,
// 	}
// }

//Menu 主菜单
func (a *Account) Menu() {

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
			a.mingxi()
		case "2":
			a.shouru()
		case "3":
			a.zhichu()
		case "4":
			a.exit()
		default:
			fmt.Println("你输入无效！")
		}
		if a.breakflag {
			fmt.Println("退出")
			break
		}
	}
	// breakTag:
	// 	fmt.Println("退出")

}
func (a *Account) shouru() {
	fmt.Println("本次收入金额：")
	fmt.Scan(&a.pay)
	a.money += a.pay
	fmt.Printf("本次收入说明：")
	fmt.Scan(&a.note)
	// 字符串的拼接
	a.str += fmt.Sprintf("\n收入\t\t%v\t\t%v\t\t%v", a.money, a.pay, a.note)
	a.flag = false

}
func (a *Account) zhichu() {
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
func (a *Account) exit() {
	for {
		fmt.Printf("你确定要退出y/n?")
		fmt.Scan(&a.k)
		if a.k == "y" {
			a.breakflag = true
			break
		} else if a.k == "n" {
			break
		}

	}
}
func (a *Account) mingxi() {
	if a.flag {
		fmt.Printf("当前没有收支明细...来一笔")
	} else {
		fmt.Println("-----------------------------当前收支明细记录-----------------------------")
		fmt.Printf(a.str)
	}
}

//NewAccount 是个构造函数
func NewAccount() *Account {

	return &Account{
		k:         "",
		note:      "",
		pay:       0.0,
		money:     1000.0,
		str:       "\n收支\t账户金额\t收支金额\t\t说明",
		flag:      true,
		breakflag: false,
	}
}
