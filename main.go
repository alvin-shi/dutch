package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := tail(os.Args)
	if len(args) == 0 {
		fmt.Println("No arguments provided")
		return
	}

	switch args[0] {
	case "add":
		add(tail(args))
	case "print":
		print()
	}
}

func add(args []string) {
	if len(args) < 3 {
		fmt.Println("Add needs item, cost and at least one contributor")
		return
	}

	item := args[0]
	cost, error := strconv.ParseFloat(args[1], 64)
	if error != nil {
		fmt.Println("error converting cost to float:", args[1])
		return
	}
	contributors := args[2:]

	portion := cost / float64(len(contributors))
	fmt.Println(item, "cost per person:", portion)
}

func print() {
	fmt.Println("Print!")
}
func tail[T any](slice []T) []T {
	return slice[1:]
}
