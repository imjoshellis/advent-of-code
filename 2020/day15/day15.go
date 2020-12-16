package main

import (
	"fmt"
)

func main() {
	initial := []int{0, 12, 6, 13, 20, 1, 17}
	target := 30000000
	fmt.Println(getNth(initial, target))
}

func getNth(initial []int, target int) int {
	l := len(initial) - 1
	next := initial[l]

	m := make(map[int]int)
	for i, n := range initial {
		m[n] = i
	}

	for i := l; i < target-1; i++ {
		if prev, ok := m[next]; ok {
			next, m[next] = i-prev, i
		} else {
			next, m[next] = 0, i
		}
	}

	return next
}
