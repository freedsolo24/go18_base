package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func numsRepeat() {

	seed := time.Now().UnixNano()
	randsrc := rand.NewSource(seed)
	rander := rand.New(randsrc)

	s := make([]int, 0, 10)
	m := make(map[int]int, 10)

	for i := 0; i < 10; i++ {
		r := rander.Intn(10) + 1
		s = append(s, r)

		if _, ok := m[r]; !ok {
			m[r] = 1
		} else {
			m[r] += 1
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(s)))

	fmt.Printf("生成的随机数降序:%v\n", s)

	for key, value := range m {
		fmt.Printf("随机数:%d,重复的次数:%d   ", key, value)
	}
	fmt.Println()

}
