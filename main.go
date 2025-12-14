package main

import (
	"fmt"
	"os"

	"github.com/alvin-shi/dutch/internal/add"
	"github.com/alvin-shi/dutch/internal/print"
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
		print.Print()
	default:
		fmt.Println("command not found")
	}
}

func tail[T any](slice []T) []T {
	return slice[1:]
}
