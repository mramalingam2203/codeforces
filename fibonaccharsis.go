// https://codeforces.com/problemset/problem/1853/B

package main

import (
	"fmt"
)

func main() {
	// read number of test cases
	var numRows int
	fmt.Scanf("%d", &numRows)

	numCols := 2 // Petya, Vasya, Tonya scores
	numbers := make([][]int, numRows)

	for i := range numbers {
		numbers[i] = make([]int, numCols)
	}

	for i := 0; i < numRows; i++ {
		fmt.Scanf("%d %d", &numbers[i][0], &numbers[i][1])
	}

	fmt.Println(numbers)
	
	yesFib := [5]int {}

}

func checkFibonacci(numbers []float64) bool {

	for i := range numbers {
		for i > 2 {
			if !numbers[i] == numbers[i-1]+numbers[i-2] {
				return false
			}
		}
	}


}

func fibonacciSequence(length int) []int {
	if length <= 0 {
		return []int{}
	}

	sequence := make([]int, length)
	sequence[0], sequence[1] = 0, 1

	for i := 2; i < length; i++ {
		sequence[i] = sequence[i-1] + sequence[i-2]
	}

	return sequence
}
	return true
}
