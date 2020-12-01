package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nums := []int{}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, n)
	}

	for i, n1 := range nums {
		m := map[int][2]int{}
		for _, n2 := range nums[i+1:] {
			if m[n2][0] != 0 {
				println(m[n2][0], m[n2][1], n2, m[n2][0]*m[n2][1]*n2)
				return
			}
			m[2020-n1-n2] = [2]int{n1, n2}
		}
	}
}
