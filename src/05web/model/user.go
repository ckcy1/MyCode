package model

import (
	"05web/utils"
	"fmt"
)

type  User struct {
	ID int
	Username string
	Password string
	Email string
}

func (u *User) AddUser()error  {
// 写SQL语句
	sqlStr :="insert into users(username,password,email)values (?,?,?)"
	stmt,err :=utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常：",err)
		return err
	}
	_,err = stmt.Exec(u.Username, u.Password, u.Email)
	if err != nil {
		fmt.Println("执行出现异常：",err)
		return err
	}
	return nil
}