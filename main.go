package main

import (
	"fmt"
)

func main() {
	fmt.Println("==========九 九 乘 法 表==========")
	table9()
	fmt.Println("==========随机数累加累乘==========")
	random_addmul()
	fmt.Println("==========斐波那契额循环==========")
	for i := 0; ; i++ {
		if n := fib_loop(i); n < 100 {
			fmt.Printf("fib(%d)=%d\n", i, n)
		} else {
			break
		}
	}
	fmt.Println("==========斐波那契额递归==========")
	for i := 0; ; i++ {
		if n := fib_recuform(i); n < 100 {
			fmt.Printf("fib(%d)=%d\n", i, n)
		} else {
			break
		}
	}
	fmt.Println("==========斐波那契额高阶==========")
	for i := 0; ; i++ {
		if n := fib_reculoop(i, 1, 1); n < 100 {
			fmt.Printf("fib(%d)=%d\n", i, n)
		} else {
			break
		}
	}
}
