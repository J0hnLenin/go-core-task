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

func (s *ServiceSuite) TestSliceSubtractingEmptyFirstSlice() {
	slice1 := []string{}
	slice2 := []string{"", "123", "abacaba"}
	
	have := SliceSubtracting(slice1, slice2)
	want := []string{}

	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestSliceSubtractingEmptySecondSlice() {
	slice1 := []string{"", "123", "abacaba"}
	slice2 := []string{}
	
	have := SliceSubtracting(slice1, slice2)
	want := []string{"", "123", "abacaba"}

	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestSliceSubtractingEmptyAllSlice() {
	slice1 := []string{}
	slice2 := []string{}
	
	have := SliceSubtracting(slice1, slice2)
	want := []string{}

	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestSliceSubtractingNotEmptySlice() {
	slice1 := []string{"", "333", "123", "123", "123"}
	slice2 := []string{"", "123", "abacaba"}
	
	have := SliceSubtracting(slice1, slice2)
	want := []string{"333"}

	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestSliceSubtractingDublicateSlice() {
	slice1 := []string{"", "333", "123", "123", "123"}
	slice2 := []string{"", "333", "123", "123", "123"}
	
	have := SliceSubtracting(slice1, slice2)
	want := []string{}

	assert.Equal(s.T(), want, have)
}