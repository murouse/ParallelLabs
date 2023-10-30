package main

import (
	"ParallelLabs/lab02/task02/utils"
	"fmt"
	"github.com/thoas/go-funk"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"math"
)

type Seed [4]uint64

func KissN(n int, seed Seed) []uint64 {
	x := make([]uint64, n)

	for i := 0; i < n; i++ {
		seed[0] = 69069*seed[0] + 123456
		seed[1] = seed[1] ^ (seed[1] << 13)
		seed[1] = seed[1] ^ (seed[1] >> 17)
		seed[1] = seed[1] ^ (seed[1] << 5)
		t := 698769069*seed[2] + seed[3]
		seed[3] = t >> 32
		seed[2] = t
		x[i] = seed[0] + seed[1] + seed[2]
	}

	return x
}

func UniformN(n int, seed Seed) []float64 {
	x := KissN(n, seed)
	y := make([]float64, n)
	for i := range x {
		y[i] = float64(x[i]) / math.MaxUint64
	}
	return y
}

func Normal(seed Seed) (float64, float64) {
	var s, u1, u2 float64
	newSeed := [4]uint64(KissN(4, seed))
	for {
		newSeed = [4]uint64(KissN(4, newSeed))
		u := UniformN(2, newSeed)
		u1 = u[0]
		u2 = u[1]

		if newSeed[0]%2 == 0 {
			u1 = -u1
		}

		if newSeed[1]%2 == 0 {
			u2 = -u2
		}

		s = u1*u1 + u2*u2
		if s > 0 && s <= 1 {
			break
		}
	}
	x := u1 * math.Pow(-2*math.Log(s)/s, 0.5)
	y := u2 * math.Pow(-2*math.Log(s)/s, 0.5)
	return x, y
}

func NormalN(n int, seed Seed) []float64 {
	x := make([]float64, n+1)
	s := KissN((n+1)*4, seed)
	for i := 0; i < n; i += 2 {
		x[i], x[i+1] = Normal(Seed(s[i*2 : i*2+4]))
	}
	return x[:n]
}

func ExponentialN(n int, lambda float64, seed Seed) []float64 {
	x := UniformN(n, seed)
	for i := range x {
		x[i] = -math.Log(1-x[i]) / lambda
	}
	return x
}

func Poisson(lambda float64, seed Seed) uint64 {
	a := math.Exp(-lambda)
	k := uint64(0)
	p := float64(1)

	for {
		k += 1
		seed = Seed(KissN(4, seed))
		u := UniformN(1, seed)[0]
		p *= u

		if p <= a {
			return k - 1
		}
	}
}

func PoissonN(n int, lambda float64, seed Seed) []uint64 {
	x := make([]uint64, n)
	for i := range x {
		seed = Seed(KissN(4, seed))
		x[i] = Poisson(lambda, seed)
	}
	return x
}

// WeibullN
// распределение https://help.fsight.ru/ru/mergedProjects/lib/05_statistics/distribution/lib_weibulldistribution.htm
// калькулятор https://homepage.divms.uiowa.edu/~mbognar/applets/weibull.html
func WeibullN(n int, lambda, k float64, seed Seed) []float64 {
	x := UniformN(n, seed)
	for i := range x {
		x[i] = lambda * math.Pow(-math.Log(x[i]), 1/k)
	}
	return x
}

func PlotHistogram(x []float64, name string, n int) {
	p := plot.New()
	hist, _ := plotter.NewHist(plotter.Values(x), n)
	p.Add(hist)
	err := p.Save(1000, 500, fmt.Sprintf("lab04/task03/plots/%s.png", name))
	if err != nil {
		panic(err)
	}
}

func main() {
	const n = 10_000_000
	const h = 200
	const threads = 8
	seed := Seed{5, 9, 267, 26342}

	normal := NormalN(n, seed)
	PlotHistogram(normal, "normal", h)
	sa := utils.ParallelSampleAverage(normal, threads)
	ubs := utils.ParallelUnbiasedBiasedSampleVariance(normal, threads)
	fmt.Println("Нормальное распределение")
	fmt.Printf("μ = %.2f\n", sa)
	fmt.Printf("σ = %.2f\n\n", ubs)

	lambda := 0.5
	exponential := ExponentialN(n, lambda, seed)
	PlotHistogram(exponential, "exponential", h)
	sa = utils.ParallelSampleAverage(exponential, threads)
	ubs = utils.ParallelUnbiasedBiasedSampleVariance(exponential, threads)
	fmt.Println("Экспоненциальное распределение")
	fmt.Printf("λ = %.2f\n", lambda)
	fmt.Printf("μ = 1/λ = %.2f\n", sa)
	fmt.Printf("σ = 1/λ = %.2f\n", math.Sqrt(ubs))
	fmt.Printf("σ^2 = (1/λ)^2 = %.2f\n\n", ubs)

	lambda = 10
	poisson := PoissonN(n, lambda, seed)
	poissonF := funk.Map(poisson, func(a uint64) float64 {
		return float64(a)
	}).([]float64)
	PlotHistogram(poissonF, "poisson", h)
	sa = utils.ParallelSampleAverage(poissonF, threads)
	ubs = utils.ParallelUnbiasedBiasedSampleVariance(poissonF, threads)
	fmt.Println("Распределение Пуассона")
	fmt.Printf("λ = %.2f\n", lambda)
	fmt.Printf("μ = λ = %.2f\n", sa)
	fmt.Printf("σ = √λ = %.2f\n", math.Sqrt(ubs))
	fmt.Printf("σ^2 = λ = %.2f\n\n", ubs)

	lambda = 1
	k := 1.5
	weibull := WeibullN(n, lambda, k, seed)
	PlotHistogram(weibull, "weibull", h)
	sa = utils.ParallelSampleAverage(weibull, threads)
	ubs = utils.ParallelUnbiasedBiasedSampleVariance(weibull, threads)
	fmt.Println("Распределение Вейбулла")
	fmt.Printf("λ = %.2f\tk = %.2f\n", lambda, k)
	fmt.Printf("μ = %.2f\n", sa)
	fmt.Printf("σ = %.2f\n", math.Sqrt(ubs))
	fmt.Printf("σ^2 = %.2f\n\n", ubs)
}
