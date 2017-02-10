package main

import (
	"log"
	"strings"
)

type Category struct {
	Name  string
	Score int
}

var Categories map[string]Category

func registerCategory(name string, score int) Category {
	category := Category{
		Name:  name,
		Score: score,
	}

	Categories[name] = category

	return category
}

var Activities map[string]Activity

func init() {
	Categories = make(map[string]Category, 6)
	registerCategory("Development", 1)
	registerCategory("General", 1)
	registerCategory("Writing", 1)

	registerCategory("Communication", 0)
	registerCategory("Uncategorized", 0)

	registerCategory("Entertainment", -1)
	registerCategory("Social", -1)

	asset, err := Asset("categories.csv")
	if err != nil {
		log.Panic(err)
	}

	lines := strings.Split(string(asset), "\n")

	Activities = make(map[string]Activity, len(lines)-1)
	for _, line := range lines {
		cols := strings.Split(line, ",")
		if len(cols[0]) == 0 {
			continue
		}

		frontApp := cols[0]
		category := Categories[cols[1]]

		Activities[frontApp] = Activity{Name: frontApp, Category: category}
	}
}
