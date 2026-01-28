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

func (s *ServiceSuite) TestSliceExampleEmptySlice() {
	slice := []int{}
	have := sliceExample(slice)
	want := []int{}
	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestSliceExampleOneElementEvenSlice() {
	slice := []int{1}
	have := sliceExample(slice)
	want := []int{}
	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestSliceExampleOneElementNotEvenSlice() {
	slice := []int{0}
	have := sliceExample(slice)
	want := []int{0}
	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestSliceExampleEvenSlice() {
	slice := []int{2, 4, 8, 8}
	have := sliceExample(slice)
	want := []int{2, 4, 8, 8}
	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestSliceExampleNotEvenSlice() {
	slice := []int{1, 1, 9, 11}
	have := sliceExample(slice)
	want := []int{}
	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestSliceExampleMixedSlice() {
	slice := []int{1, 2, 3, 4}
	have := sliceExample(slice)
	want := []int{2, 4}
	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestSliceExampleOriginalSliceNotChanged() {
	have := []int{1, 2, 3, 4}
	want := []int{1, 2, 3, 4}
	sliceExample(have)
	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestAddElementsAppend() {
	slice := []int{1, 2, 3, 4}
	addition := 5
	have := addElements(slice, addition)
	want := append(slice, addition)
	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestAddElementsEmptySlice() {
	slice := []int{}
	addition := 1
	have := addElements(slice, addition)
	want := append(slice, addition)
	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestAddElementsChangeSlice() {
	slice := []int{1, 2, 3}
	addition := 4
	have := addElements(slice, addition)
	slice[0] = -100
	notWant := slice
	assert.NotEqual(s.T(), notWant, have)
}

func (s *ServiceSuite) TestCopySliceEmptySlice() {
	slice := []int{}
	have := copySlice(slice)
	want := []int{}
	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestCopySliceNotEmptySlice() {
	slice := []int{3, 0, -100}
	have := copySlice(slice)
	want := []int{3, 0, -100}
	assert.Equal(s.T(), want, have)
}

func (s *ServiceSuite) TestCopySliceChangeSlice() {
	slice := []int{1, 2, 3}
	have := copySlice(slice)
	slice[0] = -100
	notWant := slice
	assert.NotEqual(s.T(), notWant, have)
}

func (s *ServiceSuite) TestRemoveElementOutOfRange() {
	slice := []int{}
	index := 1
	have, err := removeElement(slice, index)
	assert.Nil(s.T(), have)
	assert.Error(s.T(), err)
}

func (s *ServiceSuite) TestRemoveElementOneElement() {
	slice := []int{1}
	index := 0
	have, err := removeElement(slice, index)
	want := []int{}
	assert.Equal(s.T(), want, have)
	assert.Nil(s.T(), err)
}

func (s *ServiceSuite) TestRemoveElementAnotherElement() {
	slice := []int{1, 2, 3, 4, 5}
	index := 3
	have, err := removeElement(slice, index)
	want := []int{1, 2, 3, 5}
	assert.Equal(s.T(), want, have)
	assert.Nil(s.T(), err)
}

func (s *ServiceSuite) TestRemoveElementChangeSliceElement() {
	slice := []int{1, 2, 3, 4, 5}
	index := 3
	have, err := removeElement(slice, index)
	slice[0] = -100
	want := []int{1, 2, 3, 5}
	assert.Equal(s.T(), want, have)
	assert.Nil(s.T(), err)
}

func (s *ServiceSuite) TestRemoveElementOriginalSliceNotChanged() {
	have := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 3, 4, 5}

	_, err := removeElement(have, 3)

	assert.Equal(s.T(), want, have)
	assert.Nil(s.T(), err)
}