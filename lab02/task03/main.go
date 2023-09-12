package main

import (
	"ParallelLabs/lab02/task02/utils"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"math"
	"runtime"
	"time"
)

func main() {
	elems := utils.GenerateElems(10_000_000)

	repeatsCount := 5
	threadsCount := 8
	durations := make([][]time.Duration, threadsCount)

	bar := progressbar.Default(int64(threadsCount * repeatsCount))

	for threads := 1; threads <= threadsCount; threads++ {
		var t1, t2 time.Time

		for repeats := 1; repeats <= repeatsCount; repeats++ {
			runtime.GC()
			t1 = time.Now()
			utils.ParallelUnbiasedBiasedSampleVariance(elems, threads)
			t2 = time.Now()
			_ = bar.Add(1)
			durations[threads-1] = append(durations[threads-1], t2.Sub(t1))
		}
	}

	_ = bar.Close()
	p := plot.New()
	values := make(plotter.XYs, 0, threadsCount)

	for i, duration := range durations {
		tmp := make([]float64, 0, len(duration))
		for j := 0; j < len(duration); j++ {
			tmp = append(tmp, float64(duration[j].Milliseconds()))
		}
		a := utils.SampleAverage(tmp)
		v := math.Sqrt(utils.BiasedSampleVariance(tmp))
		values = append(values, plotter.XY{
			X: float64(i),
			Y: a,
		})

		fmt.Printf("%3d %v\t\tμ = %.2f; σ = %.2f\n", i+1, duration, a, v)
	}

	lines, _ := plotter.NewLine(values)
	p.Add(lines)
	_ = p.Save(500, 500, "lab02/task03/plot.png")
}
