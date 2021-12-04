package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, j string

	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &j)

	count := 0

	for _, char := range j {
		if strings.ContainsRune(s, char) {
			count++
		}
	}

	fmt.Println(count)
}
