package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("./input.txt")
	lns := strings.Split(string(f), "\n")

	ps := []string{}
	i := 0
	for _, ln := range lns {
		if ln != "" {
			if len(ps) == i {
				ps = append(ps, "")
			}
			ps[i] += " " + ln
		} else {
			i++
		}
	}

	c := 0
	for _, p := range ps {
		if validate(p) {
			c++
		}
	}

	fmt.Println(c)
}

func validate(p string) bool {
	sp := regexp.MustCompile(`[:\s]`).Split(p, -1)[1:]
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

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	hclRe := "#[0-9a-f]{6}"
	validHcl, _ := regexp.Match(hclRe, []byte(m["hcl"]))
	if !validHcl {
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	eclRe := "^(amb|blu|brn|gry|grn|hzl|oth)$"
	validEcl, _ := regexp.Match(eclRe, []byte(m["ecl"]))
	if !validEcl {
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	pidRe := "^[0-9]{9}$"
	validPid, _ := regexp.Match(pidRe, []byte(m["pid"]))
	if !validPid {
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

	return true
}
