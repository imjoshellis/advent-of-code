package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var myTicketInput = `191,89,73,139,71,103,109,53,97,179,59,67,79,101,113,157,61,107,181,137`

func main() {
	// read and parse the rules input
	rulesTxt, _ := ioutil.ReadFile("rules.txt")
	rules := MakeRules(rulesTxt)

	// read and parse the tickets input
	ticketsTxt, _ := ioutil.ReadFile("tickets.txt")
	tickets := ValidTickets(ticketsTxt, rules)

	// convert the tickets and rules into a map
	ticketMap := MakeTicketMap(tickets, rules)

	// parse and print the answer
	fmt.Println(ParseMyTicket(ticketMap))
}

// Rule is boolean slice where Rule[n] coresponds to whether n is valid.
type Rule []bool

// MakeRule takes a single line from rulesTxt and turns it into a Rule.
func MakeRule(ruleStr string) Rule {
	rule := make(Rule, 1000)
	ss := strings.Split(ruleStr, " ")
	low := strings.Split(ss[len(ss)-3], "-")
	lowStart, _ := strconv.Atoi(low[0])
	lowEnd, _ := strconv.Atoi(low[1])
	for i := lowStart; i <= lowEnd; i++ {
		rule[i] = true
	}

	high := strings.Split(ss[len(ss)-1], "-")
	highStart, _ := strconv.Atoi(high[0])
	highEnd, _ := strconv.Atoi(high[1])
	for i := highStart; i <= highEnd; i++ {
		rule[i] = true
	}
	return rule
}

// MakeRules converts raw the input into a slice of Rules.
func MakeRules(rulesInput []byte) []Rule {
	rulesStr := strings.Split(string(rulesInput), "\n")
	rules := make([]Rule, len(rulesStr))
	for i, rule := range rulesStr {
		rules[i] = MakeRule(rule)
	}
	return rules
}

// MergedRule is a merged version of all the rules that makes it easier to  check whether a given ticket is valid.
func MergedRule(rules []Rule) Rule {
	merged := make(Rule, 1000)
	for _, rule := range rules {
		for i := range rule {
			if rule[i] {
				merged[i] = true
			}
		}
	}
	return merged
}

// A Ticket is a slice where each element is an int representing the field.
type Ticket []int

// MakeTicket converts a raw string into a Ticket
func MakeTicket(ticketStr string) Ticket {
	split := strings.Split(ticketStr, ",")
	ticket := make([]int, len(split))
	for i, el := range split {
		n, _ := strconv.Atoi(el)
		ticket[i] = n
	}
	return ticket
}

// IsValid checks whether a ticket is valid
func (t *Ticket) IsValid(rule Rule) bool {
	for _, n := range *t {
		if !rule[n] {
			return false
		}
	}
	return true
}

// ValidTickets takes in Rules and raw ticket input and returns a list of Tickets.
func ValidTickets(ticketsInput []byte, rules []Rule) []Ticket {
	tt := strings.Split(string(ticketsInput), "\n")
	rule := MergedRule(rules)
	tickets := make([]Ticket, 0)
	for _, t := range tt {
		ticket := MakeTicket(t)
		if ticket.IsValid(rule) {
			tickets = append(tickets, ticket)
		}
	}
	return tickets
}

// MakeDP creates a 2D DP slice.
func MakeDP(l int) [][]bool {
	dp := make([][]bool, l)
	for i := range dp {
		dp[i] = make([]bool, l)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}
	return dp
}

// MakeTicketMap turns Tickets and Rules into a slice that maps from ticket fields to rules.
func MakeTicketMap(tickets []Ticket, rules []Rule) []int {
	dp := MakeDP(len(rules))
	for _, ticket := range tickets {
		for i, n := range ticket {
			for j, bools := range rules {
				if !bools[n] {
					dp[i][j] = false
				}
			}
		}
	}

	results := make([]int, len(dp))
	for count := 0; count < len(dp); count++ {
	mainloop:
		for i := range dp {
			boolIdx := -1
			for j := range dp {
				if dp[i][j] {
					if boolIdx != -1 {
						continue mainloop
					}
					boolIdx = j
				}
			}
			if (boolIdx) == -1 {
				continue
			}
			results[i] = boolIdx

			for k := range dp {
				dp[k][boolIdx] = false
			}
		}
	}

	return results
}

// ParseMyTicket uses a map to find the first 6 rules and multiply their values in order to get the final answer for the challenge.
func ParseMyTicket(m []int) int {
	product := 1
	ss := strings.Split(myTicketInput, ",")
	myTicketArr := []int{}
	for _, num := range ss {
		n, _ := strconv.Atoi(num)
		myTicketArr = append(myTicketArr, n)
	}

	for i := 0; i < 6; i++ {
		for j, n := range m {
			if n == i {
				product *= myTicketArr[j]
			}
		}
	}
	return product
}
