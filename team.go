// https://codeforces.com/problemset/problem/231/A

package main

import (
	"fmt"
	"os"
)

func main() {
	var numRows int
	fmt.Scanf("%d", &numRows)

	numCols := 3 // Petya, Vasya, Tonya scores
	scores := make([][]int, numRows)

	for i := range scores {
		scores[i] = make([]int, numCols)
	}

	count := 0
	for i := 0; i < numRows; i++ {
		fmt.Scanf("%d %d %d", &scores[i][0], &scores[i][1], &scores[i][2])
		if scores[i][0] != 0 || scores[i][0] != 1 || scores[i][1] != 0 || scores[i][1] != 1 || scores[i][2] != 0 || scores[i][2] != 1 {

			fmt.Println("Invalid")
			os.Exit(0)
		}
		if scores[i][0]+scores[i][1]+scores[i][2] >= 2 {
			count++
		}
	}
}
