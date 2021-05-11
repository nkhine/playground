package main

import "fmt"

func main() {
	type text struct {
		title string
		words int
	}
	type book struct {
		text
		isbn string
	}

	moby := book{
		text: text{title: "moby dick", words: 234567},
		isbn: "X12345",
	}
	fmt.Printf("%s has %s words (isbn: %s)\n",
		moby.text.title, moby.text.words, moby.isbn)
}
