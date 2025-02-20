package main

import (
	"fmt"
	"hangman/hg"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a dictionary file")
		fmt.Println("Usage: ./hangman words.txt")
		return
	}
	hg.StartGame(os.Args[1])
}
