package tracker

import (
	_ "embed"
	"strings"
)

//go:embed categories.csv
var asset string

type Category struct {
	Name  string
	Score int
}

var categories map[string]Category

func registerCategory(name string, score int) Category {
	category := Category{
		Name:  name,
		Score: score,
	}

	categories[name] = category

	return category
}

var activities map[string]Activity

func init() {
	categories = make(map[string]Category, 6)
	registerCategory("Development", 1)
	registerCategory("General", 1)
	registerCategory("Writing", 1)

	registerCategory("Communication", 0)
	registerCategory("Uncategorised", 0)

	registerCategory("Entertainment", -1)
	registerCategory("Social", -1)

	lines := strings.Split(asset, "\n")

	activities = make(map[string]Activity, len(lines)-1)
	for _, line := range lines {
		cols := strings.Split(line, ",")
		if len(cols[0]) == 0 {
			continue
		}

		frontApp := cols[0]
		category := categories[cols[1]]

		activities[frontApp] = Activity{Name: frontApp, Category: category}
	}
}
