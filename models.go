package main

import (
	"fmt"
)

type Book struct {
	ID       int
	Author   string
	Year     int
	Title    string
	IsIssued bool
	ReaderID *int
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("%s (%s, %d), статус %s", b.Title, b.Author, b.Year, status)
}

// IssueBook выдает книгу читателю
func (b *Book) IssueBook(reader *Reader) {
	if !reader.IsActive {
		fmt.Printf("Читатель %s %s не активен и не может получить книгу.", reader.FirstName, reader.LastName)
		return
	}
	if b.IsIssued {
		fmt.Printf("Книга %s уже кому-то выдана\n", b.Title)
		return
	}
	b.IsIssued = true
	b.ReaderID = &reader.ID

	fmt.Printf("Книга %s была выдана\n", b.Title)
}

func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга %s и так в библиотеке", b.Title)
		return
	}
	b.IsIssued = false
	b.ReaderID = nil
	fmt.Printf("Книга %s возвращена в библиотеку\n", b.Title)
}

func (r *Reader) Deactivate() {
	r.IsActive = false
}

func (r *Reader) Activate() {
	r.IsActive = true
}
