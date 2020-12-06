package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("./input.txt")
	gs := strings.Split(string(f), "\n")

	count := 0
	for i := 0; i < len(gs); i++ {
		ltrs := make([]int, 26)
		l := 0
		for len(gs[i]) > 0 {
			for _, ans := range gs[i] {
				ltrs[ans-'a']++
			}
			l++
			i++
		}
		for _, ltr := range ltrs {
			if ltr == l {
				count++
			}
		}
	}
	fmt.Println(count)
}
