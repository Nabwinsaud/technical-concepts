package main

type Circle struct {
	radius float64
}

const PI = 3.14

func (c Circle) Area() float64 {
	return PI * c.radius * c.radius
}

func (c *Circle) CalculateArea() float64 {
	return PI * c.radius * c.radius
}
