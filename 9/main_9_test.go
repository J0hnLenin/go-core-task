package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PipelineJobSuite struct {
	suite.Suite
}

func TestPipelineJobTestSuite(t *testing.T) {
	suite.Run(t, new(PipelineJobSuite))
}

func (s *PipelineJobSuite) TestPipelineJobPow3_Zero() {
	input := make(chan uint8, 1)
	output := make(chan float64, 1)

	input <- 0
	go PipelineJobPow3(input, output)

	result := <-output
	assert.Equal(s.T(), 0.0, result)
	assert.InDelta(s.T(), 0.0, result, 0.0001)
}

func (s *PipelineJobSuite) TestPipelineJobPow3_One() {
	input := make(chan uint8, 1)
	output := make(chan float64, 1)

	input <- 1
	go PipelineJobPow3(input, output)

	result := <-output
	assert.Equal(s.T(), 1.0, result)
	assert.InDelta(s.T(), 1.0, result, 0.0001)
}

func (s *PipelineJobSuite) TestPipelineJobPow3_MaxUint8() {
	input := make(chan uint8, 1)
	output := make(chan float64, 1)

	input <- 255
	go PipelineJobPow3(input, output)

	result := <-output
	expected := 255.0 * 255.0 * 255.0
	assert.InDelta(s.T(), expected, result, 0.0001)
}

func (s *PipelineJobSuite) TestPipelineJobPow3_TableTest() {
	testCases := []struct {
		name     string
		input    uint8
		expected float64
	}{
		{"3 cubed", 3, 27.0},
		{"4 cubed", 4, 64.0},
		{"5 cubed", 5, 125.0},
		{"6 cubed", 6, 216.0},
		{"10 cubed", 10, 1000.0},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			input := make(chan uint8, 1)
			output := make(chan float64, 1)

			input <- tc.input
			go PipelineJobPow3(input, output)

			result := <-output
			assert.InDelta(t, tc.expected, result, 0.0001)
		})
	}
}