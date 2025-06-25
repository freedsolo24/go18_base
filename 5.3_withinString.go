package main

import "fmt"

func merge(word1, word2 string) string {
	var shorter int
	var longer string

	if len(word1) < len(word2) {
		shorter, longer = len(word1), word2
	} else {
		shorter, longer = len(word2), word1
	}
	// totallen := len(word1) + len(word2)
	newstring := make([]byte, 2*shorter) // 没有len, 有cap, 一次性构造了底层结构

	for i := 0; i < shorter; i++ {
		newstring[2*i] = word1[i]
		newstring[2*i+1] = word2[i]
	}

	// start := shorter
	// 构建出的newstring后面的一些索引现在是0, 遍历这些索引位置, 用longer字符串的超出部分覆盖这些索引位置的值
	// for j := 2 * shorter; j < totallen; j++ {
	// 	newstring[j] = longer[start]
	// 	start++
	// }
	bytelonger := []byte(longer)
	end := bytelonger[shorter:]

	resultString := fmt.Sprintf("%s%s", string(newstring), string(end))
	return string(resultString)

}
