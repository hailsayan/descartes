package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Book struct {
	ISBN  int
	Title string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	library := make(map[int]string)

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		parts := strings.Fields(line)

		command := parts[0]
		isbn, _ := strconv.Atoi(parts[1])

		if command == "ADD" {
			title := strings.Join(parts[2:], " ")
			library[isbn] = title
		} else if command == "REMOVE" {
			delete(library, isbn)
		}
	}

	var books []Book
	for isbn, title := range library {
		books = append(books, Book{ISBN: isbn, Title: title})
	}

	sort.Slice(books, func(i, j int) bool {
		if books[i].Title == books[j].Title {
			return books[i].ISBN < books[j].ISBN
		}
		return books[i].Title < books[j].Title
	})

	for _, book := range books {
		fmt.Println(book.ISBN)
	}

}
