package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	log.Printf("maxProcs: %d", maxProcs)

	numCPU := runtime.NumCPU()
	log.Printf("numCPU: %d", numCPU)

	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

func main() {
	//runtime.GOMAXPROCS(4)
	count := MaxParallelism()
	log.Printf("MaxParallelism: %d", count)

	var wg sync.WaitGroup
	wg.Add(count)

	starts := make(map[int]time.Time, count)
	ends := make(map[int]time.Time, count)

	for i := 0; i < count; i++ {
		i := i
		starts[i] = time.Now()
		go func() {
			defer wg.Done()
			defer func() {
				ends[i] = time.Now()
			}()
			log.Printf("(!) Горутина '%d' запустилась", i+1)
			time.Sleep(time.Duration(i) * time.Second)
			log.Printf("(?) Горутина '%d' завершилась", i+1)
		}()
	}

	wg.Wait()

	for i := 0; i < count; i++ {
		sub := ends[i].Sub(starts[i])
		log.Printf("Горутина '%d' выполнилась за %.8f секунд или %.8f минут", i+1, sub.Seconds(), sub.Minutes())
	}
}
