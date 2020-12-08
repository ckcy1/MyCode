package main

import "fmt"

var (
	k     string  //输入数字
	note  string  //说明
	pay   float64 //收支金额
	money float64 // 账户金额
	str   string  // \n收支\t账户金额\t收支金额\t说明 “拼接字符”
	flag  bool    // 定义一个变量 ，记录是否有收支
)

func main() {
	str := "\n收支\t账户金额\t收支金额\t\t说明"
	money := 10000.0
	pay := 0.0
	flag := true
	for {
		fmt.Println("\n---------------------家庭收支记账软件-----------------------------")
		fmt.Println("                       1.收支明细")
		fmt.Println("                       2.登记收入")
		fmt.Println("                       3.登记支出")
		fmt.Println("                       4.退    出")
		fmt.Printf("请选择（1-4）：")
		fmt.Scan(&k)
		switch k {
		case "1":
			if flag {
				fmt.Printf("当前没有收支明细...来一笔")
			} else {
				fmt.Println("-----------------------------当前收支明细记录-----------------------------")
				fmt.Printf(str)
			}

		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scan(&pay)
			money += pay
			fmt.Printf("本次收入说明：")
			fmt.Scan(&note)
			// 字符串的拼接
			str += fmt.Sprintf("\n收入\t\t%v\t\t%v\t\t%v", money, pay, note)
			flag = false

		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scan(&pay)
			if pay > money {
				fmt.Println("余额不足")
				break
			}
			money -= pay
			fmt.Printf("本次支出说明：")
			fmt.Scan(&note)
			// 字符串的拼接
			str += fmt.Sprintf("\n支出\t\t%v\t\t%v\t\t%v", money, pay, note)
			flag = false
		case "4":
			for {
				fmt.Printf("你确定要退出y/n?")
				fmt.Scan(&k)
				if k == "y" {
					goto breakTag
				} else if k == "n" {
					break
				}
			}
		default:
			fmt.Println("你输入无效！")
		}
	}
breakTag:
	fmt.Println("退出")

}
