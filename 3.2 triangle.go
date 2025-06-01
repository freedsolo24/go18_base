package main

import "fmt"

func countup(n int) {
	switch {
	case n < 0:
		panic("n不能小于0")
	case n == 0:
		panic("n不能等于0")
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}

func countdown(n int) {
	switch {
	case n < 0:
		panic("n不能小于0")
	case n == 0:
		panic("n不能等于0")
	}
	for i := 1; i <= n; i++ {
		for j := n; j >= 1; j-- {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}

func leftdownTriangleCountup(n int) {
	switch {
	case n < 0:
		panic("n不能小于0")
	case n == 0:
		panic("n不能等于0")
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}

func leftdownTriangleCountdown(n int) {
	switch {
	case n < 0:
		panic("n不能小于0")
	case n == 0:
		panic("n不能等于0")
	}
	for i := 1; i <= n; i++ {
		for j := i; j >= 1; j-- { // 倒数应该是--
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}

}

func leftupTriangleCountup(n int) {
	switch {
	case n < 0:
		panic("n不能小于0")
	case n == 0:
		panic("n不能等于0")
	}
	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}
func leftupTriangleCountdown(n int) {
	switch {
	case n < 0:
		panic("n不能小于0")
	case n == 0:
		panic("n不能等于0")
	}
	for i := 1; i <= n; i++ {
		for j := n; j >= i; j-- {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}
func rightdownTriangleCountdown(n int) {
	switch {
	case n < 0:
		panic("n不能小于0")
	case n == 0:
		panic("n不能等于0")
	}
	lastline := ""
	for i := n; i >= 1; i-- {
		s := fmt.Sprintf("%d ", i)
		lastline = lastline + s
	}
	width := len(lastline)

	for i := 1; i <= n; i++ {
		line := ""
		for j := i; j >= 1; j-- {
			s := fmt.Sprintf("%d ", j)
			line = line + s
		}
		fmt.Printf("%[2]*[1]s\n", line, width)
	}
}

func rightdownTriangleCountup(n int) {
	switch {
	case n < 0:
		panic("n不能小于0")
	case n == 0:
		panic("n不能等于0")
	}
	lastline := ""
	for i := 3; i >= 1; i-- {
		s := fmt.Sprintf("%d ", i)
		lastline += s
	}
	width := len(lastline)

	for i := 1; i <= n; i++ {
		line := ""
		for j := 1; j <= i; j++ {
			s := fmt.Sprintf("%d ", j)
			line += s
		}
		fmt.Printf("%[2]*[1]s\n", line, width)
	}
}
func rightdownTriangleCountdownSlice(n int) {
	lastline := ""

	for i := n; i >= 1; i-- {
		if i == 1 {
			lastline += fmt.Sprintf("%d", i)
		} else {
			s := fmt.Sprintf("%d ", i)
			lastline += s
		}
	}
	width := len(lastline)

	for i := width - 1; i >= 0; i-- {
		if lastline[i] == 32 {
			fmt.Printf("%[2]*[1]s\n", lastline[i+1:], width)
		}
	}
	fmt.Println(lastline)

}
