package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	file, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	ans := 1
	for _, slope := range slopes {
		ans *= count(lines, slope)
	}

	fmt.Println(ans)
}
