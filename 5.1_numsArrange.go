package main

import (
	"fmt"
)

func nums1() {
	count := 0
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			for k := 1; k <= 4; k++ {
				if i != j && i != k && j != k {
					fmt.Println(i, j, k)
					count++
				}

			}
		}
	}

	fmt.Println(count)
}

func nums2() {
	chars := []int{1, 2, 3, 4}
	length := len(chars)
	bitmap := make([]byte, length)
	count := 0 // 记录总数

	for i := 0; i < 4; i++ {
		x := chars[i] // 先拿第i个元素
		bitmap[i] = 1 // 位图第i位=1，表示已用
		for j := 0; j < 4; j++ {
			if bitmap[j] == 1 {
				continue // 如果这个位图=1，不要，继续循环
			}
			bitmap[j] = 1 // 如果这个位图=0，我就要，并且我设置1
			y := chars[j]
			for k := 0; k < 4; k++ {
				if bitmap[k] == 1 {
					continue
				}
				// bitmap[k] = 1
				z := chars[k]
				fmt.Printf("%d%d%d\n", x, y, z)
				// bitmap[k] = 0 // 打印完了k之后，bitmap=0
				count++
			}
			bitmap[j] = 0 // 迭代下一次的时候，这次不用了，还原零
		}
		bitmap[i] = 0 // 迭代下一次的时候，这次不用了，还原零
	}
	fmt.Println(count)
}
