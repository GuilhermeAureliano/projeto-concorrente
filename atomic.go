package main

import (
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	N           int // número de goroutines
	choosing    []int32
	ticket      []int32
)

func init() {
	flag.IntVar(&N, "N", 100, "Número de goroutines")
	flag.Parse()
	choosing = make([]int32, N)
	ticket = make([]int32, N)
}

func max(arr []int32) int32 {
	m := int32(0)
	for _, value := range arr {
		if value > m {
			m = value
		}
	}
	return m
}

func lock(id int) {
	atomic.StoreInt32(&choosing[id], 1)
	atomic.StoreInt32(&ticket[id], max(ticket)+1)
	atomic.StoreInt32(&choosing[id], 0)

	for j := 0; j < N; j++ {
		for atomic.LoadInt32(&choosing[j]) == 1 { // Espera se j está escolhendo
		}
		for atomic.LoadInt32(&ticket[j]) != 0 && 
		    (atomic.LoadInt32(&ticket[j]) < atomic.LoadInt32(&ticket[id]) || 
		    (atomic.LoadInt32(&ticket[j]) == atomic.LoadInt32(&ticket[id]) && j < id)) {
		} // Espera se j tem um ticket menor ou tem o mesmo ticket mas um ID menor
	}
}

func unlock(id int) {
	atomic.StoreInt32(&ticket[id], 0)
}

func main() {
	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for k := 0; k < 1000; k++ {
				lock(id)
				unlock(id)
			}
		}(i)
	}

	wg.Wait()

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	fmt.Printf("Todas as %d goroutines completaramm. Tempo de execução: %s\n", N, elapsedTime)
}