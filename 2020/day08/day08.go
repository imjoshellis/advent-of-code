package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("./input.txt")
	l := len(strings.Split(string(f), "\n"))
	for i := 0; i < l; i++ {
		lines := strings.Split(string(f), "\n")
		seen := make(map[int]bool)
		next := 0
		acc := 0
		swap := strings.Split(lines[i], " ")
		if swap[0] == "nop" {
			lines[i] = "jmp " + swap[1]
		} else if swap[0] == "jmp" {
			lines[i] = "nop " + swap[1]
		}
		for !seen[next] {
			if next == len(lines) {
				fmt.Println("found:", acc)
				return
			}
			if next > len(lines) {
				continue
			}
			cur := strings.Split(lines[next], " ")
			op := cur[0]
			val, _ := strconv.Atoi(cur[1])
			seen[next] = true
			switch op {
			case "acc":
				acc += val
				next++
			case "nop":
				next++
			case "jmp":
				next += val
			}
		}
	}
}
