package print

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/alvin-shi/dutch/internal/utils"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func Print() {
	filename, error := utils.FindSaveFilename()
	if error != nil {
		fmt.Println(error.Error())
		return
	} else if filename == "" {
		fmt.Println("no save file found")
		return
	}

	filepath := fmt.Sprintf("/tmp/%v", filename)
	contents, error := os.ReadFile(filepath)
	if error != nil {
		fmt.Println(error.Error())
		return
	}

	// user -> item -> cost
	breakdown := make(map[string]map[string]float64)
	items := make([]any, 0)
	for line := range strings.SplitSeq(string(contents), "\n") {
		if line == "" {
			continue
		}
		splitLine := strings.Split(line, " ")
		item := splitLine[0]
		cost, error := strconv.ParseFloat(splitLine[1], 64)
		if error != nil {
			fmt.Println(error.Error())
			return
		}
		contributors := splitLine[2:]
		costPerContributor := cost / float64(len(contributors))

		items = append(items, item)
		// Should try fix this so we don't have to assume it is unique
		for _, contributor := range contributors {
			if savedContributor, found := breakdown[contributor]; !found {
				newContributor := make(map[string]float64)
				newContributor[item] = costPerContributor
				breakdown[contributor] = newContributor
			} else {
				savedContributor[item] = costPerContributor
			}
		}
	}

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	people := slices.Collect(maps.Keys(breakdown))
	headers := make([]any, len(people)+1)
	headers[0] = "Items"
	for index, person := range people {
		headers[index+1] = person

	}
	tbl := table.New(headers...)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for _, item := range items {
		row := make([]any, len(people)+1)
		row[0] = item
		for i, person := range people {
			if amount, found := breakdown[person][item.(string)]; !found {
				row[i+1] = 0.0
			} else {
				row[i+1] = amount
			}
		}
		tbl.AddRow(row...)
	}

	total := make([]any, len(people)+1)
	total[0] = "Total"
	for i, person := range people {
		amounts := maps.Values(breakdown[person])
		totalAmount := 0.0

		for amount := range amounts {
			totalAmount += amount
		}
		total[i+1] = totalAmount
	}
	tbl.AddRow(total...)
	tbl.Print()
}
