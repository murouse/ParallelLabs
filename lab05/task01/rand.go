package main

import "math"

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

func RangeN(min, max float64, n int, seed Seed) []float64 {
	x := UniformN(n, seed)
	for i := range x {
		x[i] = min + x[i]*(max-min)
	}
	return x
}

func (s *Seed) Next() {
	*s = Seed(KissN(4, *s))
}
