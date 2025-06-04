package main

import (
	"math"
)

// 定义接口类型,接口要实现area动作,不传参,返回值是float32
// 对象生成的实例既是自己结构类的实例,又是接口的实例
type Interfacer interface {
	area() float32
}

// struct定义三个对象
type Circle struct {
	radius float32
}
type Triangle struct {
	base   float32
	height float32
}
type Rectangle struct {
	length float32
	width  float32
}

// 给对象定义构造函数
func NewCircle(r float32) *Circle {
	return &Circle{radius: r}
}
func NewTriangle(b, h float32) *Triangle {
	return &Triangle{base: b, height: h}
}
func NewRectangle(l, w float32) *Rectangle {
	return &Rectangle{length: l, width: w}
}

// 定义func(receiver)接收不同的实例,调用共同的方法
func (c *Circle) area() float32 {
	return c.radius * c.radius * math.Pi
}
func (t *Triangle) area() float32 {
	return 0.5 * t.base * t.height
}
func (r *Rectangle) area() float32 {
	return r.length * r.width
}
