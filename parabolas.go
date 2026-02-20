// You can edit this code!
// Click here and start typing.

//https://codeforces.com/problemset/problem/2195/F

package main

import (
	"fmt"
	"sort"
	"strings"
)


func independentParabolas(y1 []int, y2 []int) bool {
	D := (y2[1]-y1[1])*(y2[1]-y1[1])-4*(y2[0]-y1[0])*(y2[2]-y1[2]) 
	//fmt.Println(D)
	if  D <= 0 {
		return true
	}
	return false
}





func checkForIndependentParabolas(idxSubsets[][]int, parabolas [][]int ){
	// nParabolas := len(parabolas)
	// nIndexSubsets := len(idxSubsets)
	fmt.Println(parabolas)
	// fmt.Println(nParabolas, nIndexSubsets)
	
	fmt.Println(len(idxSubsets[0]))
	
	for idx,_ := range idxSubsets{
		for j := 0; j < len(idxSubsets[idx]); j++{
			for k:=j+1; k < len(idxSubsets[idx]); k++ {
				fmt.Println(j+1, k+1 , independentParabolas(parabolas[idxSubsets[idx][j]-1], parabolas[idxSubsets[idx][k]-1]))
				//fmt.Println(idxSubsets[idx][j], idxSubsets[idx][k])
			}
		}
		fmt.Println()
	}
		
	

}


func getUniqueSubsets(n int) [][]int {

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
	indexSubsets := getUniqueSubsets(len(parabolas))
	checkForIndependentParabolas(indexSubsets, parabolas)

//	fmt.Println("Unique Subsets:", indexSubsets)
}



/*

1=find all subsets of the given set
2=filter all subsets containting a given parabola to get a list
3-choose a subset from the above list such that all parabolas in that subset are independent of each other
4-repeat(2-3)  for all parabolas

*/
