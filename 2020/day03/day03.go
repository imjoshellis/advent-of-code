package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	count := make([]int, 5)
	idx := 0
	for k := 1; k < 8; k += 2 {
		x := 0
		for _, line := range lines {
			x = x % len(line)

			if line[x] == '#' {
				count[idx]++
			}
			x += k
		}
		idx++
	}

	x := 0
	for i, line := range lines {
		if i%2 != 0 {
			continue
		}
		x = x % len(line)

		if line[x] == '#' {
			count[idx]++
		}
		x += 1
	}
	ans := 1
	for _, c := range count {
		ans *= c
	}

	fmt.Println(ans)
}
