package internal

import (
	"slices"
)

// String functions taken from https://github.com/nektro/go-util/blob/master/arrays/stringsu/strings.go

func Filter(stack []string, cb func(string) bool) []string {
	result := []string{}
	for _, item := range stack {
		if cb(item) {
			result = append(result, item)
		}
	}
	return result
}

func Remove(stack []string, search ...string) []string {
	return Filter(stack, func(s string) bool {
		return !slices.Contains(search, s)
	})
}
