package main

import (
	"fmt"
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

func (s *ServiceSuite) TestInitDataHaveInt() {
	data := initData()
	
	have := make([]string, len(data))
	for i, value := range(data) {
		have[i] = fmt.Sprintf("%T", value)
	}
	want := fmt.Sprintf("%T", 1)
	assert.Contains(s.T(), have, want)
}

func (s *ServiceSuite) TestInitDataHaveString() {
	data := initData()
	
	have := make([]string, len(data))
	for i, value := range(data) {
		have[i] = fmt.Sprintf("%T", value)
	}

	want := fmt.Sprintf("%T", "str")
	assert.Contains(s.T(), have, want)
}

func (s *ServiceSuite) TestInitDataHaveBool() {
	data := initData()
	
	have := make([]string, len(data))
	for i, value := range(data) {
		have[i] = fmt.Sprintf("%T", value)
	}

	want := fmt.Sprintf("%T", true)
	assert.Contains(s.T(), have, want)
}

func (s *ServiceSuite) TestInitDataHaveFloat() {
	data := initData()
	
	have := make([]string, len(data))
	for i, value := range(data) {
		have[i] = fmt.Sprintf("%T", value)
	}

	want := fmt.Sprintf("%T", 0.5)
	assert.Contains(s.T(), have, want)
}


func (s *ServiceSuite) TestInitDataHaveComplex() {
	data := initData()
	
	have := make([]string, len(data))
	for i, value := range(data) {
		have[i] = fmt.Sprintf("%T", value)
	}

	want := fmt.Sprintf("%T", 0 + 1i)
	assert.Contains(s.T(), have, want)
}

func (s *ServiceSuite) TestToStringAndConcatinateWithInt() {
	values := make([]any, 2)
	values[0] = 10
	values[1] = 20
	
	have := toStringAndConcatinate(values)
	want := "10 20"

	assert.Equal(s.T(), have, want)
}

func (s *ServiceSuite) TestToStringAndConcatinateWithBool() {
	values := make([]any, 2)
	values[0] = true
	values[1] = false
	
	have := toStringAndConcatinate(values)
	want := "true false"

	assert.Equal(s.T(), have, want)
}

func (s *ServiceSuite) TestToStringAndConcatinateWithString() {
	values := make([]any, 2)
	values[0] = "aba"
	values[1] = "caba"
	
	have := toStringAndConcatinate(values)
	want := "aba caba"

	assert.Equal(s.T(), have, want)
}

func (s *ServiceSuite) TestToStringAndConcatinateWithComplex() {
	values := make([]any, 2)
	values[0] = 0 + 2i
	values[1] = -1 - 10i
	
	have := toStringAndConcatinate(values)
	want := "(0+2i) (-1-10i)"

	assert.Equal(s.T(), have, want)
}

func (s *ServiceSuite) TestToStringAndConcatinateWithDifferentTypes() {
	values := make([]any, 4)
	values[0] = 1
	values[1] = "aba"
	values[2] = true
	values[3] = 1 + 4i
	
	have := toStringAndConcatinate(values)
	want := "1 aba true (1+4i)"

	assert.Equal(s.T(), have, want)
}

func (s *ServiceSuite) TestHashNotEqualObjects() {
	obj1 := []rune("abacaba")
	obj2 := []rune("cababac")
	
	have := hash(obj1)
	notWant := hash(obj2)

	assert.NotEqual(s.T(), have, notWant)
}

func (s *ServiceSuite) TestHashEqualObjects() {
	obj1 := []rune("abacaba")
	obj2 := []rune("abacaba")
	
	have := hash(obj1)
	want := hash(obj2)

	assert.Equal(s.T(), have, want)
}