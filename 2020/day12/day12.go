package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func main() {
	f, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")

	facing := 1
	ship := coordinate{0, 0}
	waypoint := coordinate{10, 1}
	dirMap := map[byte]int{'R': 1, 'L': -1, 'N': 1, 'S': -1, 'W': -1, 'E': 1}
	for _, line := range lines {
		dir := line[0]
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch dir {
		case 'N', 'S':
			waypoint.y += dist * dirMap[dir]
		case 'W', 'E':
			waypoint.x += dist * dirMap[dir]
		case 'R', 'L':
			pre := facing + 4
			facing = (pre + dist/90*dirMap[dir]) % 4
			for i := 0; i < (pre-facing)%4; i++ {
				waypoint.x, waypoint.y = -waypoint.y, waypoint.x
			}
		case 'F':
			ship.x += waypoint.x * dist
			ship.y += waypoint.y * dist
		}
	}

	fmt.Println(math.Abs(float64(ship.x)) + math.Abs(float64(ship.y)))
}
