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

func plotHistogram(x []uint64, n uint64) {
	h := uint64(math.MaxUint64) / n
	values := make(plotter.Values, n)

	for i := range x {
		values[x[i]/h] += 1
	}

	p := plot.New()
	bars, _ := plotter.NewBarChart(values, 5)
	p.Add(bars)
	err := p.Save(500, 500, fmt.Sprintf("lab03/task02/plots/histogram/plot%d.png", len(x)))
	if err != nil {
		panic(err)
	}
}

func addScatters(p *plot.Plot, x []uint64, n int) {
	xys := make(plotter.XYs, 0, n)

	for i := range x {
		xys = append(xys, plotter.XY{
			X: float64(x[i]),
			Y: math.Log(float64(n)),
		})
	}

	s, _ := plotter.NewScatter(xys)
	p.Add(s)
}

func corrCoef(X []float64, Y []float64) float64 {
	if len(X) != len(Y) {
		panic(fmt.Errorf("error"))
	}

	var sumX, sumY, sumXY, squareSumX, squareSumY float64
	n := float64(len(X))

	for i := 0; i < len(X); i++ {
		sumX += X[i]
		sumY += Y[i]
		sumXY += X[i] * Y[i]
		squareSumX += X[i] * X[i]
		squareSumY += Y[i] * Y[i]
	}
	return (n*sumXY - sumX*sumY) / (math.Sqrt((n*squareSumX - sumX*sumX) * (n*squareSumY - sumY*sumY)))
}

func plotAcf(x []uint64, lags int) {
	xys := make(plotter.XYs, 0, len(x))

	xf := make([]float64, len(x))

	for i := range x {
		xf[i] = float64(x[i])
	}

	for lag := 0; lag <= lags; lag++ {
		xys = append(xys, plotter.XY{
			X: float64(lag),
			Y: corrCoef(xf[:len(x)-lag], xf[lag:]),
		})
	}

	p := plot.New()
	line, sc, _ := plotter.NewLinePoints(xys)
	p.Add(line, sc)
	ax, _ := plotter.NewLine(plotter.XYs{plotter.XY{}, plotter.XY{X: float64(lags)}})
	p.Add(ax)

	err := p.Save(1000, 500, fmt.Sprintf("lab03/task02/plots/acf/plot%d.png", lags))
	if err != nil {
		panic(err)
	}
}

func main() {
	scatterPlot := plot.New()

	for i := 2; i <= 7; i++ {
		runtime.GC()
		n := int(math.Pow(float64(10), float64(i)))
		x := utils.Kiss(n, [4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()})

		plotHistogram(x, 50)

		if i < 5 {
			addScatters(scatterPlot, x, n)
		}

		if i == 5 {
			plotAcf(x, 50)
		}

	}

	err := scatterPlot.Save(2000, 200, "lab03/task02/plots/scatter/plot.png")
	if err != nil {
		panic(err)
	}
}
