package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var vector []int
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		os.Exit(1)
	}

	arrSize := 0
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan() && i <= arrSize; i++ {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		if i == 0 {
			arrSize = number
			continue
		}

		vector = append(vector, number)
	}

	current := 0
	max := 0

	for _, val := range vector {
		if val == 1 {
			current++
			if max < current {
				max = current
			}
		} else {
			current = 0
		}
	}

	fo, _ := os.Create("output.txt")
	defer fo.Close()

	fmt.Fprint(fo, max)
}
