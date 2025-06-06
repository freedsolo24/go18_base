package main

import "math"

// 定义对象
type shape struct {
	area float32
}
type circle struct {
	radius float32
	shape
}
type triangle struct {
	base   float32
	height float32
	shape
}
type rectangle struct {
	length float32
	width  float32
	shape
}

// 为对象创建构造函数
func newcircle(r float32) *circle {
	return &circle{radius: r}
}
func newtriangle(b, h float32) *triangle {
	return &triangle{base: b, height: h}
}
func newrectangle(l, w float32) *rectangle {
	return &rectangle{length: l, width: w}
}

// 为不同的对象创建同一个方法：求面积
func (i *circle) calcarea() float32 {
	// return math.Pi * i.radius * i.radius
	if i.area == 0 {
		i.shape.area = math.Pi * i.radius * i.radius
	}
	return i.shape.area
}
func (i *triangle) calcarea() float32 {
	// return 0.5 * i.base * i.height
	if i.area == 0 {
		i.shape.area = 0.5 * i.base * i.height
	}
	return i.shape.area
}
func (i *rectangle) calcarea() float32 {
	// return i.length * i.width
	if i.area == 0 {
		i.shape.area = i.length * i.width
	}
	return i.shape.area
}

// 定义接口要求实现方法：求面积
type interfacers interface {
	calcarea() float32
}
