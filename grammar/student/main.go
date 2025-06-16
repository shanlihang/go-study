package main

import (
	"fmt"
	"grammar/student/entity"
	"grammar/student/service"
)
var username string
var password string
var flag int
var list = make(map[int]entity.Student)

func main() {
	fmt.Println("===欢迎使用学生管理系统===")
	fmt.Print("请输入用户名：")
	fmt.Scanln(&username)
	fmt.Print("请输入密码：")
	fmt.Scanln(&password)
	if(username == "admin" && password == "123456"){
		fmt.Println("登录成功")
		for{
			fmt.Println("请选择功能")
			fmt.Println("1. 查看所有学生")
			fmt.Println("2. 新增学生")
			fmt.Println("3. 删除学生")
			fmt.Println("4. 修改学生")
			fmt.Print("请选择操作:")
			fmt.Scanln(&flag)
			switch flag{
			case 1:
				for _, student := range list{
					fmt.Println(student)
				}
			case 2:
				student := service.Add()
				list[student.Id] = *student
			case 3:
				var id int
				fmt.Print("请输入要删除的学生id:")
				fmt.Scanln(&id)
				service.Delete(id,list)
			case 4:
				student := service.Update()
				for i,v := range list{
					if v.Id == student.Id {
						list[i] = *student
					}
				}
			}
		}
		
	} else {
		fmt.Println("登录失败")
	}
	return
}