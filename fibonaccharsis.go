// https://codeforces.com/problemset/problem/1853/B

package main

import (
	"fmt"
)

func main() {
	// read number of test cases
	var numRows, k int
	fmt.Scanf("%d %d", &numRows, &k)

	numCols := 2 // Petya, Vasya, Tonya scores
	numbers := make([][]int, numRows)

	for i := range numbers {
		numbers[i] = make([]int, numCols)
	}

	for i := 0; i < numRows; i++ {
		fmt.Scanf("%d %d", &numbers[i][0], &numbers[i][1])
	}

	// Fib1 := fibonacciSequence(5)
	// Fib2 := fibonacciSequence(6)
	// noFib := []int{6, 8, 14, 23}

	// fmt.Println(Fib1, Fib2)
	// fmt.Println(checkFibonacci(Fib1), checkFibonacci(Fib2), checkFibonacci(noFib))

	seq := generateNormalSequnce(1, n)
	fmt.Println(seq)
	//generateCombinationsOfSize()
}

func checkFibonacci(numbers []int) bool {

	for i := 2; i < len(numbers); i++ {
		if !(numbers[i] == numbers[i-1]+numbers[i-2]) {
			return false
		}
	}

	return true
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

// backtrack generates all combinations of the given size (k) from the array.
func backtrack(nums []int, start, k int, path []int, result *[][]int) {
	if len(path) == k {
		// Add the current combination to the result
		combination := make([]int, len(path))
		copy(combination, path)
		*result = append(*result, combination)
		return
	}

	// Backtrack to explore other combinations
	for i := start; i < len(nums); i++ {
		// Add the current element to the path
		path = append(path, nums[i])
		// Recurse with the next element
		backtrack(nums, i+1, k, path, result)
		// Remove the last element to backtrack
		path = path[:len(path)-1]
	}
}

func generateCombinationsOfSize(nums []int, k int) [][]int {
	result := [][]int{}
	backtrack(nums, 0, k, []int{}, &result)
	return result
}

func generateNormalSequnce(low, high int) []int {
	result := make([]int, high-low+2)
	for i := low; i <= high; i++ {
		result[i] = i
	}

	return result

}
