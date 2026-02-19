// B. Another Problem about Beautiful Pairs


// https://codeforces.com/contest/2196/problem/B

package main

import "fmt"

func countBeautifulPairs(a []int){
	n := len(a)

	for i:= 0; i < n; i++{
		for j:=0; j < n; j++{
			if a[i]*a[j] == j-i{
				fmt.Println(i+1,j+1)
			}
		}

	}

}



func genRandArray(size int)[]int{

	//const size = 20000
	// Create a slice of int with the specified size.
	randomNumbers := make([]int, size)

	// Fill the slice with random numbers.
	for i := 0; i < size; i++ {
		// rand.Intn(100) generates a number in the range [0, 100)
		randomNumbers[i] = rand.Intn(100)
	}

	// Optional: Print a small part of the slice to verify
	fmt.Println("First 10 random numbers:", randomNumbers[:10])
	fmt.Println("Total numbers generated:", len(randomNumbers))

	return randomNumbers
}



func main(){
	randomNumbers(20000)
 	//array := []int{1, 1, 2, 100, 4}
 	array := []int{2, 2, 1, 1, 2, 2}

 	countBeautifulPairs(array)

}

