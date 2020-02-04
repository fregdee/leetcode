package main

import (
	"fmt"
	"strings"
)

func main() {
	moves := "LL"
	fmt.Println(judgeCircle(moves))
}

func judgeCircle(moves string) bool {
	moveSign := strings.Split(moves, "")
	current := make(map[string]int)
	current["x"] = 0
	current["y"] = 0

	for _, sign := range moveSign {
		switch sign {
		case "R":
			current["x"] += 1
		case "L":
			current["x"] -= 1
		case "U":
			current["y"] += 1
		case "D":
			current["y"] -= 1
		}
	}

	fmt.Println(current)

	var result bool

	if current["x"] == 0 && current["y"] == 0 {
		result = true
	} else {
		result = false
	}

	return result
}
