package main

func arraySlice() []int {
	a := [8]int{1, 4, 9, 16, 2, 5, 10, 15}
	s := make([]int, 0, len(a)-1)

	for i := 0; i < len(a)-1; i++ {
		s = append(s, a[i]+a[i+1])
	}

	return s
}
