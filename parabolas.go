// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func independentParabolas(y1 [3]int, y2 [3]int) bool {
	D := (y2[1]-y1[1])*(y2[1]-y1[1])-4*(y2[0]-y1[0])*(y2[2]-y1[2]) 
	fmt.Println(D)
	if  D <= 0 {
		return true
	}
	return false
}


func flagInput(){

}


func main() {
	//parabolas := [4][3]int32{{2, 2, -1}, {4, 5, 6}, {-1, 4, -5}, {1, 2, -4}}
	
	// parabolas := make([]int, 3, 4) // slice of length n and capacity m
	// parabolas[0] = []{2, 2, -1}

	parabolas := [4][3]int{ {1,2,-1},{-3, 0, -3}, {-1, 4, -5}, {1, 2,-4}}

	fmt.Println(independentParabolas(parabolas[1], parabolas[3]))
	fmt.Println(independentParabolas(parabolas[0], parabolas[3]))
	

	// for i, _ := range parabolas {
	// 	//count := 0
	// 	for j, _ := range parabolas {
	// 		if i != j {
	
	// 			if independentParabolas(parabolas[i], parabolas[j]) == true {
	// 			}
	// 			fmt.Println(i+1, j+1)

	// 		}

	// 	}
	// 		fmt.Println()
	
	// }
}
