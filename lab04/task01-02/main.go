package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func Kiss(n int, seed [4]uint64) []uint64 {
	x := make([]uint64, n)
	seed0, seed1, seed2, seed3 := seed[0], seed[1], seed[2], seed[3]
	var t uint64

	for i := 0; i < n; i++ {
		seed0 = 69069*seed0 + 123456
		seed1 = seed1 ^ (seed1 << 13)
		seed1 = seed1 ^ (seed1 >> 17)
		seed1 = seed1 ^ (seed1 << 5)
		t = 698769069*seed2 + seed3
		seed3 = t >> 32
		seed1 = t
		x[i] = seed0 + seed1 + seed2
	}

	return x
}

func KissStat(n int, seed [4]uint64) ([]uint64, float64, float64) {
	x := make([]uint64, n)
	var (
		t   uint64
		sa  float64
		usv float64
	)
	seed0, seed1, seed2, seed3 := seed[0], seed[1], seed[2], seed[3]

	for i := 0; i < n; i++ {
		seed0 = 69069*seed0 + 123456
		seed1 = seed1 ^ (seed1 << 13)
		seed1 = seed1 ^ (seed1 >> 17)
		seed1 = seed1 ^ (seed1 << 5)
		t = 698769069*seed2 + seed3
		seed3 = t >> 32
		seed1 = t
		x[i] = seed0 + seed1 + seed2

		if i != 0 {
			usv += math.Pow(float64(x[i])-sa, 2)/float64(i+1) - usv/float64(i)
		}
		sa += (float64(x[i]) - sa) / float64(i+1)
	}

	return x, sa, usv
}

func KissParallel(n int, threads int, seeds [][4]uint64) []uint64 {
	if threads != len(seeds) {
		panic(fmt.Errorf("need to set a seed for each thread"))
	}

	x := make([]uint64, n)
	wg := new(sync.WaitGroup)
	wg.Add(threads)
	h := n / threads

	for i := 0; i < threads; i++ {
		i := i

		go func() {
			defer wg.Done()
			seed0, seed1, seed2, seed3 := seeds[i][0], seeds[i][1], seeds[i][2], seeds[i][3]
			start := h * i
			var stop int
			if i < threads-1 {
				stop = h * (i + 1)
			} else {
				stop = n
			}
			var t uint64

			for j := start; j < stop; j++ {
				seed0 = 69069*seed0 + 123456
				seed1 = seed1 ^ (seed1 << 13)
				seed1 = seed1 ^ (seed1 >> 17)
				seed1 = seed1 ^ (seed1 << 5)
				t = 698769069*seed2 + seed3
				seed3 = t >> 32
				seed1 = t
				x[j] = seed0 + seed1 + seed2
			}
		}()
	}

	wg.Wait()
	return x
}

func KissStatParallel(n int, threads int, seeds [][4]uint64) ([]uint64, float64, float64) {
	if threads != len(seeds) {
		panic(fmt.Errorf("need to set a seed for each thread"))
	}

	x := make([]uint64, n)
	wg := new(sync.WaitGroup)
	wg.Add(threads)
	h := n / threads

	sas := make([]float64, threads)  // средние групп
	usvs := make([]float64, threads) // дисперсии групп
	ns := make([]int, threads)       // размеры групп

	for i := 0; i < threads; i++ {
		i := i

		go func() {
			defer wg.Done()
			seed0, seed1, seed2, seed3 := seeds[i][0], seeds[i][1], seeds[i][2], seeds[i][3]
			start := h * i
			var stop int
			if i < threads-1 {
				stop = h * (i + 1)
			} else {
				stop = n
			}
			ns[i] = stop - start

			var t uint64

			num := 0

			for j := start; j < stop; j++ {
				seed0 = 69069*seed0 + 123456
				seed1 = seed1 ^ (seed1 << 13)
				seed1 = seed1 ^ (seed1 >> 17)
				seed1 = seed1 ^ (seed1 << 5)
				t = 698769069*seed2 + seed3
				seed3 = t >> 32
				seed1 = t
				x[j] = seed0 + seed1 + seed2

				if num != 0 {
					usvs[i] += math.Pow(float64(x[j])-sas[i], 2)/float64(num+1) - usvs[i]/float64(num)
				}
				sas[i] += (float64(x[j]) - sas[i]) / float64(num+1)

				num += 1
			}
		}()
	}

	wg.Wait()

	var sa float64 // общее среднее
	for i := 0; i < threads; i++ {
		sa += sas[i] * float64(ns[i])
	}
	sa /= float64(n)

	var intra float64 // внутригрупповая дисперсия
	var inter float64 // межгрупповая дисперсия
	for i := 0; i < threads; i++ {
		intra += usvs[i] * float64(ns[i])
		inter += math.Pow(sas[i]-sa, 2) * float64(ns[i])
	}
	usv := (intra + inter) / float64(n) // общая дисперсия

	return x, sa, usv
}

func main() {
	threads := 8

	// autofill
	var seeds [][4]uint64
	for i := 0; i < threads; i++ {
		seeds = append(seeds, [4]uint64{rand.Uint64(), rand.Uint64(), rand.Uint64(), rand.Uint64()})
	}

	//manual fill
	//seeds := [][4]uint64{
	//	{1, 2, 3, 4},
	//	{5, 6, 7, 8},
	//	{9, 10, 11, 12},
	//	{13, 14, 15, 16},
	//	{17, 18, 19, 20},
	//	{21, 22, 23, 24},
	//	{25, 26, 27, 28},
	//	{29, 30, 31, 32},
	//}

	count := 10_000_0000

	var sa, usv float64
	const h = 2
	var sum time.Duration

	fmt.Println("ОДНОПОТОЧНЫЙ")
	sum = time.Duration(0)
	for i := 0; i < h; i++ {
		runtime.GC()
		t1 := time.Now()
		_ = Kiss(count, seeds[0])
		t2 := time.Now()
		sum += t2.Sub(t1)
	}
	fmt.Printf("время: %s\n\n", sum/time.Duration(h))

	fmt.Println("ОДНОПОТОЧНЫЙ СО СТАТИСТИКОЙ")
	sum = time.Duration(0)
	for i := 0; i < h; i++ {
		runtime.GC()
		t1 := time.Now()
		_, sa, usv = KissStat(count, seeds[0])
		t2 := time.Now()
		sum += t2.Sub(t1)
	}
	fmt.Printf("время: %s\n", sum/time.Duration(h))
	fmt.Printf("среднее: %v\n", sa)
	fmt.Printf("дисперсия: %v\n\n", usv)

	fmt.Println("МНОГОПОТОЧНЫЙ")
	sum = time.Duration(0)
	for i := 0; i < h; i++ {
		runtime.GC()
		t1 := time.Now()
		_ = KissParallel(count, threads, seeds)
		t2 := time.Now()
		sum += t2.Sub(t1)
	}
	fmt.Printf("время: %s\n\n", sum/time.Duration(h))

	fmt.Println("МНОГОПОТОЧНЫЙ СО СТАТИСТИКОЙ")
	sum = time.Duration(0)
	for i := 0; i < h; i++ {
		runtime.GC()
		t1 := time.Now()
		_, sa, usv = KissStatParallel(count, threads, seeds)
		t2 := time.Now()
		sum += t2.Sub(t1)
	}
	fmt.Printf("время: %s\n", sum/time.Duration(h))
	fmt.Printf("среднее: %v\n", sa)
	fmt.Printf("дисперсия: %v\n\n", usv)

}
