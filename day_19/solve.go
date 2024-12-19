package day19

import (
	"fmt"
	"os"
	"strings"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_19/input.txt"
	if isTest {
		f = "day_19/input-test.txt"
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
	patterns, designs := parseInput(data)
	designer := NewDesigner(patterns)

	ans := 0
	for _, design := range designs {
		if designer.canFormDesign(design) {
			ans++
		}
	}

	return ans, nil
}

func partTwo(data string) (any, error) {
	patterns, designs := parseInput(data)
	designer := NewDesigner(patterns)

	ans := 0
	for _, design := range designs {
		ans += designer.countWaysToFormDesign(design)
	}

	return ans, nil
}

func parseInput(data string) (patterns []string, designs []string) {
	parts := strings.Split(data, "\n\n")
	patterns = append(patterns, strings.Split(parts[0], ", ")...)
	designs = append(designs, strings.Split(parts[1], "\n")...)
	return
}

type Designer struct {
	patterns   []string
	validCache map[string]bool
	countCache map[string]int
}

func NewDesigner(patterns []string) *Designer {
	return &Designer{
		patterns:   patterns,
		validCache: make(map[string]bool),
		countCache: make(map[string]int),
	}
}

func (d *Designer) canFormDesign(design string) bool {
	if result, ok := d.validCache[design]; ok {
		return result
	}

	if design == "" {
		d.validCache[design] = true
		return true // Base case: an empty design is always valid.
	}

	for _, pattern := range d.patterns {
		if strings.HasPrefix(design, pattern) {
			// Recursively check if the remaining part of the design can be formed.
			if d.canFormDesign(design[len(pattern):]) {
				d.validCache[design] = true
				return true
			}
		}
	}

	d.validCache[design] = false
	return false // No pattern fits the current prefix.
}

func (d *Designer) countWaysToFormDesign(design string) int {
	if count, ok := d.countCache[design]; ok {
		return count
	}

	if design == "" {
		d.countCache[design] = 1
		return 1 // Base case: one way to form an empty design.
	}

	totalWays := 0
	for _, pattern := range d.patterns {
		if strings.HasPrefix(design, pattern) {
			// Recursively count ways for the remaining part of the design.
			totalWays += d.countWaysToFormDesign(design[len(pattern):])
		}
	}

	d.countCache[design] = totalWays

	return totalWays
}
