package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SpaceShip struct {
	Name            string
	ArithmeticCount int
}

func countArithmeticSequences(data []int) int {
	count := 0
	n := len(data)
	for i := 0; i < n-2; i++ {
		diff := data[i+1] - data[i]
		for j := i + 2; j < n; j++ {
			if data[j]-data[j-1] == diff {
				count++
			} else {
				break
			}
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	spaceships := make([]SpaceShip, 0, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		words := strings.Fields(scanner.Text())
		name := words[0]
		fuelData := make([]int, len(words)-1)
		for j := 1; j < len(words); j++ {
			fuelData[j-1], _ = strconv.Atoi(words[j])
		}
		count := countArithmeticSequences(fuelData)
		spaceships = append(spaceships, SpaceShip{Name: name, ArithmeticCount: count})
	}

	sort.Slice(spaceships, func(i, j int) bool {
		if spaceships[i].ArithmeticCount == spaceships[j].ArithmeticCount {
			return spaceships[i].Name < spaceships[j].Name
		}
		return spaceships[i].ArithmeticCount > spaceships[j].ArithmeticCount
	})

	for _, spaceship := range spaceships {
		fmt.Printf("%s %d\n", spaceship.Name, spaceship.ArithmeticCount)
	}
}
