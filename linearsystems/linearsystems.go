package linearsystems

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

const SIZE = 12
const SIZE1 = int(4.0)
const SIZE2 = 4

type Matrix struct {
	Matrix [SIZE1][SIZE2]float64
}

func FormatMatrixResult(matrix mat.Dense) {
	fa := mat.Formatted(&matrix, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func CreateMatrix(array []float64, col int, row int) *mat.Dense {
	matrix := mat.NewDense(row, col, array)
	return matrix
}

func getVectorFromRow(row mat.Vector) []float64{
	v := make([]float64, row.Len())
	for i := 0; i < row.Len(); i++ {
		v[i] = row.At(i,0)
	}
	return v
}

func ChangeLines(matrix mat.Dense, line1 int, line2 int) mat.Dense {
	firstLine := getVectorFromRow(matrix.RowView(line1-1))
	secondLine := getVectorFromRow(matrix.RowView(line2-1))

	matrix.SetRow(line2-1, firstLine)
	matrix.SetRow(line1-1, secondLine)

	return matrix
}

func MultiplyLineElements(matrix mat.Dense, line int, constant float64) mat.Dense {
	new_row := getVectorFromRow(matrix.RowView(line-1))
	
	for index, elements := range new_row{
		new_row[index] = elements * constant
	}

	matrix.SetRow(line-1, new_row)

	return matrix
}

func AddLineElements(matrix mat.Dense, line1 int, line2 int) mat.Dense {
	firstLine := mat.VecDenseCopyOf(matrix.RowView(line1-1))
	secondLine := mat.VecDenseCopyOf(matrix.RowView(line2-1))

	newRow := 

	matrix.SetRow(line-1, firstLine)

	return matrix
}

func AddLineWeightedElements(matrix [SIZE1][SIZE2]float64, line1 int, line2 int, weight float64) [SIZE1][SIZE2]float64 {
	aux := matrix

	for position := range matrix[line1] {
		aux[line1-1][position] = matrix[line1-1][position] + (matrix[line2-1][position] * weight)
	}

	return aux
}

func Jacobi(matrix [SIZE1][SIZE2]float64, results [SIZE1]float64, estimation [SIZE1]float64, n int, p [SIZE]int){
	iteracao := 1
    i2 := 0;

    for k:=0; k<n;k++{
        for i:=0; i<SIZE1; i++{
            bi:= results[i];
            for j:=0; j<SIZE2;j++{
                if j != i {
					bi -= matrix[i][j]*estimation[j];

				}
            }
            bi/=matrix[i][i];
            if iteracao == p[i2]{
                fmt.Printf("%.16f,\n", bi)
            }
            estimation[i]=bi;
        }
        if iteracao == p[i2]{
            i2++;
        }
        iteracao++;
    }
}

// func Gauss(matrix [SIZE1][SIZE2]float64){
// 	for j:=0; j<SIZE2; j++{
// 		for i:=j;j<SIZE2; i++{
// 			if matrix[i][j] != 0{
// 				if i!=j {
// 					var temp float64
// 					for k:=0; k<SIZE1; k++{
// 						temp = matrix[i][k]
// 						matrix[i][k] = matrix[j][k]
// 						matrix[j][k] = temp
// 					}
// 				}
// 				for m:=j+1; m<SIZE1; m++{
// 					a := float64(-matrix[m][j])/float64(matrix[j][j])
// 					for n:=j; n<SIZE1; n++{
// 						matrix[m][n] += a * matrix[j][n]
// 					}
// 					FormatMatrixResult(matrix)
// 					break
// 				}
// 			} else {
// 				if (i==SIZE2-1){
//                     fmt.Println("Sistema não possui solução")
//                 }
// 			}
// 		}
// 	}
// }