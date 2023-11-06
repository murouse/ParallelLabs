package main

import "math"

type Borders struct {
	left1  float64
	right1 float64
	left2  float64
	right2 float64
	bottom float64
	top    float64
}

func rectBorders(a float64) *Borders {
	return &Borders{
		left1:  -math.Sqrt(2) * a,
		right1: -math.Cos(math.Pi/6) * a,
		left2:  math.Cos(math.Pi/6) * a,
		right2: math.Sqrt(2) * a,
		bottom: -math.Sin(math.Pi/6) * a,
		top:    math.Sin(math.Pi/6) * a,
	}
}

func (b *Borders) Area() float64 {
	return (b.right1 - b.left1) * (b.top - b.bottom) * 2
}
