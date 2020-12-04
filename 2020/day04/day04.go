package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(file), "\n")

	passports := []string{}
	i := 0
	for _, line := range lines {
		if line != "" {
			if len(passports) == i {
				passports = append(passports, "")
			}
			passports[i] += " " + line
		} else {
			i++
		}
	}

	count := 0
	for _, passport := range passports {
		if validate(passport) {
			count++
		}
	}

	fmt.Println(count)
}

func validate(p string) bool {
	re := regexp.MustCompile(`[:\s]`)
	sp := re.Split(p, -1)[1:]
	m := map[string]string{}
	for i := 0; i < len(sp)-1; i += 2 {
		m[sp[i]] = sp[i+1]
	}
	if len(m) < 7 {
		return false
	}

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	if x, _ := strconv.Atoi(m["byr"]); x < 1920 || x > 2002 {
		return false
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	if x, _ := strconv.Atoi(m["iyr"]); x < 2010 || x > 2020 {
		return false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	if x, _ := strconv.Atoi(m["eyr"]); x < 2020 || x > 2030 {
		return false
	}

	// hgt (Height) - a number followed by either cm or in:
	if x := m["hgt"]; strings.Contains(x, "cm") {
		// If cm, the number must be at least 150 and at most 193.
		if h, _ := strconv.Atoi(strings.Split(x, "cm")[0]); h < 150 || h > 193 {
			return false
		}
	} else if strings.Contains(x, "in") {
		// If in, the number must be at least 59 and at most 76.
		if h, _ := strconv.Atoi(strings.Split(x, "in")[0]); h < 59 || h > 76 {
			return false
		}
	} else {
		return false
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	if x := m["hcl"]; true {
		valid, _ := regexp.Match("#[0-9a-f]{6}", []byte(x))
		if !valid {
			return false
		}
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	if x := m["ecl"]; true {
		valid, _ := regexp.Match("^(amb|blu|brn|gry|grn|hzl|oth)$", []byte(x))
		if !valid {
			return false
		}
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	if x := m["pid"]; true {
		valid, _ := regexp.Match("^[0-9]{9}$", []byte(x))
		if !valid {
			return false
		}
	}

	return true
}
