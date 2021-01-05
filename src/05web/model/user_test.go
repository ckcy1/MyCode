package model

import (
	"fmt"
	"testing"
)

//func TestUser(t *testing.T) {
//	fmt.Println("开始测试USER中的相关方法")
//	t.Run("测试添加用户：",TestAddUser)
//}
func TestAddUser(t *testing.T) {
	fmt.Println("测试添加用户：")
	User:=&User{
		Username: "admin",
		Password: "123",
		Email:    "admin@qq.com",
	}
	User.AddUser()
}