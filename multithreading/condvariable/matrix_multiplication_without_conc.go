package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MATRIX_SIZE = 250

func generateMatrix(matrix *[MATRIX_SIZE][MATRIX_SIZE]int) {
	for i := 0; i < MATRIX_SIZE; i++ {
		for j := 0; j < MATRIX_SIZE; j++ {
			matrix[i][j] = rand.Intn(100) - 50
		}
	}
}

func rowMatrixMultiply(row int, matrixA, matrixB, result *[MATRIX_SIZE][MATRIX_SIZE]int) {
	for j := 0; j < MATRIX_SIZE; j++ {
		for k := 0; k < MATRIX_SIZE; k++ {
			result[row][j] += matrixA[row][k] * matrixB[k][j]
		}
	}
}
func main() {
	matrixA := [MATRIX_SIZE][MATRIX_SIZE]int{}
	matrixB := [MATRIX_SIZE][MATRIX_SIZE]int{}
	Result := [MATRIX_SIZE][MATRIX_SIZE]int{}
	start := time.Now()
	for matrixNum := 0; matrixNum < 100; matrixNum++ {
		generateMatrix(&matrixA)
		generateMatrix(&matrixB)
		for i := 0; i < MATRIX_SIZE; i++ {
			rowMatrixMultiply(i, &matrixA, &matrixB, &Result)
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Time taken: ", elapsed)
}
