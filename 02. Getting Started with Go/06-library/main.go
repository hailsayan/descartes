package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	books := make(map[int]string)

	for i := 0; i < n; i++ {
		scanner.Scan()
		inp := scanner.Text()
		inpSpl := strings.SplitN(inp, " ", 3) // Use SplitN to correctly handle titles with spaces

		command := inpSpl[0]
		isbn, _ := strconv.Atoi(inpSpl[1])

		if command == "ADD" {
			if len(inpSpl) == 3 {
				title := inpSpl[2]
				books[isbn] = title
			}
		} else if command == "REMOVE" {
			delete(books, isbn)
		}
	}

	var isbns []int
	for isbn := range books {
		isbns = append(isbns, isbn)
	}

	sort.Slice(isbns, func(i, j int) bool {
		titleI, titleJ := books[isbns[i]], books[isbns[j]]
		if titleI == titleJ {
			return isbns[i] < isbns[j] // Sort by ISBN if titles are the same
		}
		return titleI < titleJ // Sort by title alphabetically
	})

	for _, isbn := range isbns {
		fmt.Println(isbn)
	}
}
