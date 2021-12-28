package main

import "math"

// How does something become a shape?
// We just tell Go what a Shape is using an interface declaration

type Shape interface {
	Area() float64
}

/* Rectangle has a method called Area that returns a float64,
so it satisfies the Shape interface */
type Rectangle struct {
	Width  float64
	Height float64
}
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

/* Circle has a method called Area that returns a float64,
so it satisfies the Shape interface */

type Circle struct {
	Radius float64
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base 	float64
	Height 	float64
}
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

