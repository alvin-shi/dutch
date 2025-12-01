package main

import (
	"errors"
	"fmt"
	"strconv"
)

type breakdown struct {
	Item         string
	Cost         float64
	Contributors []string
}

func add(args []string) {
	data, error := parseArgs(args)
	if error != nil {
		fmt.Println(error)
	}

	portion := data.Cost / float64(len(data.Contributors))
	fmt.Println(data.Item, "cost per person:", portion)
}

func parseArgs(args []string) (breakdown, error) {
	if len(args) < 3 {
		return breakdown{}, errors.New("Add needs item, cost and at least one contributor")
	}

	item := args[0]
	cost, error := strconv.ParseFloat(args[1], 64)
	if error != nil {
		errorString := fmt.Sprintf("error converting cost to float %v", args[1])
		return breakdown{}, errors.New(errorString)
	}
	contributors := args[2:]

	return breakdown{
		Item:         item,
		Cost:         cost,
		Contributors: contributors,
	}, nil
}
