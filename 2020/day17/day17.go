package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type (
	Grid     [][][][]int
	Position struct {
		x int
		y int
		z int
		w int
	}
)

func main() {
	f, _ := ioutil.ReadFile("input.txt")
	split := strings.Split(string(f), "\n")
	grid := initialize(split)
	steps := 5
	for i := 0; i <= steps; i++ {
		grid.update()
	}
	count := 0
	for w := range grid {
		for z := range grid {
			for y := range grid[z] {
				for x := range grid[z][y] {
					count += grid[w][z][y][x]
				}
			}
		}
	}
	fmt.Println("count:", count)
}

func initialize(input []string) Grid {
	inWidth := len(input)
	size := inWidth + 11
	grid := make(Grid, size)
	offset := (size - inWidth) / 2
	for w := range grid {
		grid[w] = make([][][]int, size)
		for z := range grid[w] {
			grid[w][z] = make([][]int, size)
			for y := range grid[w][z] {
				grid[w][z][y] = make([]int, size)
				for x := range grid[w][z][y] {
					grid[w][z][y][x] = 0
					if z == 6 && w == 6 && x >= offset && x < offset+inWidth && y >= offset && y < offset+inWidth && input[y-offset][x-offset] == '#' {
						grid[w][z][y][x] = 1
					}
				}
			}
		}
	}
	return grid
}

func (g *Grid) update() {
	old := make(Grid, len(*g))
	for w := range *g {
		old[w] = make([][][]int, len(*g))
		for z := range (*g)[w] {
			old[w][z] = make([][]int, len(*g))
			for y := range (*g)[w][z] {
				old[w][z][y] = make([]int, len(*g))
				for x := range (*g)[w][z][y] {
					old[w][z][y][x] = (*g)[w][z][y][x]
				}
			}
		}
	}
	for w := range old {
		for z := range old[w] {
			for y := range old[w][z] {
				for x := range old[w][z][y] {
					(*g)[w][z][y][x] = nextCell(old, Position{x, y, z, w})
				}
			}
		}
	}
}

func nextCell(g Grid, p Position) int {
	// If a cube is active and exactly 2 or 3 of its neighbors are also active, the cube remains active. Otherwise, the cube becomes inactive.
	// If a cube is inactive but exactly 3 of its neighbors are active, the cube becomes active. Otherwise, the cube remains inactive.
	n := countActiveNeighbors(g, p)
	cur := g[p.w][p.z][p.y][p.x]
	if cur == 1 {
		if n == 2 || n == 3 {
			return 1
		}
		return 0
	}
	if n == 3 {
		return 1
	}
	return 0
}

func countActiveNeighbors(g Grid, p Position) int {
	count := 0
	for _, d := range dirs {
		dp := Position{p.x + d.x, p.y + d.y, p.z + d.z, p.w + d.w}
		if inGrid(g, dp) && g[dp.w][dp.z][dp.y][dp.x] == 1 {
			count++
		}
	}
	return count
}

func inGrid(g Grid, p Position) bool {
	if p.x < 0 || p.x >= len(g) {
		return false
	}
	if p.y < 0 || p.y >= len(g) {
		return false
	}
	if p.z < 0 || p.z >= len(g) {
		return false
	}
	if p.w < 0 || p.w >= len(g) {
		return false
	}
	return true
}

func makeDirs() []Position {
	dirs := make([]Position, 0)
	for w := -1; w <= 1; w++ {
		for z := -1; z <= 1; z++ {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if w == 0 && z == 0 && y == 0 && x == 0 {
						continue
					}
					dirs = append(dirs, Position{x, y, z, w})
				}
			}
		}
	}
	return dirs
}

var dirs = makeDirs()
