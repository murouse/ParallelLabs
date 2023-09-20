package main

import (
	"ParallelLabs/lab03/task01/utils"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"math"
	"math/rand"
	"runtime"
)

func toValues(x []uint64) plotter.Values {
	values := make(plotter.Values, len(x))

	for i := range x {
		values[i] = float64(x[i])
	}

	return values
}

func main() {
	bar := progressbar.Default(int64(4))

	for i := 2; i <= 5; i++ {
		runtime.GC()
		n := int(math.Pow(float64(10), float64(i)))
		x := utils.Kiss(n, [4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()})

		p := plot.New()
		bars, _ := plotter.NewBarChart(toValues(x), 1)
		p.Add(bars)
		err := p.Save(2000, 500, fmt.Sprintf("lab03/task02/plots/plot%d.png", n))
		if err != nil {
			panic(err)
		}
		_ = bar.Add(1)
	}

	_ = bar.Close()

}
