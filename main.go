package main

import (
	"fmt"
)

func main() {
	fmt.Println("==========高阶函数==========")
	x := PUT(1, 2, func(x, y int) int {
		return x + y
	})
	fmt.Printf("调用高阶函数: %d\n", x)
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

	fmt.Println("==========第二章：随机数重复次数============")
	numsRepeat()

	fmt.Println("==========第三章：阶乘循环============")
	fmt.Printf("%d的阶乘=%d\n", 3, factorial_loop(3))

	fmt.Println("==========第三章：阶乘递归============")
	fmt.Printf("%d的阶乘=%d\n", 4, factorial_recuform(4))

	fmt.Println("==========第三章: 阶乘循环递归============")
	fmt.Printf("%d的阶乘=%d\n", 5, factorial_reculoop(5, 1))

	fmt.Println("==========第三章: 01.打印正数矩阵============")
	countup(3)
	fmt.Println("==========第三章: 02.打印倒数矩阵============")
	countdown(3)
	fmt.Println("==========第三章: 03.打印左下三角正数============")
	leftdownTriangleCountup(3)
	fmt.Println("==========第三章: 04.打印左下三角倒数============")
	leftdownTriangleCountdown(3)
	fmt.Println("==========第三章: 05.打印左上三角正数============")
	leftupTriangleCountup(3)
	fmt.Println("==========第三章: 06.打印左上三角倒数============")
	leftupTriangleCountdown(3)
	fmt.Println("==========第三章: 07.打印右下三角倒数============")
	rightdownTriangleCountdown(3)
	fmt.Println("==========第三章: 08.打印右下三角正数============")
	rightdownTriangleCountup(3)
	fmt.Println("==========第三章: 09.打印右下三角倒数切片解法=====")
	rightdownTriangleCountdownSlice(3)
	fmt.Println("==========第四章: 01.求面积=====")
	circle := NewCircle(4)
	triangle := NewTriangle(2, 2)
	rectangle := NewRectangle(2, 3)
	var s Interfacer
	s = circle
	fmt.Printf("圆形的面积是:  %f\n", s.area())
	s = triangle
	fmt.Printf("三角形的面积是:%f\n", s.area())
	s = rectangle
	fmt.Printf("长方形的面积是:%f\n", s.area())
	fmt.Println("==========在理解一下高阶函数的用法=====")

	// 函数的签名要一致,里面定义我想让他怎么做,这是一个匿函,调用之后就消亡
	// 加法要这么做
	i := highorder(3, 4, func(a, b int) int {
		z := a + b
		return z
	})
	fmt.Printf("求加法:%d\n", i)

	// 乘法要这么做
	j := highorder(5, 6, func(a, b int) int {
		z := a * b
		return z
	})
	fmt.Printf("求乘法:%d\n", j)
}
