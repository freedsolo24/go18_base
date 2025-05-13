package main

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

func fib_recuform(n int) int {
	switch {
	case n < 0:
		panic("项数不能为0")
	case n == 0:
		return 0
	case n < 3:
		return 1
	}
	return fib_recuform(n-1) + fib_recuform(n-2)
}

func fib_reculoop(n int, a, b int) int { // n是项数，a和b是两个初值
	switch {
	case n < 0:
		panic("项数不能为0")
	case n == 0:
		return 0
	case n < 3:
		return b
	}
	tmp := a + b
	a = b
	b = tmp
	return fib_reculoop(n-1, a, b)
}
