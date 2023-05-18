package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const MATRIX_SIZE = 250

var (
	// Locker para ser usada na condicionalVariable
	rwLock = sync.RWMutex{}
	// cond variable. Note que ela vai olhar somente a leitura, uma vez que vamos
	// ter uma goroutine para cada linha da matriz. Ou seja, a escrita vai ser feita
	// por uma única goroutine.
	cond = sync.NewCond(rwLock.RLocker())
	// wg para esperar todas as goroutines terminarem
	wg = sync.WaitGroup{}
)

func generateMatrix(matrix *[MATRIX_SIZE][MATRIX_SIZE]int) {
	for i := 0; i < MATRIX_SIZE; i++ {
		for j := 0; j < MATRIX_SIZE; j++ {
			matrix[i][j] = rand.Intn(100) - 50
		}
	}
}

func rowMatrixMultiply(row int, matrixA, matrixB, result *[MATRIX_SIZE][MATRIX_SIZE]int) {
	rwLock.RLock()
	for {
		// Em vez de usar o defer wg.Done(), vamos usá-lo aqui após a primeira interação.
		// Isso porque, se usarmos o defer, corremos o risco do sinal ser enviado antes
		// de todas as goroutines estarem prontas para receber o sinal (estar na condição de wait).
		// Isso vai causar com que a gente perca esse sinal broadcast.
		wg.Done()
		cond.Wait()
		for j := 0; j < MATRIX_SIZE; j++ {
			for k := 0; k < MATRIX_SIZE; k++ {
				result[row][j] += matrixA[row][k] * matrixB[k][j]
			}
		}
	}
}
func main() {
	// The code without concurrency takes around 4s to run with MATRIX_SIZE = 250 and 100 iterations
	// The code with concurrency takes around 835.298166ms to run with MATRIX_SIZE = 250 and 100 iterations
	matrixA := [MATRIX_SIZE][MATRIX_SIZE]int{}
	matrixB := [MATRIX_SIZE][MATRIX_SIZE]int{}
	Result := [MATRIX_SIZE][MATRIX_SIZE]int{}
	start := time.Now()
	wg.Add(MATRIX_SIZE)
	for i := 0; i < MATRIX_SIZE; i++ {
		go rowMatrixMultiply(i, &matrixA, &matrixB, &Result)
	}
	fmt.Println("Starting matrix multiplication")
	for matrixNum := 0; matrixNum < 100; matrixNum++ {
		// Aqui nós esperamos para ter certeza que todos os goroutines estão prontas para receber o sinal,
		// caso contrário, podemos perder o sinal. Isso irá funcionar pois a gente está dando um wg.Done()
		// antes de dar o cond.Wait() na função rowMatrixMultiply.
		// Resumindo, o wg está servindo para garantir que todas as goroutines estão inicializadas com seus
		// respectivos locks.
		wg.Wait()
		rwLock.Lock()
		generateMatrix(&matrixA)
		generateMatrix(&matrixB)
		// Após gerar as matrizes, nós damos o sinal para todas as goroutines que estão esperando.
		wg.Add(MATRIX_SIZE)
		rwLock.Unlock()
		cond.Broadcast()
	}
	elapsed := time.Since(start)
	fmt.Println("Time taken: ", elapsed)
}
