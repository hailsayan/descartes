package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	values := strings.Split(line, " ")
	p, _ := strconv.Atoi(values[0])
	q, _ := strconv.Atoi(values[1])

	for i := 1; i <= q; i++ {
		if i%p == 0 {
			count := i / p
			for j := 0; j < count; j++ {
				fmt.Print("Hope ")
			}
			fmt.Println()
		} else {
			fmt.Println(i)
		}
	}
}