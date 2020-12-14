package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestProcess(t *testing.T) {
	for i := 0; i < 36; i++ {
		mask := []byte("mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
		mask[i+7] = '1'
		mem := "mem[1] = 0"
		lines := []string{string(mask), mem}
		expect := []byte(fmt.Sprintf("%036b", 0))
		expect[i] = '1'
		n, _ := strconv.ParseInt(string(expect), 2, 36)
		got := part1(lines)
		if got != n {
			t.Fatalf("Expected %v, got %v", n, got)
		}

	}
}
