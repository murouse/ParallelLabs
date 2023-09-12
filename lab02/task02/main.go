package main

import (
	"ParallelLabs/lab02/task02/utils"
	"fmt"
	"runtime"
	"time"
)

func main() {
	//test()
	elems := utils.GenerateElems(100_000_000)

	t1 := time.Now()
	unbiasedSampleVariance := utils.UnbiasedSampleVariance(elems)
	t2 := time.Now()
	fmt.Println("Без параллелизма: ", t2.Sub(t1), unbiasedSampleVariance)

	for i := 1; i <= 8; i++ {
		runtime.GC()
		t1 = time.Now()
		parallelUnbiasedSampleVariance := utils.ParallelUnbiasedBiasedSampleVariance(elems, i)
		t2 = time.Now()
		fmt.Printf("%d потоков: %v %v\n", i, t2.Sub(t1), parallelUnbiasedSampleVariance)
	}
}

// общая дисперсия равна сумме средней из внутригрупповых и межгрупповой дисперсий
// http://mathprofi.ru/dispersii.html
