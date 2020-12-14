package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")
	fmt.Println(part2(lines))
}

func part1(lines []string) int64 {
	mask := ""
	mem := make([]int64, 65469+1)
	r, _ := regexp.Compile(`\d+`)
	for _, line := range lines {
		split := strings.Split(line, " = ")
		if split[0] == "mask" {
			mask = split[1]
		} else {
			n, _ := strconv.Atoi(split[1])
			b := []byte(fmt.Sprintf("%036b", n))
			for i := range mask {
				if mask[i] != 'X' {
					b[i] = mask[i]
				}
			}
			k, _ := strconv.Atoi(r.FindString(split[0]))
			v, _ := strconv.ParseInt(string(b), 2, 37)
			mem[k] = v
		}
	}
	var sum int64
	for _, n := range mem {
		sum += n
	}
	return sum
}

func part2(lines []string) int {
	mask := ""
	mem := make(map[string]int)
	r, _ := regexp.Compile(`\d+`)
	for _, line := range lines {
		split := strings.Split(line, " = ")
		if split[0] == "mask" {
			mask = split[1]
		} else {
			v, _ := strconv.Atoi(split[1])
			k, _ := strconv.Atoi(r.FindString(split[0]))
			b := []byte(fmt.Sprintf("%036b", k))
			st := [][]byte{b}
			for i := range mask {
				l := len(st)
				for j := 0; j < l; j++ {
					cur := st[0]
					// n, _ := strconv.ParseInt(string(cur), 2, 37)
					// fmt.Println(i, string(cur), n)
					st = st[1:]
					if mask[i] == 'X' {
						cur1 := make([]byte, len(cur))
						copy(cur1, cur)
						cur2 := make([]byte, len(cur))
						copy(cur2, cur)
						cur1[i] = '1'
						cur2[i] = '0'
						st = append(st, [][]byte{cur1, cur2}...)
					} else if mask[i] == '1' {
						cur[i] = mask[i]
						st = append(st, cur)
					} else {
						st = append(st, cur)
					}
				}
			}

			for _, address := range st {
				mem[string(address)] = v
			}
		}
	}
	var sum int
	for _, n := range mem {
		sum += n
	}
	return sum
}
