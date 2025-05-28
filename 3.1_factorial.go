package main

func factorial_loop(n int) int {

	if n < 0 {
		panic("n不能小于0")
	} else if n < 2 {
		return 1
	}

	tmp := 1

	// for ; n > 1; n-- {
	// 	tmp = n * tmp
	// }

	for i := 2; i <= n; i++ {
		tmp = i * tmp
	}
	return tmp
}

func factorial_recuform(n int) int {
	switch {
	case n < 0:
		panic("项数不能小于0")
	case n < 2:
		return 1
	}

	return n * factorial_recuform(n-1)
}

func factorial_reculoop(n int, p int) int {

	switch {
	case n < 0:
		panic("项数不能小于0")
	case n < 2:
		return p

	}
	p = n * p

	return factorial_reculoop(n-1, p)
}
