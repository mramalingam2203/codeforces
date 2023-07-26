// https://codeforces.com/problemset/problem/4/A

package main

import "fmt"

func main() {

	var i int
	fmt.Scanf("%d", &i)

	if i < 1 || i > 100 {
		fmt.Println("NO")
		return
	}

	if i%2 == 0 {
		fmt.Println("YES")
		return
	} else {
		fmt.Println("NO")
		return
	}

}
