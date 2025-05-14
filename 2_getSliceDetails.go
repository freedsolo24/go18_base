package main

import "fmt"

func getSliceDetails(s []int) string {
	if len(s) > 0 {
		str := fmt.Sprintf("内容:%v,标头值指针:%p,第一个元素的指针:%p, 长度:%d,容量:%d", s, &s, &s[0], len(s), cap(s))
		return str
	} else {
		return "切片中没有元素"
	}
}
