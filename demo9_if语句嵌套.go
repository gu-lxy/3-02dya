package main

import "fmt"

func main() {
	/**
	   if 条件表达式 {
	       //执行的代码
	       if 条件表达式{
			} else{
	        }
	    }else {
	        //执行的代码
			if 条件表达式 {
	        }
	     }
	 */
	//结婚年龄 ：国家法定结婚年龄男生不得低于22周岁，女生不得低于20周岁
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age) // 接受键盘输入,并把输入的值赋值变量age
	var sex int // 1表示男生 2表示女生
	fmt.Println("请输入你的性别（男生请输入1，女生请输入2）：")
	fmt.Scanln(&sex) // 接受键盘的输入，并把输入的值赋值给变量sex

	//1、先判断性别
	if sex == 1 { // 男的
		if age >= 22 {
			fmt.Println("恭喜你小哥哥，可以领证结婚了")
		}else{
			fmt.Println("请耐心等待国家分配")
		}
	}else { //除了男的，剩下的都是女的
		if age >= 20 {
			fmt.Println("恭喜你小姐姐，可以领证结婚了")
		} else {
			fmt.Println("耐心你的小哥哥出现吧")
		}
	}
}