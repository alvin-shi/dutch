package add

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Add(args []string) {
	if ok, error := validate(args); !ok {
		fmt.Println(error.Error())
		return
	}

	store(args)
}

func validate(args []string) (bool, error) {
	if len(args) < 3 {
		return false, errors.New("Add needs item, cost and at least one contributor")
	}

	_, error := strconv.ParseFloat(args[1], 64)
	if error != nil {
		errorString := fmt.Sprintf("error converting cost to float %v", args[1])
		return false, errors.New(errorString)
	}

	return true, nil
}

func store(args []string) {
	fp, error := createAndOrOpen()
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	defer fp.Close()

	fmt.Println("writing to file:", fp.Name())
	_, err := fp.Write(fmt.Appendln([]byte(strings.Join(args, " "))))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func createAndOrOpen() (*os.File, error) {
	entries, error := os.ReadDir(os.TempDir())
	if error != nil {
		return nil, error
	}

	index := slices.IndexFunc(entries, func(entry os.DirEntry) bool {
		return strings.HasPrefix(entry.Name(), "dutch_state")
	})

	if index < 0 {
		fmt.Println("creating new file")
		return os.CreateTemp("", "dutch_state")
	}

	filepath := fmt.Sprintf("/tmp/%v", entries[index].Name())
	return os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
}
