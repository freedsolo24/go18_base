package main

import "fmt"

func table9() {
	var width int
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
func main() {
	table9()
}
