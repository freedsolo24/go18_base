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
	fmt.Printf("初始的s1切片情况:                  %s\n", getSliceDetails(s1))
	s2 := append(s1, 1)
	fmt.Printf("在s1切片基础上追加元素变成新s2切片: %s\n", getSliceDetails(s2))
	s2 = append(s2, 100)
	fmt.Printf("在s2切片基础上追加元素扩容s2切片:   %s\n", getSliceDetails(s2))
}
