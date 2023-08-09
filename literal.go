package main

import "github.com/caarlos0/markscribe/literal"

func literalClubCurrentlyReading(count int) []literal.Book {
	books, err := literal.CurrentlyReading()
	if err != nil {
		panic(err)
	}
	if len(books) > count {
		return books[:count]
	}
	return books
}
