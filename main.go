package main

import (
	"fmt"
	"math/rand"
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

	var sum int = 0
	var mul uint64 = 1
	for i := 0; i < 10; i++ {
		random := rand.Intn(20) + 1
		if random&1 == 1 {
			sum += random // 任何数与1与，结果是1，是奇数
		} else {
			mul *= uint64(random)
		}
		fmt.Printf("%d  ", random)
	}
	fmt.Printf("\n奇数随机数的累和=%d,偶数随机数的累乘=%d\n", sum, mul)
}

func main() {
	table9()
	fmt.Printf("++++++++++++\n")
	random_addmul()
	fmt.Printf("++++++++++++\n")

}
