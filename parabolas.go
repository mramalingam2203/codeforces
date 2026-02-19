// You can edit this code!
// Click here and start typing.

//https://codeforces.com/problemset/problem/2195/F

package main

import (
	"fmt"
	"sort"
	"strings"
)


func independentParabolas(y1 [3]int, y2 [3]int) bool {
	D := (y2[1]-y1[1])*(y2[1]-y1[1])-4*(y2[0]-y1[0])*(y2[2]-y1[2]) 
	//fmt.Println(D)
	if  D <= 0 {
		return true
	}
	return false
}


func flagInput(){

}

/*
func main() {
	//parabolas := [4][3]int32{{2, 2, -1}, {4, 5, 6}, {-1, 4, -5}, {1, 2, -4}}
	
	// parabolas := make([]int, 3, 4) // slice of length n and capacity m
	// parabolas[0] = []{2, 2, -1}

	parabolas := [4][3]int{ {1,2,-1},{-3, 0, -3}, {-1, 4, -5}, {1, 2,-4}}

	fmt.Println(independentParabolas(parabolas[1], parabolas[3]))
	fmt.Println(independentParabolas(parabolas[0], parabolas[3]))
	
	

	for i, _ := range parabolas {
		//count := 0
		for j, _ := range parabolas {
			if i != j {
	
				if independentParabolas(parabolas[i], parabolas[j]) == true {
								fmt.Println(i+1, j+1)

				}
	
			}

		}
			fmt.Println()
	
	}
}

*/



func getUniqueSubsets(parabolas[][]int) [][]int {
	n := len(parabolas)

	nums := make([]int, n)

	for index, _ := range nums{
		nums[index] = index +1
	}

	subsetMap := make(map[string][]int)
	
	// 2. Iterate through all 2^n bitmasks
	numSubsets := 1 << n
	for i := 0; i < numSubsets; i++ {
		var currentSubset []int
		var sb strings.Builder
		
		for j := 0; j < n; j++ {
			// Check if the j-th bit is set in the mask i
			if (i >> j) & 1 == 1 {
				currentSubset = append(currentSubset, nums[j])
				sb.WriteString(fmt.Sprintf("%d,", nums[j]))
			}
		}
		
		// 3. Use the string representation as a key in a map to filter duplicates
		key := sb.String()
		if _, exists := subsetMap[key]; !exists {
			subsetMap[key] = currentSubset
		}
	}

	// Convert map values back to a slice
	var subsets [][]int
	for _, subset := range subsetMap {
		if len(subset) >= 2{
			subsets = append(subsets, subset)
		}
	}

	// Sort slice of slices descending by inner length
	sort.Slice(subsets, func(i, j int) bool {
		return len(subsets[i]) > len(subsets[j]) // > for descending
	})

	return subsets
}

func main() {
	//nums := []int{1, 2, 3, 4}
	parabolas := [][]int{ {1,2,-1},{-3, 0, -3}, {-1, 4, -5}, {1, 2,-4}}
	subsets := getUniqueSubsets(parabolas)
	fmt.Println("Unique Subsets:", subsets)
}



/*

1=find all subsets of the given set
2=filter all subsets containting a given parabola to get a list
3-choose a subset from the above list such that all parabolas in that subset are independent of each other
4-repeat(2-3)  for all parabolas

*/
