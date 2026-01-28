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

func (s *ServiceSuite) TestIntersectionEmptyFirstSlice() {
	slice1 := []int{}
	slice2 := []int{1, 2, 3}
	
	notEmpty, have := Intersection(slice1, slice2)
	
	want := []int{}

	assert.Equal(s.T(), want, have)
	assert.False(s.T(), notEmpty)
}

func (s *ServiceSuite) TestIntersectionEmptySecondSlice() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{}
	
	notEmpty, have := Intersection(slice1, slice2)
	
	want := []int{}

	assert.Equal(s.T(), want, have)
	assert.False(s.T(), notEmpty)
}

func (s *ServiceSuite) TestIntersectionTwoSlice() {
	slice1 := []int{}
	slice2 := []int{}
	
	notEmpty, have := Intersection(slice1, slice2)
	
	want := []int{}

	assert.Equal(s.T(), want, have)
	assert.False(s.T(), notEmpty)
}

func (s *ServiceSuite) TestIntersectionEmpty() {
	slice1 := []int{1, 3, 5}
	slice2 := []int{2, 4, 6}
	
	notEmpty, have := Intersection(slice1, slice2)
	
	want := []int{}

	assert.Equal(s.T(), want, have)
	assert.False(s.T(), notEmpty)
}

func (s *ServiceSuite) TestIntersectionDublicateSlice() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{1, 2, 3, 4, 5}
	
	notEmpty, have := Intersection(slice1, slice2)
	
	want := []int{1, 2, 3, 4, 5}

	assert.ElementsMatch(s.T(), want, have)
	assert.True(s.T(), notEmpty)
}

func (s *ServiceSuite) TestIntersection() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{3, 4, 5}
	
	notEmpty, have := Intersection(slice1, slice2)
	
	want := []int{3}

	assert.ElementsMatch(s.T(), want, have)
	assert.True(s.T(), notEmpty)
}

func (s *ServiceSuite) TestIntersectionWithDublicates() {
	slice1 := []int{1, 2, 3, 3, 3, 1, 2}
	slice2 := []int{4, 5, 5, 4, 4, 4, 3, 3}
	
	notEmpty, have := Intersection(slice1, slice2)
	
	want := []int{3}

	assert.ElementsMatch(s.T(), want, have)
	assert.True(s.T(), notEmpty)
}