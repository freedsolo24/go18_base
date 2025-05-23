package main

func PUT(x, y int, opfunc func(int, int) int) int {

	return opfunc(x, y)
}
