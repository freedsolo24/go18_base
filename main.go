package main

import (
	"fmt"
)

func main() {
	fmt.Println("==========第一章：九 九 乘 法 表==========")
	table9()
	fmt.Println("==========第一章：随机数累加累乘==========")
	random_addmul()
	fmt.Println("==========第一章：斐波那契额循环==========")
	for i := 0; ; i++ {
		if n := fib_loop(i); n < 100 {
			fmt.Printf("fib(%d)=%d\n", i, n)
		} else {
			break
		}
	}
	fmt.Println("==========第一章：斐波那契额递归==========")
	for i := 0; ; i++ {
		if n := fib_recuform(i); n < 100 {
			fmt.Printf("fib(%d)=%d\n", i, n)
		} else {
			break
		}
	}
	fmt.Println("==========第一章：斐波那契额高阶==========")
	for i := 0; ; i++ {
		if n := fib_reculoop(i, 1, 1); n < 100 {
			fmt.Printf("fib(%d)=%d\n", i, n)
		} else {
			break
		}
	}
	fmt.Println("==========第二章：打印切片内容============")
	s1 := make([]int, 3, 4)
	str1 := "初始s1切片情况:"
	fmt.Printf("|%-28s|%s", str1, getSliceDetails(s1))
	s2 := append(s1, 1)
	str2 := "在s1切片基础上追加元素变成s2切片:"
	fmt.Printf("|%-22s|%s", str2, getSliceDetails(s2))
	s2 = append(s2, 100)
	str3 := "在s2切片基础上追加元素扩容s2切片:"
	fmt.Printf("|%-22s|%s", str3, getSliceDetails(s2))
	// 一个字符算一个宽度，不管中文，英文，标点

	fmt.Println("==========第二章：数组转成切片============")
	fmt.Printf("数组转成切片，类型:%[1]T, 内容:%[1]v\n", arraySlice())

}
