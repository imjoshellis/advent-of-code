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
	p := coordinate{0, 0}
	w := coordinate{10, 1}
	m := map[byte]int{'R': 1, 'L': -1, 'N': 1, 'S': -1, 'W': -1, 'E': 1}
	for _, line := range lines {
		dir := line[0]
		n, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		switch dir {
		case 'F':
			p.x += w.x * n
			p.y += w.y * n
		case 'N', 'S':
			w.y += n * m[dir]
		case 'W', 'E':
			w.x += n * m[dir]
		case 'R', 'L':
			pre := facing + 4
			facing = (pre + n/90*m[dir]) % 4
			for i := 0; i < (pre-facing)%4; i++ {
				w.x, w.y = -w.y, w.x
			}
		}
	}

	fmt.Println(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}
