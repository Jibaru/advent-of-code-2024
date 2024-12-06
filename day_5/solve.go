package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_5/input.txt"
	if isTest {
		f = "day_5/input-test.txt"
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
	rules, updates := parse(data)
	graph := newDependencyGraph(rules)

	sum := 0
	for _, update := range updates {
		if graph.isValid(update) {
			sum += update[len(update)/2]
		}
	}

	return sum, nil
}

func partTwo(data string) (any, error) {
	rules, updates := parse(data)
	graph := newDependencyGraph(rules)

	ans := 0
	for _, update := range updates {
		if graph.isValid(update) {
			continue
		}

		sortedUpdate := graph.topologicalSort(update)
		ans += sortedUpdate[len(update)/2]
	}

	return ans, nil
}

type Rule struct {
	before int
	after  int
}

type Graph map[int][]int

func parse(data string) ([]Rule, [][]int) {
	parts := strings.Split(strings.TrimSpace(data), "\n\n")
	rulesSection := strings.Split(parts[0], "\n")
	updatesSection := strings.Split(parts[1], "\n")

	var rules []Rule
	for _, rule := range rulesSection {
		parts := strings.Split(rule, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		rules = append(rules, Rule{before: before, after: after})
	}

	var updates [][]int
	for _, update := range updatesSection {
		pagesStr := strings.Split(update, ",")
		var pages []int
		for _, pageStr := range pagesStr {
			page, _ := strconv.Atoi(pageStr)
			pages = append(pages, page)
		}
		updates = append(updates, pages)
	}

	return rules, updates
}

func newDependencyGraph(rules []Rule) Graph {
	graph := make(Graph)
	for _, rule := range rules {
		graph[rule.before] = append(graph[rule.before], rule.after)
	}
	return graph
}

func (g Graph) isValid(update []int) bool {
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}

	for x, dependents := range g {
		if posX, exists := position[x]; exists {
			for _, y := range dependents {
				if posY, exists := position[y]; exists && posX > posY {
					return false
				}
			}
		}
	}

	return true
}

func (g Graph) topologicalSort(update []int) []int {
	// Build a restricted graph and in-degree map for the current update
	inUpdate := make(map[int]bool)
	for _, page := range update {
		inUpdate[page] = true
	}

	restrictedGraph := make(map[int][]int)
	inDegree := make(map[int]int)
	for page := range inUpdate {
		for _, dependent := range g[page] {
			if inUpdate[dependent] {
				restrictedGraph[page] = append(restrictedGraph[page], dependent)
				inDegree[dependent]++
			}
		}
	}

	// Initialize the queue with pages having no prerequisites
	var queue []int
	for _, page := range update {
		if inDegree[page] == 0 {
			queue = append(queue, page)
		}
	}

	// Perform topological sort
	var sorted []int
	visited := make(map[int]bool)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		sorted = append(sorted, current)
		visited[current] = true

		for _, dependent := range restrictedGraph[current] {
			inDegree[dependent]--
			if inDegree[dependent] == 0 && !visited[dependent] {
				queue = append(queue, dependent)
			}
		}
	}

	// If some pages are left unsorted (due to cycles), append them as-is
	for _, page := range update {
		if !visited[page] {
			sorted = append(sorted, page)
		}
	}

	return sorted
}
