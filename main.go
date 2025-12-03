package main

import (
	"fmt"
	"os"

	"github.com/alvin-shi/dutch/internal/add"
)

func main() {
	args := tail(os.Args)
	if len(args) == 0 {
		fmt.Println("No arguments provided")
		return
	}

	switch args[0] {
	case "add":
		add.Add(tail(args))
	case "print":
		print()
	}
}

func print() {
	fmt.Println("Print!")
}
func tail[T any](slice []T) []T {
	return slice[1:]
}
