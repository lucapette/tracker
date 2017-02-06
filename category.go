package main

type Category struct {
	Name  string
	Score int
}

var Communication = Category{Name: "Communication", Score: 0}
var Development = Category{Name: "Development", Score: 1}
var Social = Category{Name: "Social", Score: -1}
var Uncategorized = Category{Name: "Uncategorized", Score: 0}

var categories map[string]Category

func init() {
	categories = map[string]Category{
		"iTerm2":            Development,
		"github.com":        Development,
		"stackoverflow.com": Development,
		"Dash":              Development,
		"localhost":         Development,
		"twitter.com":       Social,
		"reddit.com":        Social,
		"medium.com":        Social,
		"linkedin.com":      Social,
		"airmail":           Communication,
		"slack.com":         Communication,
	}
}
