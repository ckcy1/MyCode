package main

import "fmt"

// var m string

// func main() {

// 	fmt.Println("请输入：")
// 	fmt.Scan(&m)
// 	fmt.Println(m)
// }
func main() {
	var a *int
	a = new(int)

	fmt.Println(*a)
	fmt.Println(a)

}
