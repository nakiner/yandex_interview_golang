package main

import (
	"fmt"
)

func main() {
	var count int
	fmt.Scanf("%d", &count)
	gen("", 0, 0, count)
}

func gen(res string, l, r, counts int) {
	if r == counts && l == counts {
		fmt.Println(res)
	} else {
		if l < counts {
			gen(res+"(", l+1, r, counts)
		}
		if r < l {
			gen(res+")", l, r+1, counts)
		}
	}
}
