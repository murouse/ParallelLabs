package main

type Point struct {
	X   float64
	Y   float64
	Hit bool
}

func (ps Points) Len() int {
	return len(ps)
}

func (ps Points) XY(i int) (x, y float64) {
	return ps[i].X, ps[i].Y
}

type Points []*Point
