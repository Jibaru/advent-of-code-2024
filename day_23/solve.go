package day23

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_23/input.txt"
	if isTest {
		f = "day_23/input-test.txt"
	}

	body, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	switch part {
	case 1:
		return partOne(string(body))
	case 2:
		return partTwo(string(body))
	}

	return nil, fmt.Errorf("part should be only 1 or 2")
}

func partOne(data string) (any, error) {
	conns := parseInput(data)

	paths := cycles(conns, 3)
	ans := 0
	for _, path := range paths {
		if someComputerStartsWith(path, "t") {
			ans++
		}
	}

	return ans, nil
}

func partTwo(data string) (any, error) {
	conns := parseInput(data)
	graph := buildGraph(conns)
	bronkerbosh := &Bronkerbosh{graph: graph}
	names := bronkerbosh.FindLargest()

	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})

	return strings.Join(names, ","), nil
}

func parseInput(data string) [][]ComputerName {
	conns := make([][]ComputerName, 0)

	for _, line := range strings.Split(data, "\n") {
		parts := strings.Split(line, "-")
		conns = append(conns, []ComputerName{parts[0], parts[1]})
	}

	return conns
}

type ComputerName = string

type Set map[ComputerName]bool

func (s Set) Add(name ComputerName) {
	s[name] = true
}

func (s Set) Union(other Set) Set {
	result := make(Set)
	for k := range s {
		result[k] = true
	}
	for k := range other {
		result[k] = true
	}
	return result
}

func (s Set) Intersect(other Set) Set {
	result := make(Set)
	for k := range s {
		if other[k] {
			result[k] = true
		}
	}
	return result
}

func (s Set) Subtract(other Set) Set {
	result := make(Set)
	for k := range s {
		if !other[k] {
			result[k] = true
		}
	}
	return result
}

func (s Set) Copy() Set {
	result := make(Set)
	for k, v := range s {
		result[k] = v
	}
	return result
}

func someComputerStartsWith(names []ComputerName, letter string) bool {
	for _, name := range names {
		if strings.HasPrefix(string(name), letter) {
			return true
		}
	}
	return false
}

// buildGraph builds an adjacency list from the connections
func buildGraph(conns [][]ComputerName) map[ComputerName]Set {
	graph := make(map[ComputerName]Set)
	for _, conn := range conns {
		a, b := conn[0], conn[1]
		if graph[a] == nil {
			graph[a] = make(Set)
		}
		if graph[b] == nil {
			graph[b] = make(Set)
		}
		graph[a].Add(b)
		graph[b].Add(a)
	}
	return graph
}

// cycles finds all cliques of size n
func cycles(conns [][]ComputerName, n int) [][]ComputerName {
	graph := buildGraph(conns)

	// Generate all subsets of size n and check for cliques
	var nodes []ComputerName
	for node := range graph {
		nodes = append(nodes, node)
	}

	combiner := NewCombiner(nodes, n)
	subsets := combiner.Generate()

	var result [][]ComputerName
	for _, subset := range subsets {
		if isClique(graph, subset) {
			result = append(result, subset)
		}
	}

	return result
}

type Combiner struct {
	nodes []ComputerName
	size  int
}

func NewCombiner(nodes []ComputerName, size int) *Combiner {
	return &Combiner{nodes: nodes, size: size}
}

func (c *Combiner) Generate() [][]ComputerName {
	var result [][]ComputerName
	var temp []ComputerName
	c.combine(0, temp, &result)
	return result
}

func (c *Combiner) combine(start int, temp []ComputerName, result *[][]ComputerName) {
	if len(temp) == c.size {
		*result = append(*result, append([]ComputerName{}, temp...))
		return
	}
	for i := start; i < len(c.nodes); i++ {
		c.combine(i+1, append(temp, c.nodes[i]), result)
	}
}

// isClique checks if the subset is a clique
func isClique(graph map[ComputerName]Set, subset []ComputerName) bool {
	for i := 0; i < len(subset); i++ {
		for j := i + 1; j < len(subset); j++ {
			if !graph[subset[i]][subset[j]] {
				return false
			}
		}
	}
	return true
}

type Bronkerbosh struct {
	largest []ComputerName
	graph   map[ComputerName]Set
}

func (b *Bronkerbosh) FindLargest() []ComputerName {
	// Initialize sets
	allNodes := make(Set)
	for node := range b.graph {
		allNodes[node] = true
	}

	b.largest = make([]ComputerName, 0)
	b.Do(make(Set), allNodes, make(Set))

	return b.largest
}

func (b *Bronkerbosh) Do(currentCycle, candidates, excluded Set) {
	if len(candidates) == 0 && len(excluded) == 0 {
		// Found a maximal clique
		if len(currentCycle) > len(b.largest) {
			b.largest = make([]ComputerName, 0, len(currentCycle))
			for node := range currentCycle {
				b.largest = append(b.largest, node)
			}
		}
		return
	}

	// Choose a pivot to reduce branching
	var u ComputerName
	for u = range candidates.Union(excluded) {
		break
	}

	// Iterate over candidates not in neighbors of the pivot
	for v := range candidates.Subtract(b.graph[u]) {
		newCycle := currentCycle.Copy()
		newCycle[v] = true

		newCandidates := candidates.Intersect(b.graph[v])
		newExcluded := excluded.Intersect(b.graph[v])

		b.Do(newCycle, newCandidates, newExcluded)

		delete(candidates, v)
		excluded[v] = true
	}
}
