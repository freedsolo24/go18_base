package main

import (
	"fmt"
	"math/rand"
	"time"
)

func table9() {
	var width int
	/* for i:=1;i<=9;i++ {
	      for j:=1;j<=9;j++ {
	        if j==1 {
	   	   width=2
	   	   fmt.Printf("%dx%d=%-[4]*[3]d",j,i,i*j,width)
	   	 } else {
	   	   width=3
	   	   fmt.Printf("%dx%d=%-[4]*[3]d",j,i,i*j,width)
	   	 }
	      }
	   }
	*/
	for i := 1; i <= 9; i++ {
		width = 2
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%-[4]*[3]d", j, i, i*j, width)
			if j == 1 {
				width++
			}
		}
		fmt.Println()
	}
}

func random_addmul() {

	seed := time.Now().UnixNano() // 随机数种子
	src := rand.NewSource(seed)   // 随机数发生源
	r := rand.New(src)            // 随机数生成器

	var sum int = 0
	var mul uint64 = 1
	for i := 0; i < 10; i++ {
		random := r.Intn(20) + 1
		if random&1 == 1 {
			sum += random // 任何数与1与，结果是1，是奇数
		} else {
			mul *= uint64(random)
		}
		fmt.Printf("%d  ", random)
	}
	fmt.Printf("\n奇数随机数的累和=%d,偶数随机数的累乘=%d\n", sum, mul)
}

func fib_loop(n int) int {
	switch {
	case n < 0:
		panic("项数不能为负数")
	case n == 0:
		return 0
	case n == 1:
		return 1
	case n == 2:
		return 1
	}
	a, b := 1, 1
	for i := 0; i < n-2; i++ {
		tmp := a + b
		a = b
		b = tmp
	}
	return b // 最后return b，函数体消亡
}

func fib_rcrs(n int) int {
	switch {
	case n < 0:
		panic("项数不能为0")
	case n == 0:
		return 0
	case n < 3:
		return 1
	}
	return fib_rcrs(n-1) + fib_rcrs(n-2)
}

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
		if n := fib_rcrs(i); n < 100 {
			fmt.Printf("fib(%d)=%d\n", i, n)
		} else {
			break
		}
	}
}
