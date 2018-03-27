package go_playground

import "errors"

func CreateMatrix_empty(rows int, columns int) ([][]float64, error) {
	if rows < 1 || columns < 1 {
		return nil, errors.New("invalid matrix size")
	}
	matrix := createMatrix(
		func(rowIdx int, columns int) []float64 {
			return make([]float64, columns)
		},
		rows,
		columns)
	return matrix, nil
}

func CreateMatrix_initialized(val float64, rows int, columns int) ([][]float64, error) {
	if rows < 1 || columns < 1 {
		return nil, errors.New("invalid matrix size")
	}
	raw := make([]float64, rows * columns)
	for i := range raw {
		raw[i] = val
	}
	matrix := createMatrix(
		func(rowIdx int, columns int) []float64 {
			base := rowIdx * columns
			return raw[base : base + columns]
		},
		rows,
		columns)
	return matrix, nil
}

func createMatrix(fil filler, rows int, columns int) [][]float64 {
	matrix := make([][]float64, rows)
	for i := range matrix {
		matrix[i] = fil(i, columns)
	}
	return matrix
}

type filler func(rowIdx int, columns int) []float64
