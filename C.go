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

	seen := map[int]struct{}{}
	scanner := bufio.NewScanner(file)
	arrSize := 0
	for i := 0; scanner.Scan() && i <= arrSize; i++ {
		tx := scanner.Text()
		number, _ := strconv.Atoi(tx)
		if err != nil {
			continue
		}
		if i == 0 {
			arrSize = number
			continue
		}

		_, ok := seen[number]
		if !ok {
			seen[number] = struct{}{}
			vector = append(vector, number)
		}
	}

	fo, _ := os.Create("output.txt")
	defer fo.Close()

	for _, value := range vector {
		fmt.Fprintln(fo, value)
	}
}
