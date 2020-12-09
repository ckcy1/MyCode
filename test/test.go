package main

import "fmt"

type account struct {
	k         string  //输入数字
	note      string  //说明
	pay       float64 //收支金额
	money     float64 // 账户金额
	str       string  // \n收支\t账户金额\t收支金额\t说明 “拼接字符”
	flag      bool    // 定义一个变量 ，记录是否有收支
	breakflag bool    // 跳出循环
}

func newaccount(k string, note string, pay float64,
	money float64, str string, flag bool, breakflag bool) *account {
	return &account{
		k:         k,
		note:      note,
		pay:       pay,
		money:     money,
		str:       str,
		flag:      flag,
		breakflag: breakflag,
	}
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
			a.breakflag = true
			break
		} else if a.k == "n" {
			break
		}

	}
}
func (a *account) mingxi() {
	if a.flag {
		fmt.Printf("当前没有收支明细...来一笔")
	} else {
		fmt.Println("-----------------------------当前收支明细记录-----------------------------")
		fmt.Printf(a.str)
	}

}
func main() {
	a := newaccount("", "", 3.0, 100.0, "\n收支\t账户金额\t收支金额\t\t说明", true, false)
	a.menu()

}
