package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Book struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Borrowed bool   `json:"borrowed"`
}

type Library struct {
	Books []Book
}

type Server struct {
	library *Library
	port    string
}

func NewServer(port string) *Server {
	library := &Library{Books: []Book{}}
	return &Server{library, port}
}

func (s *Server) Start() {
	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			s.handleAddBook(w, r)
		case http.MethodPut:
			s.handleBorrowBook(w, r)
		case http.MethodDelete:
			s.handleDeleteBook(w, r)
		case http.MethodGet:
			s.handleGetBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":"+s.port, nil)
}

func (s *Server) handleAddBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("author")

	title = strings.ToLower(title)
	author = strings.ToLower(author)

	if title == "" || author == "" {
		http.Error(w, `{"Result": "", "Error": "title or author cannot be empty"}`, http.StatusBadRequest)
		return
	}

	for _, book := range s.library.Books {
		if book.Title == title && book.Author == author {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"Result": "this book is already in the library", "Error": ""}`)
			return
		}
	}

	newBook := Book{
		Title:    title,
		Author:   author,
		Borrowed: false,
	}
	s.library.Books = append(s.library.Books, newBook)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"Result": "added book %s by %s", "Error": ""}`, title, author)
}

type RequestBodyBorrow struct {
	Borrow bool `json:"borrow"`
}

func (s *Server) handleBorrowBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("author")

	title = strings.ToLower(title)
	author = strings.ToLower(author)

	if title == "" || author == "" {
		http.Error(w, `{"Result": "", "Error": "title or author cannot be empty"}`, http.StatusBadRequest)
		return
	}

	var requestData RequestBodyBorrow
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	borrow := requestData.Borrow

	for i, book := range s.library.Books {
		if book.Title == title && book.Author == author {
			if borrow {
				if book.Borrowed {
					http.Error(w, `{"Result": "", "Error": "this book is already borrowed"}`, http.StatusBadRequest)
					return
				}
				s.library.Books[i].Borrowed = true
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{"Result": "you have borrowed this book successfully", "Error": ""}`)
				return
			} else {
				if !book.Borrowed {
					http.Error(w, `{"Result": "", "Error": "this book is already in the library"}`, http.StatusBadRequest)
					return
				}
				s.library.Books[i].Borrowed = false
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, `{"Result": "thank you for returning this book", "Error": ""}`)
				return
			}
		}
	}

	http.Error(w, `{"Result": "", "Error": "this book does not exist"}`, http.StatusBadRequest)
}

func (s *Server) handleDeleteBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("author")

	title = strings.ToLower(title)
	author = strings.ToLower(author)

	if title == "" || author == "" {
		http.Error(w, `{"Result": "", "Error": "title or author cannot be empty"}`, http.StatusBadRequest)
		return
	}

	for i, book := range s.library.Books {
		if book.Title == title && book.Author == author {
			s.library.Books = append(s.library.Books[:i], s.library.Books[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, `{"Result": "successfully deleted", "Error": ""}`)
			return
		}
	}

	http.Error(w, `{"Result": "", "Error": "this book does not exist"}`, http.StatusBadRequest)
}

func (s *Server) handleGetBook(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	author := r.FormValue("author")

	title = strings.ToLower(title)
	author = strings.ToLower(author)

	if title == "" || author == "" {
		http.Error(w, `{"Result": "", "Error": "title or author cannot be empty"}`, http.StatusBadRequest)
		return
	}

	for _, book := range s.library.Books {
		if book.Title == title && book.Author == author {
			if book.Borrowed {
				http.Error(w, `{"Result": "", "Error": "this book is borrowed"}`, http.StatusBadRequest)
				return
			}
			response, _ := json.Marshal(book)
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
			return
		}
	}

	http.Error(w, `{"Result": "", "Error": "this book does not exist"}`, http.StatusBadRequest)
}