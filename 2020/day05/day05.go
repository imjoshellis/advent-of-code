package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("./input.txt")
	seats := strings.Split(string(f), "\n")
	tickets := [][]int{}
	for _, s := range seats[1:] {
		tickets = append(tickets, processSeat(s))
	}

	sort.Slice(tickets, func(i, j int) bool {
		return tickets[i][2] > tickets[j][2]
	})

	for i := range tickets {
		if tickets[i][2]-2 == tickets[i+1][2] {
			fmt.Print(tickets[i], tickets[i+1])
			return
		}
	}
}

func processSeat(s string) []int {
	rs, cs := s[:7], s[7:]
	minr, maxr := 0, 127
	minc, maxc := 0, 7

	for _, p := range rs {
		if p == 'F' {
			maxr = minr + (maxr-minr)/2
		} else {
			minr = minr + (maxr-minr)/2
		}
	}

	for _, p := range cs {
		if p == 'L' {
			maxc = minc + (maxc-minc)/2
		} else {
			minc = minc + (maxc-minc)/2
		}
	}
	r := maxr
	c := maxc
	id := r*8 + c
	return []int{r, c, id}
}
