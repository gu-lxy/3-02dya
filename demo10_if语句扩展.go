package main

import "fmt"

func main() {
	var sex int
	fmt.Println("请输入你的性别：")
	fmt.Println("友情提示：1表示男，2表示女，3表示未知，请如实填写。")
	fmt.Scanln(&sex) // 读取键盘输入，并把值赋值给sex变量
	// 1表示男，2表示女，3表示未知
	if sex == 1{
		fmt.Println("性别为男")
	}else if sex ==2 {
		fmt.Println("性别为女")
	}else if sex == 3 {
		fmt.Println("性别未知")
	}
	/**
	  if语句语法：
	     1、 if 表达式 {}
	     2、 if 表达式 {
	         }else {
	         }
	     3、if 表达式的嵌套
	        if 表达式{
	           if 表达式{}
	        }
	    4、if 表达式{
	       }else if 表达式{
	       }...
	 */
}
