package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Color string

type Bag struct {
	color Color
}

type BagEdge struct {
	bag   *Bag
	count int
}

type BagGraph struct {
	bags  map[Color]*Bag
	edges map[Bag][]*BagEdge
}

// ToString serializes a bag to string
func (b *Bag) ToString() string {
	return fmt.Sprintf("%v", b.color)
}

// AddBag adds a bag (node) to the graph
func (g *BagGraph) AddBag(b *Bag) *Bag {
	if g.bags == nil {
		g.bags = make(map[Color]*Bag)
	}
	g.bags[b.color] = b
	return b
}

// AddEdge adds a directional edge to the graph
func (g *BagGraph) AddEdge(b1, b2 *Bag, count int) {
	if g.edges == nil {
		g.edges = make(map[Bag][]*BagEdge)
	}
	edge := BagEdge{b2, count}
	g.edges[*b1] = append(g.edges[*b1], &edge)
}

// GetBag is a wrapper for looking up bags in the bag map
func (g *BagGraph) GetBag(c Color) (*Bag, bool) {
	bag, ok := g.bags[c]
	return bag, ok
}

func main() {
	// read the input (which is manually cleaned up)
	f, _ := ioutil.ReadFile("./input.txt")
	lines := strings.Split(string(f), "\n")

	// initialize graph
	graph := BagGraph{}

	// parse input line by line
	for _, line := range lines {
		// split into bag / edges
		// because of how input was formatted,
		// this is guaranteed to be [2]string
		// where the first element is a bag color
		// and the second is a list of what it contains
		kv := strings.Split(line, "#")

		// get the color off the first element
		color := Color(kv[0])

		// get edges off the second element
		edges := strings.Split(kv[1], ", ")

		// get bag if it's in the graph
		bag, ok := graph.GetBag(color)
		if !ok {
			// if bag isn't in graph, make new and add
			bag = graph.AddBag(&Bag{color: color})
		}

		// iterate over edges
		for _, edge := range edges {
			// split edge by spaces
			split := strings.Split(edge, " ")

			// the first element will be the count
			count, _ := strconv.Atoi(split[0])

			// the second element is the color
			color := Color(strings.Join(split[1:], " "))

			// if the bag is in the graph...
			if bag2, ok := graph.GetBag(color); ok {
				// ...add it as an edge to the current bag
				graph.AddEdge(bag, bag2, count)
			} else {
				// ...otherwise, make a new bag and add it as an edge
				bag2 := graph.AddBag(&Bag{color: color})
				graph.AddEdge(bag, bag2, count)
			}
		}
	}

	// part 1
	count := 0
	for _, bag := range graph.bags {
		// count every bag that could contain shiny gold
		if bag.CanContain("shiny gold", &graph) {
			count++
		}
	}
	fmt.Println(count, "bags can contain shiny gold /", len(graph.bags), "total bags")

	// part 2: count how many bags would go inside shiny gold
	sg, _ := graph.GetBag("shiny gold")
	fmt.Println("shiny gold would have", sg.CountInside(&graph), "bags inside it")
}

func (b *Bag) CountInside(g *BagGraph) (count int) {
	// special type for the stack
	type BagStack struct {
		bag   Bag
		count int
	}
	st := []BagStack{{bag: *b, count: 1}}

	for len(st) > 0 {
		// take the head off of stack
		cur := st[0]
		st = st[1:]

		// add to count
		count += cur.count

		// iterate over edges
		for _, edge := range g.edges[cur.bag] {
			// calculate how many of this bag there will be
			curCount := edge.count * cur.count

			// add to stack
			st = append(st, BagStack{bag: *edge.bag, count: curCount})
		}
	}

	// return count - 1 because the first bag shouldn't be counted
	return count - 1
}

func (b *Bag) CanContain(c Color, g *BagGraph) bool {
	// initialize visited map
	visited := make(map[Bag]bool)

	// initialize stack of bags
	st := []Bag{*b}

	// loop until stack is empty
	for len(st) > 0 {
		// get first element off of stack
		cur := st[0]
		st = st[1:]

		// skip if already visited
		if visited[cur] {
			continue
		}

		// update visited
		visited[cur] = true

		// iterate over edges
		for _, edge := range g.edges[cur] {
			// if edge color is the target, we can return
			if edge.bag.color == c {
				return true
			}

			// if the edge isn't in visited, add it to the stack
			// could just add them to stack and let the earlier
			// conditional take care of it...
			if _, ok := visited[*edge.bag]; !ok {
				st = append(st, *edge.bag)
			}
		}
	}

	// if we get here, there's no path to target color
	return false
}
