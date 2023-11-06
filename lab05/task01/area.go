package main

import "math"

func lemniscateArea(a float64, p *Point) bool {
	x2 := math.Pow(p.X, 2)
	y2 := math.Pow(p.Y, 2)
	a2 := math.Pow(a, 2)
	return math.Pow(x2+y2, 2) < 2*a2*(x2-y2)
}

func circleArea(a float64, p *Point) bool {
	return p.X*p.X+p.Y*p.Y < a*a
}

func shadedArea(a float64, p *Point) bool {
	return lemniscateArea(a, p) && !circleArea(a, p)
}
