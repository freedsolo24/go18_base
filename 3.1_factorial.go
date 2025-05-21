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

func factorial_recu(n int) int {

	return 0
}
