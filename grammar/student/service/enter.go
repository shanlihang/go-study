package service

import (
	"fmt"
	"grammar/student/entity"
)
func Add() *entity.Student{
	var id int
	var name string
	var age int
	var gender string
	var phone string
	var email string
	fmt.Print("请输入学生id：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入学生年龄：")
	fmt.Scanln(&age)
	fmt.Print("请输入学生性别：")
	fmt.Scanln(&gender)
	fmt.Print("请输入学生手机号：")
	fmt.Scanln(&phone)
	fmt.Print("请输入学生邮箱：")
	fmt.Scanln(&email)
	return &entity.Student{
		Id:     id,
		Name:   name,
		Age:    age,
		Gender: gender,
		Phone:  phone,
		Email:  email,
	}
}

func Delete(id int,students map[int]entity.Student)  {
	delete(students,id)
}

func Update()*entity.Student{
	var id int
	var name string
	var age int
	var gender string
	var phone string
	var email string
	fmt.Print("请输入要修改学生的id：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)
	fmt.Print("请输入学生年龄：")
	fmt.Scanln(&age)
	fmt.Print("请输入学生性别：")
	fmt.Scanln(&gender)
	fmt.Print("请输入学生手机号：")
	fmt.Scanln(&phone)
	fmt.Print("请输入学生邮箱：")
	fmt.Scanln(&email)
	return &entity.Student{
		Id:     id,
		Name:   name,
		Age:    age,
		Gender: gender,
		Phone:  phone,
		Email:  email,
	}
}