package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(f), "\n")
	nums := make([]int, len(lines))
	m := make(map[int]bool)
	for i := range lines {
		nums[i], _ = strconv.Atoi(lines[i])
		m[nums[i]] = true
	}
	sort.Ints(nums)
	// nums = append(nums, nums[len(nums)-1]+3)
	arr := make([]int, nums[len(nums)-1])
	if m[1] {
		arr[0] = 1
	}
	if m[2] {
		arr[1] = arr[0] + 1
	}
	if m[3] {
		arr[2] = arr[1] + arr[0] + 1
	}
	for i := 3; i < len(arr); i++ {
		if m[i+1] {
			arr[i] = arr[i-1] + arr[i-2] + arr[i-3]
		}
	}
	fmt.Println(arr[len(arr)-1])
}
