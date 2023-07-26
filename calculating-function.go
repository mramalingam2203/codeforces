// https://codeforces.com/problemset/problem/486/A

package main

import (
	"fmt"
)

func main() {
	fmt.Println(function(1e15))
}

func function(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum += i
		} else {
			sum += (-1 * i)
		}
	}
	return sum
}
