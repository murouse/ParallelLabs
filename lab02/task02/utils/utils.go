package utils

import (
	"math"
	"math/rand"
	"sync"
)

// SampleAverage возвращает выборочное среднее
func SampleAverage(elems []float64) float64 { //
	var sum float64
	for _, elem := range elems {
		sum += elem
	}
	return sum / float64(len(elems))
}

// ParallelSampleAverage возвращает выборочное среднее (параллелизм)
func ParallelSampleAverage(elems []float64, n int) float64 {
	groups := ElemsToGroups(elems, n)
	var sum float64
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(n)

	for _, group := range groups {
		group := group
		go func() {
			defer wg.Done()
			var s float64
			for _, elem := range group {
				s += elem
			}
			mu.Lock()
			defer mu.Unlock()
			sum += s
		}()
	}

	wg.Wait()
	return sum / float64(len(elems))
}

// UnbiasedSampleVariance возвращает несмещенную выборочную дисперсию
func UnbiasedSampleVariance(elems []float64) float64 {
	xAvg := SampleAverage(elems)
	var sum float64
	for _, elem := range elems {
		sum += math.Pow(elem-xAvg, 2)
	}
	return sum / float64(len(elems)-1)
}

// BiasedSampleVariance возвращает смещенную выборочную дисперсию
func BiasedSampleVariance(elems []float64) float64 {
	xAvg := SampleAverage(elems)
	var sum float64
	for _, elem := range elems {
		sum += math.Pow(elem-xAvg, 2)
	}
	return sum / float64(len(elems))
}

//func test() {
//	elems := []float64{3, 7, 1, 9, -2, 8, 5}
//	unbiasedSampleVariance := UnbiasedSampleVariance(elems)
//	fmt.Println(unbiasedSampleVariance) //15.95
//}

// ParallelUnbiasedBiasedSampleVariance возвращает несмещенную выборочную дисперсию (параллелизм)
func ParallelUnbiasedBiasedSampleVariance(elems []float64, n int) float64 {
	groups := ElemsToGroups(elems, n)
	a := ParallelSampleAverage(elems, n)

	var sum1 float64
	var sum2 float64
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(n)

	for _, group := range groups {
		group := group
		go func() {
			defer wg.Done()
			tmp1 := BiasedSampleVariance(group) * float64(len(group))
			tmp2 := math.Pow(SampleAverage(group)-a, 2) * float64(len(group))

			mu.Lock()
			defer mu.Unlock()
			sum1 += tmp1
			sum2 += tmp2
		}()
	}

	wg.Wait()

	s1 := sum1 / float64(len(elems)) // внутригрупповая дисперсия (средняя из групповых)
	s2 := sum2 / float64(len(elems)) // межгрупповая дисперсия

	return (s1 + s2) * float64(len(elems)) / (float64(len(elems) - 1))
}

// ElemsToGroups возвращает по возможности равномерные группы элементов
func ElemsToGroups(elems []float64, n int) [][]float64 {
	h := len(elems) / n
	elemGroups := make([][]float64, 0, h)
	for i := 0; i < n-1; i++ {
		elemGroups = append(elemGroups, elems[h*i:h*(i+1)])
	}
	return append(elemGroups, elems[h*(n-1):])
}

func GenerateElems(n int) []float64 {
	elems := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		elems = append(elems, rand.Float64()*100)
	}
	return elems
}
