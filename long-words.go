// https://codeforces.com/problemset/problem/71/A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var file_name string = "long-words-input.txt"
	str := readInput(file_name)
	encoded := encode(str)
	for i := range encoded {
		fmt.Println(encoded[i])
	}

}

func encode(array []string) []string {
	if len(array) < 1 || len(array) > 100 {
		fmt.Println("range invalid")
		os.Exit(0)
	}

	for i := range array {
		if len(array[i]) < 1 || len(array[i]) > 100 {
			fmt.Println("len ranges invalid")
		}
	}

	results := make([]string, 0)
	for i := range array {
		n := len(array[i])
		if n <= 10 {
			results = append(results, array[i])
		} else {
			results = append(results, string(array[i][0])+strconv.Itoa(len(array[i][0:n-2]))+string(array[i][n-1]))
		}
	}
	return results

}

func readInput(filename string) []string {
	// Step 1: Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		os.Exit(0)
	}
	defer file.Close() // Step 5: Close the file when we're done

	// Step 2: Create a new bufio.Scanner
	scanner := bufio.NewScanner(file)
	var i, j int
	s := make([]string, 0)

	// Step 3: Iterate over the lines
	for scanner.Scan() {
		// Step 4: Process each line
		line := scanner.Text()
		if j == 0 {
			i, _ = strconv.Atoi(line)

		}
		if j <= i {
			s = append(s, line)
		}
		j++
		//fmt.Println(line)
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning the file:", err)
	}

	return s
}
