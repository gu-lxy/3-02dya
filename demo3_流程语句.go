package main
import "fmt"
func main() {
	/**
	   if单词： 如果, 一种假设情况
	   if语句: 在程序当中，负责做出选择
	   比如： 如果肺炎疫情结束，我们就要去上学
	          如果考试不及格，就拿不到对应学分
	 */
	//grade: 成绩
	var grade int
	// 成绩大于60分算及格
	grade = 61 // 不及格 grade = 61 及格
	if grade >= 60 {  // 条件表达式位true ，则表示满足 if语句
		fmt.Printf("成绩是:%d, 成绩合格\n",grade)
	}
	fmt.Println("程序执行结束")


}
