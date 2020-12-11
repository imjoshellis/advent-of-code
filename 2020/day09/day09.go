package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(f), "\n")
	nums := make([]int, len(lines))
	for i := range lines {
		nums[i], _ = strconv.Atoi(lines[i])
	}

	// part 1
	// MainLoop:
	// 	for i := 0; i < len(lines)-25; i++ {
	// 		target, _ := strconv.Atoi(lines[i+25])
	// 		for j := i; j < i+25; j++ {
	// 			for k := j + 1; k < i+25; k++ {
	// 				n1, _ := strconv.Atoi(lines[j])
	// 				n2, _ := strconv.Atoi(lines[k])
	// 				if n1+n2 == target {
	// 					continue MainLoop
	// 				}
	// 			}
	// 		}
	// 		fmt.Println(i, target)
	// 		return
	// 	}

	for start := 0; start < len(nums)-1; start++ {
		for end := start + 1; sum(nums[start:end+1]) <= 1398413738; end++ {
			if a := nums[start : end+1]; sum(a) == 1398413738 {
				fmt.Println(start, end, min(a)+max(a))
				return
			}
		}
	}
	fmt.Println("not found")
}

func sum(arr []int) int {
	result := 0
	for _, v := range arr {
		result += v
	}
	return result
}

func min(arr []int) int {
	res := math.MaxInt32
	for _, v := range arr {
		if v < res {
			res = v
		}
	}
	return res
}

func max(arr []int) int {
	res := arr[0]
	for _, v := range arr {
		if v > res {
			res = v
		}
	}
	return res
}
