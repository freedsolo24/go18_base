package main

import "fmt"

func getSliceDetails(s []int) string {
	if len(s) > 0 {
		str1 := fmt.Sprintf("内容:%v", s)
		str2 := fmt.Sprintf("标头值指针:%p", &s)
		str3 := fmt.Sprintf("第一个元素的指针:%p", &s[0])
		str4 := fmt.Sprintf("长度:%d", len(s))
		str5 := fmt.Sprintf("容量:%d", cap(s))
		str := fmt.Sprintf("%-16s|%s|%s|%s|%s\n", str1, str2, str3, str4, str5)
		return str
	} else {
		return "切片中没有元素"
	}
}
