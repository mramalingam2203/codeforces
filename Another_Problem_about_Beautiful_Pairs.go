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


func main(){

 	//array := []int{1, 1, 2, 100, 4}
 	array := []int{2, 2, 1, 1, 2, 2}
 	countBeautifulPairs(array)

}

