package main

import (
	"fmt"
	"math"
)

type shape interface {
	perimeter() float64
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length  float64
	breadth float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	length float64
}

func (s square) area() float64 {
	if s.length < 0 {
		return -1
	}
	return s.length * s.length
}

func (s square) perimeter() float64 {
	if s.length < 0 {
		return -1
	}
	return 4 * s.length
}

func (r rectangle) area() float64 {
	if r.length < 0 || r.breadth < 0 {
		return -1
	}
	return r.length * r.breadth
}

func (r rectangle) perimeter() float64 {
	if r.length < 0 || r.breadth < 0 {
		return -1
	}
	return 2 * (r.length + r.breadth)
}

func (c circle) perimeter() float64 {
	if c.radius < 0 {
		return -1
	}
	return 2 * pi * c.radius
}

func (c circle) area() float64 {
	if c.radius < 0 {
		return -1
	}
	return pi * c.radius * c.radius
}

func (t triangle) perimeter() float64 {
	if t.base < 0 || t.height < 0 {
		return -1
	}
	return t.base + t.height + math.Sqrt(t.base*t.base+t.height*t.height)
}

func (t triangle) area() float64 {
	if t.base < 0 || t.height < 0 {
		return -1
	}
	return 0.5 * t.base * t.height
}

func getDetails(s shape) string {
	return fmt.Sprintf("Shape Details => Area: %.1f Perimeter: %.1f", s.area(), s.perimeter())
}
