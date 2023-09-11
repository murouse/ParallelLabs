package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat/distuv"
	"math"
	"strings"
)

//func Factorial(n int) int {
//	res := 1
//	for i := 2; i <= n; i++ {
//		res *= i
//	}
//	return res
//}
//
//func Poisson(lambda float64, k int) float64 {
//	return math.Exp(-lambda) * math.Pow(lambda, float64(k)) / float64(Factorial(k))
//}

func main() {
	h := 100000
	m := make(map[int]int)
	p := distuv.Poisson{Lambda: 10}

	for i := 0; i < h; i++ {
		m[int(math.Round(p.Rand()))] += 1
	}

	maxKey := 0
	for key := range m {
		if key > maxKey {
			maxKey = key
		}
	}

	for i := 0; i <= maxKey; i++ {
		fmt.Printf("%s\n", strings.Repeat("*", m[i]/(h/200)))
	}
}
