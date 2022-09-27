package linearsystems

import (
	"fmt"
)

const SIZE = 12
const SIZE1 = int(4.0)
const SIZE2 = 4

type Matrix struct {
	Matrix [SIZE1][SIZE2]float64
}

func FormatMatrixResult(matrix [SIZE1][SIZE2]float64) {
	for position := range matrix {
		for _, element := range matrix[position] {
			fmt.Printf("%.8f, ", element)
		}
	}

}

func CreateMatrix(array [SIZE1 * SIZE2]float64) [SIZE1][SIZE2]float64 {
	count := 0
	var matrix [SIZE1][SIZE2]float64

	for i := 0; i < SIZE1; i++ {
		for j := 0; j < SIZE2; j++ {
			matrix[i][j] = array[count]
			count++
		}
	}

	return matrix
}

func ChangeLines(matrix [SIZE1][SIZE2]float64, line1 int, line2 int) [SIZE1][SIZE2]float64 {
	aux := matrix

	matrix[line1-1] = aux[line2-1]
	matrix[line2-1] = aux[line1-1]

	return matrix
}

func MultiplyLineElements(matrix [SIZE1][SIZE2]float64, line int, constant float64) [SIZE1][SIZE2]float64 {
	for position, element := range matrix[line-1] {
		matrix[line-1][position] = element * constant
	}
	return matrix

}

func AddLineElements(matrix [SIZE1][SIZE2]float64, line1 int, line2 int) [SIZE1][SIZE2]float64 {
	aux := matrix

	for position := range matrix {
		aux[line1-1][position] = matrix[line1-1][position] + matrix[line2-1][position]
	}

	return aux
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

func Gauss(matrix [SIZE1][SIZE2]float64){
	for j:=0; j<SIZE2; j++{
		for i:=j;j<SIZE2; i++{
			if matrix[i][j] != 0{
				if i!=j {
					var temp float64
					for k:=0; k<SIZE1; k++{
						temp = matrix[i][k]
						matrix[i][k] = matrix[j][k]
						matrix[j][k] = temp
					}
				}
				for m:=j+1; m<SIZE1; m++{
					a := float64(-matrix[m][j])/float64(matrix[j][j])
					for n:=j; n<SIZE1; n++{
						matrix[m][n] += a * matrix[j][n]
					}
					FormatMatrixResult(matrix)
					break
				}
			} else {
				if (i==SIZE2-1){
                    fmt.Println("Sistema não possui solução")
                }
			}
		}
	}
}