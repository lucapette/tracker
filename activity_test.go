package main

import (
	"reflect"
	"testing"
)

func TestNewActivity(t *testing.T) {
	tests := []struct {
		name     string
		expected Activity
	}{
		{name: "iTerm2", expected: Activity{Name: "iTerm2", Category: Development}},
		{name: "https://twitter.com", expected: Activity{Name: "twitter.com", Category: Social}},
		{name: "airmail", expected: Activity{Name: "airmail", Category: Communication}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := NewActivity(test.name)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("Expected %v to equal %v", actual, test.expected)
			}
		})
	}
}