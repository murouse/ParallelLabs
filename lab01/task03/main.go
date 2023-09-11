package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"math"
	"math/rand"
)

func Bernoulli(p float64) int {
	if rand.Float64() < p {
		return 1
	}
	return 0
}

func Binomial(p float64, n int) (res int) {
	for i := 0; i < n; i++ {
		res += Bernoulli(p)
	}
	return
}

func NormalPDF(x float64, mu float64, sigma float64) float64 {
	return 1 / (sigma * math.Sqrt(2*math.Pi)) * math.Exp(-math.Pow(x-mu, 2)/(2*math.Pow(sigma, 2)))
}

func main() {
	n := 25
	h := 10000000
	prob := 0.35

	valuesBinomial := make(plotter.Values, n+1)
	for i := 0; i < h; i++ {
		valuesBinomial[Binomial(prob, n)] += 1 / float64(h)
	}

	p := plot.New()
	bars, _ := plotter.NewBarChart(valuesBinomial, 10)
	p.Add(bars)

	var values2 plotter.XYs
	for i := float64(0); i <= float64(n); i += 0.01 {
		values2 = append(values2, plotter.XY{
			X: i,
			Y: NormalPDF(i, float64(n)*prob, math.Sqrt(float64(n)*prob*(1-prob))), //σ = √ np (1-p)
		})
	}

	lines, _ := plotter.NewLine(values2)
	p.Add(lines)

	_ = p.Save(500, 500, "lab01/task03/plot.png")
}
