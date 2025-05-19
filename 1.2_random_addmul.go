package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random_addmul() {

	seed := time.Now().UnixNano() // 随机数种子
	src := rand.NewSource(seed)   // 随机数发生源
	r := rand.New(src)            // 随机数生成器

	var sum int = 0
	var mul uint64 = 1
	for i := 0; i < 10; i++ {
		random := r.Intn(20) + 1
		if random&1 == 1 {
			sum += random // 任何数与1与，结果是1，是奇数
		} else {
			mul *= uint64(random)
		}
		fmt.Printf("%d  ", random)
	}
	fmt.Printf("\n奇数随机数的累和=%d,偶数随机数的累乘=%d\n", sum, mul)
}
