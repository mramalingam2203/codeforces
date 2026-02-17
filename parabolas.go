// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func independentParabolas(y1 [3]int, y2 [3]int) bool {
	if (y2[1]-y1[1])*(y2[1]-y1[1])-4*(y2[0]-y1[0])*(y2[2]-y2[2]) <= 0 {
		return true
	}
	return false
}

func main() {
	//parabolas := [4][3]int32{{2, 2, -1}, {4, 5, 6}, {-1, 4, -5}, {1, 2, -4}}
	
	parabolas := make([]int, 3, 4) // slice of length n and capacity m
	parabolas[0] = []{2, 2, -1}



	for i, _ := range parabolas {
		count := 0
		for j, _ := range parabolas {
			if i != j {
	
				if independentParabolas(parabolas[i], parabolas[j]) == true {
					count++
				}
			}
		}

		fmt.Println(count)

	}
}
