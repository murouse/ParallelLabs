package main

import (
	"fmt"
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

func ParallelBiasedSampleVariance(elems []float64, n int) {
	groups := ElemsToGroups(elems, n)
	biasedSampleVariances := make([]float64, 0, n)
	var sum float64
	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(n)

	for _, group := range groups {
		group := group
		go func() {
			defer wg.Done()
			bsv := BiasedSampleVariance(group)
			biasedSampleVariances = append(biasedSampleVariances, bsv)
			mu.Lock()
			defer mu.Unlock()
			sum += bsv / float64(len(group))
		}()
	}

	wg.Wait()

	a := SampleAverage(biasedSampleVariances) + BiasedSampleVariance(biasedSampleVariances)*float64(len(elems))/float64(len(elems)+3000)
	fmt.Println("Парал", a)
}

func ElemsToGroups(elems []float64, n int) [][]float64 {
	h := len(elems) / n
	elemGroups := make([][]float64, 0, h)
	for i := 0; i < n-1; i++ {
		elemGroups = append(elemGroups, elems[h*i:h*(i+1)])
	}
	return append(elemGroups, elems[h*(n-1):])
}

func main() {
	//test()
	//n := 10_000_008 //делится на 2, 3, 4

	n := 10_000_000

	elems := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		elems = append(elems, rand.Float64()*100)
	}

	unbiasedSampleVariance := UnbiasedSampleVariance(elems)
	biasedSampleVariance := BiasedSampleVariance(elems)

	fmt.Println(unbiasedSampleVariance)
	fmt.Println(biasedSampleVariance)

	ParallelBiasedSampleVariance(elems, 8)

}

// общая дисперсия равна сумме средней из внутригрупповых и межгрупповой дисперсий
// http://mathprofi.ru/dispersii.html
