package main

import "fmt"

func main() {
	fmt.Println("程序开始执行")
	age := 22 // 年龄
	// 中国法定结婚年龄：男22岁，女20
	if age >= 22 {
		fmt.Println("该男士满足法定结婚年龄，可以着手准备彩礼了")
	}
	fmt.Println("程序结束")

}
