package main

import (
	"ParallelLabs/lab02/task02/utils"
	"fmt"
	"math"
	"math/rand"
)

func KissPlus(n int, seed [4]uint64) ([]uint64, float64, float64) {
	x := make([]uint64, n)
	var (
		t   uint64
		sa  float64
		usv float64
	)

	for i := 0; i < n; i++ {
		seed[0] = 69069*seed[0] + 123456
		seed[1] = seed[1] ^ (seed[1] << 13)
		seed[1] = seed[1] ^ (seed[1] >> 17)
		seed[1] = seed[1] ^ (seed[1] << 5)
		t = 698769069*seed[2] + seed[3]
		seed[3] = t >> 32
		seed[1] = t
		x[i] = seed[0] + seed[1] + seed[2]

		if i != 0 {
			usv += math.Pow(float64(x[i])-sa, 2)/float64(i+1) - usv/float64(i)
		}
		sa += (float64(x[i]) - sa) / float64(i+1)
	}

	return x, sa, usv
}

func main() {
	for _, n := range []int{10_000_000, 100_000_000} {
		x, saf, usvf := KissPlus(n, [4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()})

		xf := make([]float64, len(x))

		for i := range x {
			xf[i] = float64(x[i])
		}

		sa := utils.SampleAverage(xf)
		usv := utils.UnbiasedSampleVariance(xf)

		fmt.Printf("Количество элементов:\t%d\n", n)
		fmt.Printf("Ожидаемое среднее:\t%d\n", math.MaxUint64/2)
		fmt.Printf("Ожидаемая дисперсия:\t%.0f\n", math.Pow(float64(math.MaxUint64), 2)/12)
		fmt.Printf("Полученное среднее:\t%.0f\n", sa)
		fmt.Printf("Полученная дисперсия:\t%.0f\n", usv)
		fmt.Printf("Уэлфорд среднее:\t%.0f\n", saf)
		fmt.Printf("Уэлфорд дисперсия:\t%.0f\n\n", usvf)
	}
}

//func test() {
//	x := []float64{1, 5, 8, 3, 4, 1, 7, 8, 3}
//
//	usv := float64(0)
//	sa := x[0]
//
//	for i := 1; i < len(x); i++ {
//		usv += math.Pow(x[i]-sa, 2)/float64(i+1) - usv/float64(i)
//		sa += (x[i] - sa) / float64(i+1)
//	}
//
//	fmt.Println(utils.SampleAverage(x))
//	fmt.Println(sa)
//
//	fmt.Println(utils.UnbiasedSampleVariance(x))
//	fmt.Println(usv)
//}
