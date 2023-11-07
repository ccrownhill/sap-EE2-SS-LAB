package main

import (
	"reflect"
	"testing"
)

func TestMatrixVectorMult(t *testing.T) {
	testMatrixVectorMult := func(t *testing.T, matrix [][]float64,
		vector, want []float64) {
		got, err := matrixVectorMult(matrix, vector)

		if err != nil {
			t.Errorf("got error: %q", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v want: %v", got, want)
		}
	}

	t.Run("test 1", func(t *testing.T) {
		matrix := [][]float64{{1, 2}, {3, 4}}
		vector := []float64{5, 6}
		want := []float64{17, 39}
		testMatrixVectorMult(t, matrix, vector, want)
	})

	t.Run("test 2", func(t *testing.T) {
		matrix := [][]float64{{1, 0}, {0, 1}}
		vector := []float64{10, 21}
		want := []float64{10, 21}
		testMatrixVectorMult(t, matrix, vector, want)
	})
}

func TestLtiConvolution(t *testing.T) {
	convTest := func(t *testing.T, x, h, want []float64) {
		got, err := ltiConvolution(x, h)
		if err != nil {
			t.Errorf("got error: %q", err)
		}
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v got %v", want, got)
		}
	}

	t.Run("test 1", func(t *testing.T) {
		x := []float64{1, 1}
		h := []float64{3, 1}
		want := []float64{3, 4, 1}
		convTest(t, x, h, want)
	})

	t.Run("test 2", func(t *testing.T) {
		x := []float64{2, 3, -2}
		h := []float64{1, -0.5}
		want := []float64{2, 2, -3.5, 1}
		convTest(t, x, h, want)
	})
}
