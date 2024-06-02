package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// X >= 80 returns Excellent
// X >= 60 && X < 80 returns Very Good
// X >= 40 && X < 60 returns Good
// else returns Fair

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	words := strings.Fields(scanner.Text())

	n, _ := strconv.Atoi(words[0])

	for i := 0; i < n; i++ {
		// teachers
		scanner.Scan()
		teacherName := scanner.Text()

		// students' percentages
		scanner.Scan()
		words = strings.Fields(scanner.Text())

		total := 0
		for _, word := range words {
			score, _ := strconv.Atoi(word)
			total += score
		}

		average := float64(total) / float64(len(words))

		evaluation := ""
		if average >= 80 {
			evaluation = "Excellent"
		} else if average >= 60 {
			evaluation = "Very Good"
		} else if average >= 40 {
			evaluation = "Good"
		} else {
			evaluation = "Fair"
		}
		fmt.Printf("%s %s\n", teacherName, evaluation)
	}

}