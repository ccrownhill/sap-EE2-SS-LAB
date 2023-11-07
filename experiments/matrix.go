package main

import (
	"errors"
	"fmt"
)

func areRowsSameLen(m [][]float64) bool {
	length := -1
	for _, row := range m {
		if length == -1 {
			length = len(row)
		} else if length != len(row) {
			return false
		}
	}
	return true
}

func matrixVectorMult(m [][]float64, v []float64) ([]float64, error) {
	if (len(v) > 0 && len(m) <= 0) ||
		(len(m) > 0 && len(m[0]) != len(v)) ||
		(!areRowsSameLen(m)) {
		return []float64{-1}, errors.New("matrix vector dimension mismatch")
	}

	out := make([]float64, len(m))

	for ri, row := range m {
		for ci, el := range row {
			out[ri] += v[ci] * el
		}
	}

	return out, nil
}

// convolute input with impulse response h to get output of LTI system
func ltiConvolution(x, h []float64) ([]float64, error) {
	convMat := make([][]float64, len(x)+len(h)-1)

	// fill matrix used for convolution
	for ri := range convMat {
		convMat[ri] = make([]float64, len(x))
		for ci := range convMat[ri] {
			if ri-ci >= 0 && ri-ci < len(h) {
				convMat[ri][ci] = h[ri-ci]
			} else {
				convMat[ri][ci] = 0
			}
		}
	}

	return matrixVectorMult(convMat, x)

}

func main() {
	fmt.Println("Hello world")
}
