package main

import "math"

func PipelineJobPow3(input <-chan uint8, output chan<- float64) {
	output <- math.Pow(float64(<-input), 3)
}