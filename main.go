package main

import (
	"fmt"
)

type Library struct {
	Books   map[int]*Book
	Readers map[int]*Reader
}

func main() {
	fmt.Printf("Простая библиотека запущена\n")

	book1 := Book{
		ID:     1,
		Title:  "Война и мир",
		Author: "Лев Толстой",
		Year:   1869,
	}

	user1 := Reader{
		ID:        1,
		FirstName: "Петр",
		LastName:  "Петров",
		IsActive:  true,
	}

	book1.IssueBook(&user1)
	fmt.Println(book1)

	reader2 := Reader{
		ID:        2,
		FirstName: "Sergey",
		LastName:  "Meniaylo",
		IsActive:  true,
	}
	book1.IssueBook(&reader2)

	user1.Deactivate()
	fmt.Println(user1)
	fmt.Println("---")

	lib := Library{
		Books:   make(map[int]*Book),
		Readers: make(map[int]*Reader),
	}

	lib.Books[book1.ID] = &book1
	lib.Readers[user1.ID] = &user1
	lib.Readers[reader2.ID] = &reader2

	book1.IssueBook(&user1)
}
