package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")
	ans := 0
	for _, line := range lines {
		n, _ := strconv.Atoi(calculate(parse(line)))
		ans += n
	}
	fmt.Println(ans)
}

func parse(s string) []string {
	arr := make([]string, 0)
	start := 0
	for i := 0; i < len(s); i++ {
		r := s[i]
		if r == ' ' {
			continue
		}
		if r == '(' {
			start = i + 1
			count := 1
			for j := start; j < len(s); j++ {
				if s[j] == ')' && count == 1 {
					i = j
					break
				}
				if s[j] == '(' {
					count++
				} else if s[j] == ')' {
					count--
				}
			}
			arr = append(arr, calculate(parse(s[start:i])))
			continue
		}
		arr = append(arr, string(r))
	}
	return arr
}

func calculate(s []string) string {
	pre := 0
	op := "+"
	mult := make([]int, 0)
	for _, cur := range s {
		switch cur {
		case "+", "*":
			op = cur
		default:
			n, _ := strconv.Atoi(cur)
			if op == "+" {
				pre = pre + n
			} else {
				mult = append(mult, pre)
				pre = n
			}
		}
	}
	mult = append(mult, pre)
	ans := 1
	for _, cur := range mult {
		ans *= cur
	}
	return fmt.Sprintf("%d", ans)
}
