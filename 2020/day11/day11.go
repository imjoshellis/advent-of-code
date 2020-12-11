package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(f), "\n")
	arr := make([][]string, len(lines))
	for i := range arr {
		arr[i] = make([]string, len(lines[i]))
		for j := range arr[i] {
			arr[i][j] = string(lines[i][j])
		}
	}

	changed := true
	for changed {
		changed = false
		next := make([][]string, len(arr))
		for i := range arr {
			next[i] = make([]string, len(arr[i]))
			for j := range next[i] {
				next[i][j] = arr[i][j]
			}
		}
		for i := range arr {
			for j := range arr[i] {
				if arr[i][j] == "L" && countOccupied(arr, i, j) == 0 {
					next[i][j] = "#"
					changed = true
				}
				if arr[i][j] == "#" && countOccupied(arr, i, j) >= 5 {
					next[i][j] = "L"
					changed = true
				}
			}
		}
		copy(arr, next)
	}

	count := 0
	for i := range arr {
		for j := range arr[i] {
			if arr[i][j] == "#" {
				count++
			}
		}
	}

	fmt.Println(count)
}

func countOccupied(arr [][]string, i, j int) int {
	dirs := [][]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}, {0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	count := 0
mainloop:
	for _, dir := range dirs {
		di := i + dir[0]
		dj := j + dir[1]
		if outOfGrid(arr, di, dj) {
			continue mainloop
		}
		for arr[di][dj] == "." {
			fmt.Println(arr[di][dj], di, dj, i, j)
			di += dir[0]
			dj += dir[1]
			if outOfGrid(arr, di, dj) {
				continue mainloop
			}
		}
		if arr[di][dj] == "#" {
			count++
		}
	}
	return count
}

func outOfGrid(arr [][]string, i, j int) bool {
	return i < 0 || i >= len(arr) || j < 0 || j >= len(arr[0])
}
