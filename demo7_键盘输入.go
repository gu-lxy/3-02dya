package main

import "fmt"

func main() {

	//age := 19 // 19是编写程序时写死的一个数据
	var age int // 默认值是0

	// 需求：在程序执行的时候，程序开发者自己输入一个变量值，然后程序执行判断，而不是在程序中把该值固定
	// 在go语言当中，提供了这种允许程序人员输入具体值得软件包
	// fmt包可以实现

	/**
	   fmt : go语言中官方提供的一个软件包,  主要用于实现格式化打印和输入/输出相关的操作.
	   单词Input：输入
	    单词Output：输出
	     IO操作
	 */
	//fmt包： 用于读取编程人员在键盘上的输入
	//sacn: 横扫、观察
	fmt.Scanln(&age) // 读取键盘上的输入，并把相应的值赋值给age变量
	fmt.Println(&age)
	fmt.Println(age) // 23

	if age >= 22 {
		fmt.Println("可以结婚了")
	}else {
		fmt.Println("对不起，还要再等等")
	}
}
