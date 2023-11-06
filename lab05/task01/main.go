package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"runtime"
	"sync"
	"time"
)

func throwPoints(b *Borders, n int, seed Seed) Points {
	seed.Next()
	x1 := RangeN(b.left1, b.right1, n/2, seed)

	seed.Next()
	y1 := RangeN(b.bottom, b.top, n/2, seed)

	seed.Next()
	x2 := RangeN(b.left2, b.right2, n/2, seed)

	seed.Next()
	y2 := RangeN(b.bottom, b.top, n/2, seed)

	points := make(Points, 0, n)

	for i := range x1 {
		points = append(points, &Point{
			X: x1[i],
			Y: y1[i],
		})
	}

	for i := range x2 {
		points = append(points, &Point{
			X: x2[i],
			Y: y2[i],
		})
	}

	return points
}

func hitRate(a float64, points Points) float64 {
	count := 0
	for i := range points {
		if shadedArea(a, points[i]) {
			count += 1
			points[i].Hit = true
		}
	}
	return float64(count) / float64(len(points))
}

func plotFigure() {
	const n = 1_000 // количество бросаний
	const a = 4     // радиус

	borders := rectBorders(a)
	points := throwPoints(borders, n, Seed{4748, 25647, 325, 6313})
	area := borders.Area()
	result := hitRate(a, points) * area

	fmt.Println(result)

	plots := plot.New()

	plots.Add(CirclePlot(a))
	plots.Add(PointsPlots(points))
	plots.Add(LemniscatePlot(a))

	if err := plots.Save(540, 400, "lab05/task01/plot.svg"); err != nil {
		panic(err)
	}
}

func main() {
	plotFigure()

	const n = 1_048_576 // количество бросаний
	const k = 8         // количество тестов
	const a = 4         // радиус
	const maxThreads = 8

	seeds := []Seed{
		{47478, 235647, 3525, 6313},
		{3637, 352346, 142345, 526},
		{6735, 12314, 48574, 96532},
		{13782, 9654, 213588, 1245},
		{57637, 3563, 133024, 9611},
		{51714, 95432, 2465, 36222},
		{5854, 368, 1365, 61569341},
		{93257, 292141, 8534, 1224},
	}

	p := plot.New()
	values := make(plotter.XYs, 0, maxThreads)

	for threads := 1; threads <= maxThreads; threads++ {
		var generalTime time.Duration
		var generalResult float64
		mu := &sync.Mutex{}

		for i := 1; i <= k; i++ {
			runtime.GC()
			t1 := time.Now()
			wg := &sync.WaitGroup{}
			wg.Add(threads)

			for thread := 0; thread < threads; thread++ {

				go func(thread int) {
					defer wg.Done()

					borders := rectBorders(a)
					points := throwPoints(borders, n/threads, seeds[thread%8])
					area := borders.Area()
					singleResult := hitRate(a, points) * area

					mu.Lock()
					defer mu.Unlock()
					generalResult += singleResult
				}(thread)

			}
			wg.Wait()
			t2 := time.Now()
			generalTime += t2.Sub(t1)
		}

		fmt.Printf("Потоков:\t%v\n", threads)
		fmt.Printf("Время:\t\t%v\n", generalTime/k)
		fmt.Printf("Площадь:\t%v\n\n", generalResult/float64(threads)/k)

		values = append(values, plotter.XY{
			X: float64(threads),
			Y: float64(generalTime / k),
		})
	}

	lines, _ := plotter.NewLine(values)
	p.Add(lines)
	_ = p.Save(1000, 500, "lab05/task01/threads.svg")
}
