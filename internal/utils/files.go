package utils

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

const FILE_PREFIX = "dutch_state"

func FindSaveFilename() (string, error) {
	entries, error := os.ReadDir(os.TempDir())
	if error != nil {
		return "", error
	}

	index := slices.IndexFunc(entries, func(entry os.DirEntry) bool {
		return strings.HasPrefix(entry.Name(), FILE_PREFIX)
	})

	if index < 0 {
		return "", nil
	}

	return entries[index].Name(), nil
}

func CreateAndOrOpen() (*os.File, error) {
	if filename, error := FindSaveFilename(); error != nil {
		return nil, error
	} else if filename == "" {
		return os.CreateTemp("", FILE_PREFIX)
	} else {
		filepath := fmt.Sprintf("/tmp/%v", filename)
		return os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}
}
