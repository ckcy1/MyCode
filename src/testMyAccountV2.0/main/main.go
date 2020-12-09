package main

import "testMyAccountV2.0/utils"

func main() {
	// 	a := utils.newAccount("", "", 3.0, 100.0, "\n收支\t账户金额\t收支金额\t\t说明", true, false)
	// 	a.menu()
	utils.NewAccount().Menu()
}
