package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		i := 0
		one := ""
		two := ""
		for line[i] != '-' {
			one += string(line[i])
			i++
		}
		i++
		for line[i] != ' ' {
			two += string(line[i])
			i++
		}
		i++
		ltr := line[i]

		oneIdx, _ := strconv.Atoi(one)
		twoIdx, _ := strconv.Atoi(two)
		oneIdx = i + oneIdx + 2
		twoIdx = i + twoIdx + 2
		valid := false
		if line[oneIdx] == ltr || line[twoIdx] == ltr {
			if line[oneIdx] != line[twoIdx] {
				valid = true
				count++
			}
		}
		fmt.Println(line)
		fmt.Println(string(ltr), string(line[oneIdx]), string(line[twoIdx]), oneIdx, twoIdx, valid)
		fmt.Println("---")
	}

	fmt.Println(count)
}

/* part 1
func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		i := 0
		min := ""
		max := ""
		for line[i] != '-' {
			min += string(line[i])
			i++
		}
		i++
		for line[i] != ' ' {
			max += string(line[i])
			i++
		}
		i++
		ltr := line[i]
		i += 2

		passCount := 0
		for ; i < len(line); i++ {
			if line[i] == ltr {
				passCount++
			}
		}

		minInt, _ := strconv.Atoi(min)
		maxInt, _ := strconv.Atoi(max)
		if passCount >= minInt && passCount <= maxInt {
			count++
		}
	}

	fmt.Println(count)
}
*/
