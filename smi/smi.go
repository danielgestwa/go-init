package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (rect Rectangle) Perimiter() float64 {
	return 2*rect.Height + 2*rect.Width
}

func (rect Rectangle) Area() float64 {
	return rect.Height * rect.Width
}

func (circ Circle) Area() float64 {
	return math.Pi * circ.Radius * circ.Radius
}

func (tri Triangle) Area() float64 {
	return tri.Base / 2 * tri.Height
}
