package main

// fibCache变量是map

var fibCache = map[int]int{0: 0, 1: 1, 2: 1}

func fibUseCache(n int) int {
	switch {
	case n < 0:
		panic("n不能小于0")
	case n == 1:
		return fibCache[1]
	case n == 2:
		return fibCache[2]
	}
	// 简单容易理解的逻辑的写法
	// if v,ok:=fibCache[n];ok {
	// 	return v
	// } else {
	// 	v:=fibUseCache(n-1)+fibUseCache(n-2)
	// 	fibCache[n]=v
	// 	return v
	// }
	// 把v抽出来
	// if _,ok:=fibCache[n];ok {
	// 	return fibCache[n]
	// } else {
	// 	fibCache[n]=fibUseCache(n-1)+fibUseCache(n-2)
	// 	return	fibCache[n]
	// }
	// 把return抽出来
	// fmt.Printf("刚开始: fib(%d)=%d\n", n, fibCache[n])
	if _, ok := fibCache[n]; !ok {
		fibCache[n] = fibUseCache(n-1) + fibUseCache(n-2)
	}

	return fibCache[n]

}
