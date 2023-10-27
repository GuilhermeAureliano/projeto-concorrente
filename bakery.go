package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

type Bakery struct {
	choosing []bool
	ticket   []int
}

var (
	N      int // número de goroutines
	bakery *Bakery
)

func init() {
	flag.IntVar(&N, "N", 100, "Número de goroutines")
	flag.Parse()
}

func max(arr []int) int {
	m := 0
	for _, value := range arr {
		if value > m {
			m = value
		}
	}
	return m
}

func (b *Bakery) lock(id int) {
	if len(b.choosing) != N {
		b.choosing = make([]bool, N)
		b.ticket = make([]int, N)
	}

	b.choosing[id] = true
	b.ticket[id] = max(b.ticket) + 1
	b.choosing[id] = false

	for j := 0; j < N; j++ {
		for b.choosing[j] { // Espera se j está escolhendo
		}
		for b.ticket[j] != 0 && (b.ticket[j] < b.ticket[id] || (b.ticket[j] == b.ticket[id] && j < id)) {
		} // Espera se j tem um ticket menor ou tem o mesmo ticket mas um ID menor
	}
}

func (b *Bakery) unlock(id int) {
	b.ticket[id] = 0
}

func main() {
	bakery := &Bakery{}

	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for k := 0; k < 1000; k++ {
				bakery.lock(id)
				bakery.unlock(id)
			}
		}(i)
	}

	wg.Wait()

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	fmt.Printf("Todas as %d goroutines completaram. Tempo de execução: %s\n", N, elapsedTime)
}
