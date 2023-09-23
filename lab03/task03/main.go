package main

import (
	utils2 "ParallelLabs/lab02/task02/utils"
	"ParallelLabs/lab03/task01/utils"
	"fmt"
	"math"
	"math/rand"
)

func main() {
	for _, n := range []int{10, 100, 1_000, 10_000, 100_000, 1_000_000} {
		x := utils.Kiss(n, [4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()})

		xf := make([]float64, len(x))

		for i := range x {
			xf[i] = float64(x[i])
		}

		sa := utils2.SampleAverage(xf)
		usv := utils2.UnbiasedSampleVariance(xf)

		fmt.Printf("Количество элементов:\t%d\n", n)
		fmt.Printf("Ожидаемое среднее:\t%d\n", math.MaxUint64/2)
		fmt.Printf("Ожидаемая дисперсия:\t%.0f\n", math.Pow(float64(math.MaxUint64), 2)/12)
		fmt.Printf("Полученное среднее:\t%.0f\n", sa)
		fmt.Printf("Полученная дисперсия:\t%.0f\n\n", usv)
	}
}
