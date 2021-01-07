package model

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("开始测试USER中的相关方法")
	t.Run("测试添加用户：", testGetUsers)
}
func testAddUser(t *testing.T) {
	fmt.Println("测试添加用户：")
	User := &User{
		Username: "admin11",
		Password: "12311",
		Email:    "admin11@qq.com",
	}
	User.AddUser()
}
func testGetUserById(t *testing.T) {
	fmt.Println("查询用户：")
	user := User{
		ID: 3,
	}
	u, err := user.GetUserById()
	if err != nil {
		fmt.Println("出错拉！！！")
	}
	fmt.Printf("我的ID：%d,姓名：%s,密码：%s,邮箱：%s.\n", u.ID, u.Username, u.Password, u.Email)
}
func testGetUsers(t *testing.T) {
	fmt.Println("查询多用户")
	user := &User{}
	u, _ := user.GetUsers()
	for k, v := range u {
		fmt.Printf("第%v个用户是%v:\n", k+1, v)
	}
}
