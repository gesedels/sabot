package main

import (
	"fmt"
	"os"

	comm "github.com/gesedels/sabot/sabot/comms"
	"github.com/gesedels/sabot/sabot/items/book"
)

func try(err error) {
	if err != nil {
		fmt.Printf("Error: %s.\n", err)
		os.Exit(1)
	}
}

func main() {
	book, err := book.NewPath("./sabot.db")
	try(err)
	try(comm.Run(os.Stdout, book, os.Args[1:]))
	try(book.Close())
}
