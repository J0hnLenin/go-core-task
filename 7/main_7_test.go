package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServiceSuite struct {
	suite.Suite
}
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

func (s *ServiceSuite) TestReduceOneElementChannel() {
	
	n := 10
	want := make([]int64, n)
	for i := range(n) {
		want[i] = int64(i)
	}
	
	wg := &sync.WaitGroup{}
    wg.Add(n)
	channels := make([]chan int64, n)
	for i := range(n) {
		channels[i] = make(chan int64)

		go func(ch chan int64, value int64) {
			defer wg.Done()
			ch <- value
		}(channels[i], int64(i))
	}

	go func() {
        wg.Wait()
        for _, ch := range channels {
            close(ch)
        }
    }()

	have := make([]int64, 0, n)
	for i := range Reduce(channels) {
		have = append(have, i)
	}

	assert.ElementsMatch(s.T(), want, have)
}

func (s *ServiceSuite) TestReduceManyElementsChannel() {
	
	n := 10
	want := make([]int64, n*n)
	for i := range(n*n) {
		want[i] = int64(i/n)
	}
	
	wg := &sync.WaitGroup{}
    wg.Add(n)
	channels := make([]chan int64, n)
	for i := range(n) {
		channels[i] = make(chan int64)

		go func(ch chan int64) {
			defer wg.Done()
			for j := range(n) {
				ch <- int64(j)
			}
		}(channels[i])
	}

	go func() {
        wg.Wait()
        for _, ch := range channels {
            close(ch)
        }
    }()

	have := make([]int64, 0, n*n)
	for i := range Reduce(channels) {
		have = append(have, i)
	}

	assert.ElementsMatch(s.T(), want, have)
}

func (s *ServiceSuite) TestReduceOneChannel() {
	
	n := 1
	want := make([]int64, n*n)
	for i := range(n*n) {
		want[i] = int64(i/n)
	}
	
	wg := &sync.WaitGroup{}
    wg.Add(n)
	channels := make([]chan int64, n)
	for i := range(n) {
		channels[i] = make(chan int64)

		go func(ch chan int64) {
			defer wg.Done()
			for j := range(n) {
				ch <- int64(j)
			}
		}(channels[i])
	}

	go func() {
        wg.Wait()
        for _, ch := range channels {
            close(ch)
        }
    }()

	have := make([]int64, 0, n*n)
	for i := range Reduce(channels) {
		have = append(have, i)
	}

	assert.ElementsMatch(s.T(), want, have)
}