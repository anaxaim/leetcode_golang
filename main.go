package main

import "fmt"

type BookStore interface {
	AddBook(book string)
}

type MyBookStore struct {
	books []string
}

func (store *MyBookStore) AddBook(book string) {
	store.books = append(store.books, book)
}

func main() {
	var sl = []string{"AA", "ABC"}
	var booker BookStore = &MyBookStore{sl}
	booker.AddBook("CCCC")

	fmt.Println(booker)
}
