package main

import (
	"fmt"
)

func main() {
	ns := []int{0, 12, 6, 13, 20, 1}
	m := make(map[int]int)
	for i, n := range ns {
		m[n] = i
	}

	nxt := 17
	for i := 6; i < 30000000-1; i++ {
		if prev, ok := m[nxt]; ok {
			nxt, m[nxt] = i-prev, i
		} else {
			nxt, m[nxt] = 0, i
		}
	}

	fmt.Println(nxt)
}
