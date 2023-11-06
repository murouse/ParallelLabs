package main

import (
	"gonum.org/v1/plot/plotter"
	"image/color"
	"math"
)

func CirclePlot(radius float64) *plotter.Line {
	xysTop := FuncXYs(-radius, radius, 1000, func(x float64) float64 {
		return CircleTopFunc(x, radius)
	})

	xysBottom := FuncXYs(radius, -radius, 1000, func(x float64) float64 {
		return CircleBottomFunc(x, radius)
	})

	s, _ := plotter.NewLine(append(append(xysTop, xysBottom...), xysTop[0]))
	s.Color = color.RGBA{R: 255, A: 255}
	return s
}

func LemniscatePlot(a float64) *plotter.Line {
	xysTop := FuncXYs(-math.Sqrt(2)*a, math.Sqrt(2)*a, 1000, func(x float64) float64 {
		return LemniscateTopFunc(x, a)
	})

	xysBottom := FuncXYs(math.Sqrt(2)*a, -math.Sqrt(2)*a, 1000, func(x float64) float64 {
		return LemniscateBottomFunc(x, a)
	})

	s, _ := plotter.NewLine(append(append(xysTop, xysBottom...), xysTop[0]))
	s.Color = color.RGBA{B: 255, A: 255}
	return s
}

func CircleTopFunc(x, radius float64) float64 {
	return math.Sqrt(math.Pow(radius, 2) - math.Pow(x, 2))
}

func CircleBottomFunc(x, radius float64) float64 {
	return -CircleTopFunc(x, radius)
}

func LemniscateTopFunc(x, a float64) float64 {
	return math.Sqrt(math.Sqrt(math.Pow(a, 2)*(math.Pow(a, 2)+4*math.Pow(x, 2))) - math.Pow(a, 2) - math.Pow(x, 2))
}

func LemniscateBottomFunc(x, a float64) float64 {
	return -LemniscateTopFunc(x, a)
}

func FuncXYs(from float64, to float64, n int, f func(float64) float64) plotter.XYs {
	x := from
	h := (to - from) / float64(n)

	xys := make(plotter.XYs, 0, n+1)

	for {
		if h > 0 {
			if x >= to {
				break
			}
		} else {
			if x <= to {
				break
			}
		}

		x += h
		y := f(x)
		if math.IsNaN(y) {
			continue
		}
		xys = append(xys, plotter.XY{
			X: x,
			Y: y,
		})
	}

	return xys
}

func PointsPlots(xys Points) (*plotter.Scatter, *plotter.Scatter) {
	var hit, noHit Points
	for i := range xys {
		if xys[i].Hit {
			hit = append(hit, xys[i])
		} else {
			noHit = append(noHit, xys[i])
		}
	}

	hitScatter, _ := plotter.NewScatter(hit)
	hitScatter.Color = color.RGBA{R: 255, A: 255}

	noHitScatter, _ := plotter.NewScatter(noHit)
	noHitScatter.Color = color.RGBA{A: 255}

	return hitScatter, noHitScatter
}
