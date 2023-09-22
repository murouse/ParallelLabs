package main

import (
	"ParallelLabs/lab03/task01/utils"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"math"
	"math/rand"
	"runtime"
)

//func toValues(x []uint64) plotter.Values {
//	values := make(plotter.Values, len(x))
//
//	for i := range x {
//		values[i] = float64(x[i])
//	}
//
//	return values
//}

func toBars(x []uint64, n uint64) plotter.Values {
	h := uint64(math.MaxUint64) / n
	values := make(plotter.Values, n)

	for i := range x {
		values[x[i]/h] += 1
	}

	return values
}

func histogram(x []uint64) {
	p := plot.New()
	bars, _ := plotter.NewBarChart(toBars(x, 50), 5)
	p.Add(bars)
	err := p.Save(500, 500, fmt.Sprintf("lab03/task02/plots/histogram/plot%d.png", len(x)))
	if err != nil {
		panic(err)
	}
}

func toScatters(x []uint64, n int) plotter.XYs {
	xys := make(plotter.XYs, 0, n)

	for i := range x {
		xys = append(xys, plotter.XY{
			X: float64(x[i]),
			Y: math.Log(float64(n)),
		})
	}

	return xys
}

func scattering(p *plot.Plot, x []uint64, n int) {
	s, _ := plotter.NewScatter(toScatters(x, n))
	p.Add(s)
}

func main() {
	scatterPlot := plot.New()

	for i := 2; i <= 7; i++ {
		runtime.GC()
		n := int(math.Pow(float64(10), float64(i)))
		x := utils.Kiss(n, [4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()})

		histogram(x)

		if i < 5 {
			scattering(scatterPlot, x, n)
		}

	}

	err := scatterPlot.Save(2000, 200, "lab03/task02/plots/scatter/plot.png")
	if err != nil {
		panic(err)
	}
}
