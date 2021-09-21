package main

import (
	"fmt"
	"strings"
)

func main() {
	result := isAdjacent("122")
	fmt.Println(result)
}

func isAdjacent(n string) bool {
	var adjacent bool
	splited := strings.Split(n, "")
	for i := range splited {
		if i != (len(splited) - 1) {
			if splited[i] == splited[i+1] {
				adjacent = true
				return adjacent
			}
		} else {
			adjacent = false
		}
	}
	return adjacent
}
