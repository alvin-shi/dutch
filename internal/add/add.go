package add

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/alvin-shi/dutch/internal/utils"
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
	fp, error := utils.CreateAndOrOpen()
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	defer fp.Close()

	_, err := fp.Write(fmt.Appendln([]byte(strings.Join(args, " "))))
	if err != nil {
		fmt.Println(err.Error())
	}
}
