package utils

import "fmt"

type account struct {
	k     string  //输入数字
	note  string  //说明
	pay   float64 //收支金额
	money float64 // 账户金额
	str   string  // \n收支\t账户金额\t收支金额\t说明 “拼接字符”
	flag  bool    // 定义一个变量 ，记录是否有收支
}

func (a account) menu() {

	k := a.k
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

		case "2":

		case "3":

		case "4":

		default:
			fmt.Println("你输入无效！")
		}
	}
	// breakTag:
	// 	fmt.Println("退出")

}
