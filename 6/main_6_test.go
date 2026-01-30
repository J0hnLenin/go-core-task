package main

import (
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

func (s *ServiceSuite) TestDifferentGenerator() {
	// Проверяем что два разных генератора выдают разные результаты

	n := 10
	generator1 := NewRandIntGenerator(-10000, 10000)
	generator2 := NewRandIntGenerator(-10000, 10000)

	slice1 := make([]int, n)
	slice2 := make([]int, n)
	
	for i := range(n) {
		slice1[i] = generator1.Get()
		slice2[i] = generator2.Get()
	}

	assert.NotEqual(s.T(), slice1, slice2)
}

func (s *ServiceSuite) TestGeneratorRange() {
	// Проверяем что генератор выдаёт все числа из диапазона [L, R]

	n := 1000
	lowerBound := 0
	upperBound := 10
	generator := NewRandIntGenerator(lowerBound, upperBound)
	counter := make([]int, upperBound+1)
	
	for range(n) {
		index := generator.Get()

		indexInBound := lowerBound <= index && index <= upperBound 
		assert.True(s.T(), indexInBound)
		if indexInBound {
			counter[index] += 1
		}
	}

	assert.NotContains(s.T(), counter, 0)
}

