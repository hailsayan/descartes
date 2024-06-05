package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n := parseInt(scanner.Text())

	countryCodes := make(map[string]string)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Split(line, " ")
		country := parts[0]
		code := parts[1]
		countryCodes[code] = country
	}

	scanner.Scan()
	q := parseInt(scanner.Text())

	for i := 0; i < q; i++ {
		scanner.Scan()
		phone := scanner.Text()
		found := false

		for code, country := range countryCodes {
			if strings.HasPrefix(phone, code) {
				fmt.Println(country)
				found = true
				break
			}
		}

		if !found {
			fmt.Println("Invalid Number")
		}
	}
}

func parseInt(s string) int {
	var num int
	fmt.Sscanf(s, "%d", &num)
	return num
}
