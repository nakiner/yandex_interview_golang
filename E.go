package main

import (
	"fmt"
)

func main() {
	var one, two string

	fmt.Scanf("%s", &one)
	fmt.Scanf("%s", &two)

	counts := map[rune]int{}

	if len(one) != len(two) {
		fmt.Println(0)
		return
	}

	for _, char := range one {
		if _, ok := counts[char]; ok {
			counts[char] += 1
		} else {
			counts[char] = 1
		}
	}

	for _, char := range two {
		if _, ok := counts[char]; ok {
			counts[char] -= 1
		}
	}

	for _, val := range counts {
		if val != 0 {
			fmt.Println(0)
			return
		}
	}

	fmt.Println(1)
}
