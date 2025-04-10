package utils

import (
	"slices"
	"strings"
)

func PathParser(path string) []string {
	parse := strings.Split(path, "/")
	for i := range parse {
		if strings.TrimSpace(parse[i]) == "" {
			parse = slices.Delete(parse, i, i+1)
		}
	}
	return strings.Split(path, "/")
}
