package main

import "fmt"

func main() {
	//最初始的赋值 := 也是一种赋值运算符
	// 定义两个变量 num1 = 5, num2 = 8, 交换两个变量的值
   num1 := 5 // 第一个格子
   num2 := 8 // 第二个格子
	fmt.Println(num1,num2)

   // 临时的,帮忙的变量
   num3 := num1 // 把num1内存盒子当中的数，赋值给num3，先保存下来
   num1 = num2 //
   num2 = num3
   fmt.Println(num1,num2)

   // 第二种方法
   // 变量定义的时候
   age1 , age2 := 15,18 // 多变量定义法
   fmt.Println(age1,age2)

   // go语言当中赋值运算符是支持多变量的
   age1, age2 = age2, age1 // 多值交换, 直接操作多变量
   fmt.Println(age1,age2)

   // 求和
   a := 5
   b := 8
   a = a + b // a= 13 b = 8
   b = a - b // a = 13 b = 5
   a = a - b // a = 13 a = 8

}
