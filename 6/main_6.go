package main

import (
	"math/rand/v2"
	"sync"
)

type RandIntGenerator struct {
	lowerBound int
	upperBound int
	chanel     chan int
}

func NewRandIntGenerator(lowerBound int, upperBound int) *RandIntGenerator {
	generator := &RandIntGenerator{
		lowerBound: lowerBound,
		upperBound: upperBound,
		chanel:     make(chan int),
	}
	var wg sync.WaitGroup
    wg.Add(1)
    
    go func() {
        wg.Done()
        generator.produce()
    }()
    
    wg.Wait()
	return generator
}

func (r *RandIntGenerator) Get() int {
	return <-r.chanel
}

func (r *RandIntGenerator) produce() {
	for {
		r.chanel <- rand.IntN(r.upperBound-r.lowerBound+1) + r.lowerBound
	}
}

