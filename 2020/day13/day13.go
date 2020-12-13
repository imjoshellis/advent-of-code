package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"
	"strings"
)

// part 1
// func main() {
// 	f, _ := ioutil.ReadFile("input.txt")
// 	lines := strings.Split(string(f), "\n")
// 	timestamp := lines[0]
// 	intTS, _ := strconv.Atoi(timestamp)
// 	IDs := strings.Split(lines[1], ",")
// 	cleanIDs := make([]int, 0)
// 	stamps := make([]int, 0)
// 	for _, ID := range IDs {
// 		if ID != "x" {
// 			intID, _ := strconv.Atoi(ID)
// 			cleanIDs = append(cleanIDs, intID)
// 			var i int
// 			for i = intID; i < intTS; i += intID {
// 			}
// 			stamps = append(stamps, i)
// 		}
// 	}
// 	fmt.Println(intTS)
// 	fmt.Println(cleanIDs)
// 	fmt.Println(stamps)
// }

func main() {
	f, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")
	IDs := strings.Split(lines[1], ",")
	cleanIDs := make([]int64, 0)
	for _, ID := range IDs {
		if ID != "x" {
			intID, _ := strconv.Atoi(ID)
			cleanIDs = append(cleanIDs, int64(intID))
		} else {
			cleanIDs = append(cleanIDs, 1)
		}
	}
	// found using math site...
	// work smarter, not harder???
	// no clue how to code chinese remainder
	// https://www.dcode.fr/chinese-remainder
	// x = 1106724616194525

	// using code (see helpers below)...
	n, a := []*big.Int{}, []*big.Int{}

	var j int64
	for j = 0; int(j) < len(cleanIDs); j++ {
		if cleanIDs[j] != 1 {
			a = append(a, big.NewInt((cleanIDs[j]-j)%cleanIDs[j]))
			n = append(n, big.NewInt(cleanIDs[j]))
		}
	}
	fmt.Println(crt(a, n))
}

// code from https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
var one = big.NewInt(1)

func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}
