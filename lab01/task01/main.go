package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"math/rand"
)

type Point struct {
	x float64
	y float64
}

func (ps Points) Len() int {
	return len(ps)
}

func (ps Points) XY(i int) (x, y float64) {
	return ps[i].x, ps[i].y
}

type Points []Point

type AreaFunc func(*Point) bool

func MarkPoints(points Points, areaFunc AreaFunc) (inside Points, outside Points) {
	for _, point := range points {
		if areaFunc(&point) {
			inside = append(inside, point)
			continue
		}
		outside = append(outside, point)
	}
	return
}

func SavePlot(inside Points, outside Points, file string) {
	p := plot.New()
	p.Title.Text = "Area"
	_ = plotutil.AddScatters(p, "Inside", inside, "Outside", outside)
	_ = p.Save(500, 500, file)
}

func GeneratePoints(num int) (points Points) {
	for i := 0; i < num; i++ {
		xr := rand.Float64()
		yr := rand.Float64()
		points = append(points, Point{
			x: -10 + 20*xr,
			y: -10 + 20*yr,
		})
	}
	return
}

func main() {
	//points := Points{
	//	{1, 2},
	//	{2, 4},
	//	{5, 7},
	//}

	//areaFunc := func(point *Point) bool {
	//	return point.y > point.x*point.x
	//}

	areaFunc := func(point *Point) bool {
		const a = 4
		return point.y+point.x < 5/2.0*a && point.x*point.y > a*a
	}

	points := GeneratePoints(10_000)
	inside, outside := MarkPoints(points, areaFunc)
	SavePlot(inside, outside, "lab01/task01/plot1.png")

	//points = utils.RemoveIf[Point](points, func(point *Point) bool {
	//	return !areaFunc(point)
	//})
	//SavePlot(points, Points{}, "lab01/task01/plot2.png")
}
