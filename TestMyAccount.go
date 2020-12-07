package main

import "fmt"

var (
	k     string  //输入数字
	note  string  //说明
	pay   float64 //收支金额
	money float64 // 账户金额
	str   string  // \n收支\t账户金额\t收支金额\t说明 “拼接字符”
)

func main() {
	str :="\n收支\t账户金额\t收支金额\t\t说明"
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
			fmt.Println("-----------------------------当前收支明细记录-----------------------------")
			fmt.Printf(str)
		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scan(&pay)
			money += pay
			fmt.Printf("本次收入说明：")
			fmt.Scan(&note)
			str += fmt.Sprintf("\n收入\t\t%v\t\t%v\t\t%v",money,pay,note)
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scan(&pay)
			money += pay
			fmt.Printf("本次支出说明：")
			fmt.Scan(&note)
			str += fmt.Sprintf("\n支出\t\t%v\t\t%v\t\t%v",money,pay,note)
		case "4":
			goto breakTag
		default:
			fmt.Println("你输入无效！")
		}
	}
breakTag:
	fmt.Println("退出")

}
