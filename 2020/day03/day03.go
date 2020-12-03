package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func count(m []string, s []int) int {
	c := 0
	for x, y := 0, 0; y < len(m); x, y = x+s[0], y+s[1] {
		x = x % len(m[0])
		if m[y][x] == '#' {
			c++
		}
	}
	return c
}

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(file), "\n")

	ans := 1
	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, slope := range slopes {
		ans *= count(lines, slope)
	}

	fmt.Println(ans)
}
