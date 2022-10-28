package linearsystems

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

const SIZE = 12
const SIZE1 = int(4.0)
const SIZE2 = 4

func FormatMatrixResult(matrix mat.Dense) {
	fa := mat.Formatted(&matrix, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func CreateMatrix(array []float64, col int, row int) *mat.Dense {
	matrix := mat.NewDense(row, col, array)
	return matrix
}

func getVectorFromRow(row mat.Vector) []float64 {
	v := make([]float64, row.Len())
	for i := 0; i < row.Len(); i++ {
		v[i] = row.At(i, 0)
	}
	return v
}

func ChangeLines(matrix mat.Dense, lineMain int, lineRef int) {
	firstLine := getVectorFromRow(matrix.RowView(lineMain - 1))
	secondLine := getVectorFromRow(matrix.RowView(lineRef - 1))

	matrix.SetRow(lineRef-1, firstLine)
	matrix.SetRow(lineMain-1, secondLine)

}

func MultiplyLineElements(matrix mat.Dense, line int, constant float64) {
	op := mat.VecDenseCopyOf(matrix.RowView(line - 1))

	op.ScaleVec(constant, op)

	matrix.SetRow(line-1, getVectorFromRow(op))

}

func AddLineElements(matrix mat.Dense, lineMain int, lineRef int) {
	firstLine := mat.VecDenseCopyOf(matrix.RowView(lineMain - 1))
	secondLine := mat.VecDenseCopyOf(matrix.RowView(lineRef - 1))

	newRow := mat.NewVecDense(firstLine.Len(), nil)
	newRow.AddVec(firstLine, secondLine)

	matrix.SetRow(lineMain-1, getVectorFromRow(newRow))

}

func AddLineWeightedElements(matrix mat.Dense, lineMain int, lineRef int, weight float64) {
	secondLine := mat.VecDenseCopyOf(matrix.RowView(lineRef - 1))

	secondLine.ScaleVec(weight, secondLine)
	secondLine.AddVec(secondLine, matrix.RowView(lineMain-1))

	matrix.SetRow(lineMain-1, getVectorFromRow(secondLine))

}

func LinearSolve(matrix mat.Dense, results mat.VecDense) mat.VecDense {
	err := results.SolveVec(&matrix, &results)
	if err != nil {
		fmt.Printf("unexpected error from dense vector solve: %v", err)
	}
	return results
	
}

func Jacobi(matrix mat.Dense, results mat.VecDense, estimation mat.VecDense, p mat.VecDense, n int) {
	iteracao := 1
	i2 := 0

	for k := 0; k < n; k++ {
		for i := 0; i < SIZE1; i++ {
			bi := results.AtVec(i)
			for j := 0; j < SIZE2; j++ {
				if j != i {
					bi -= matrix.At(i, j) * estimation.AtVec(j)

				}
			}
			bi /= matrix.At(i, i)
			if iteracao == int(p.AtVec(i2)) {
				fmt.Printf("%.16f,\n", bi)
			}
			estimation.SetVec(i, bi)
		}
		if iteracao == int(p.AtVec(i2)) {
			i2++
		}
		iteracao++
	}
}

func Gauss(matrix mat.Dense) {
	for j := 0; j < matrix.RawMatrix().Cols-1; j++ {
		for i := j; j < matrix.RawMatrix().Cols-1; i++ {
			if matrix.At(i, j) != 0 {
				if i != j {
					var temp float64
					for k := 0; k < matrix.RawMatrix().Rows-1; k++ {
						temp = matrix.At(i, k)
						matrix.Set(i, k, matrix.At(j, k))
						matrix.Set(j, k, temp)
					}
				}
				for m := j + 1; m < matrix.RawMatrix().Rows-1; m++ {
					a := float64(-matrix.At(m, j)) / float64(matrix.At(j, j))
					for n := j; n < matrix.RawMatrix().Rows; n++ {
						matrix.Set(m, n, (matrix.At(m, n) + (a * matrix.At(j, n))))
					}
					FormatMatrixResult(matrix)
					break
				}
			} else {
				if i == matrix.RawMatrix().Cols-1 {
					fmt.Println("Sistema não possui solução")
				}
			}
		}
	}
}
