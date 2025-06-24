package main

import (
	"math/rand"
	"strings"
	"time"
)

var letters = "abcdefghigklmnopqrstuvwxyz"

func getString(n int) string {
	if n <= 0 {
		return ""
	}
	var building strings.Builder
	for i := 0; i < n; i++ {
		seed := time.Now().UnixNano()
		src := rand.NewSource(seed)
		rand := rand.New(src)
		random := rand.Intn(26)
		building.WriteString(string(letters[random]))
	}
	return building.String()
}

func getReverse(s string) string {

	length := len(s) // len函数针对string类型, 取的是它的字节数
	// s是string类型, 虽然可以通过s[i]对string进行遍历, 但是不能对s[i]做修改
	b := []byte(s)
	var swap int
	if length&1 == 0 {
		swap = length / 2
	} else {
		swap = (length - 1) / 2
	}
	for i := 0; i < swap; i++ {
		b[i], b[length-1-i] = b[length-1-i], b[i]
	}

	return string(b)
}
