package main

import "fmt"

func main() {
	// 给定一个变量num, 求输出该变量的奇偶性
	num := 20
	//奇数：不能被2整除的是奇数
	//偶数：能被2整除的是偶数
	// 判断： num /2 取余数，看余数是否为0，如果余数为0，则能够整除，是偶数；否则，则不能被整除，就是奇数
	// 取余就是计算机中的取模运算
	if num % 2 == 0 {
		fmt.Println("该变量是一个偶数")
	}else {
		fmt.Println("该变量是一个奇数")
	}
	fmt.Println("程序执行结束")
}
