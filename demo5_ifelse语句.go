package main

import "fmt"

func main() {
	/**
	  if...else语句
	  if 条件表达式 {
	     // 条件成立，执行此处的代码
	  }else {
	    // 条件不成立，执行此处else的代码
	  }
	  年龄：0 - 120
	  如果男士年龄大于22周岁，则满足国家法定结婚年龄，可以结婚；否则不能结婚
	  城市的男女比例是1:4，往大城市里去。
	 */
	//else: 其他
	age := 22
	if age >= 22 {
		fmt.Println("该男士满足国家法定结婚年龄，可以着手准备彩礼结婚了")
	}else {
		fmt.Println("对不起，还不满足国家法定结婚年龄，请耐心等待国家分配")
	}
	fmt.Println("程序结束")

}
