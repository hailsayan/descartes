package main

import (
	"fmt"
)

func main() {
	var p, q int
	fmt.Scanf("%d %d", &p, &q)

	for i := 1; i <= q; i++ {
		if i%p == 0 {
			count := i / p
			for j := 0; j < count-1; j++ {
				fmt.Printf("%s ", "Hope")
			}
			fmt.Printf("Hope\n")
		} else {
			fmt.Println(i)
		}
	}
}