package main

// 再次理解高阶函数, 形参里面是一个函数,只需要定义传什么形参,返什么返回值
// 具体形参函数怎么做我不管,只管把我拿到的传参转给你

/*
calc func(int int) int { z:=a+b return z}
*/
func highorder(x, y int, calc func(int, int) int) int {
	z := calc(x, y)
	return z
}
