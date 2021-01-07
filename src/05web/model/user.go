package model

import (
	"05web/utils"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *User) AddUser() error {
	// 写SQL语句
	sqlStr := "insert into users(username,password,email)values (?,?,?)"
	stmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常：", err)
		return err
	}
	_, err = stmt.Exec(u.Username, u.Password, u.Email)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}
func (u *User) GetUserById() (*User, error) {
	sqlStr := "select id,username,password,email from users where id =?"
	row := utils.Db.QueryRow(sqlStr, u.ID)
	var (
		id       int
		username string
		password string
		email    string
	)
	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		return nil, err
	}
	user := &User{
		ID:       id,
		Username: username,
		Password: password,
		Email:    email,
	}
	return user, err
}
func (u *User) GetUsers() ([]*User, error) {
	sqlStr := "select id,username,password,email from users"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var users []*User
	for rows.Next() {
		var (
			id       int
			username string
			password string
			email    string
		)
		err := rows.Scan(&id, &username, &password, &email)
		if err != nil {
			return nil, err
		}
		user := &User{
			ID:       id,
			Username: username,
			Password: password,
			Email:    email,
		}
		users = append(users, user)
	}
	return users, nil
}
