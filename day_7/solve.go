package day7

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_7/input.txt"
	if isTest {
		f = "day_7/input-test.txt"
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
	operators := []string{"+", "*"}
	return solveBoth(data, operators)
}

func partTwo(data string) (any, error) {
	operators := []string{"+", "*", "||"}
	return solveBoth(data, operators)
}

func solveBoth(data string, operators []string) (any, error) {
	operations := newOperations(data)
	combinator := newCombinator()

	ans := 0
	for _, op := range operations {
		combinations := combinator.Generate(len(op.Operands)-1, len(operators))
		for _, combination := range combinations {
			result := evaluate(op.Operands, combination, operators)
			if result == op.Result {
				ans += result
				break
			}
		}
	}
	return ans, nil
}

type Operation struct {
	Result   int
	Operands []int
}

func newOperations(data string) []Operation {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	operations := []Operation{}

	for _, line := range lines {
		resultWithOperands := strings.Split(line, ":")
		operands := strings.Split(strings.TrimSpace(resultWithOperands[1]), " ")

		result, _ := strconv.ParseInt(resultWithOperands[0], 10, 64)

		var ops []int
		for _, op := range operands {
			parsedOp, _ := strconv.ParseInt(op, 10, 64)
			ops = append(ops, int(parsedOp))
		}

		operations = append(operations, Operation{
			Result:   int(result),
			Operands: ops,
		})
	}

	return operations
}

type Combinator struct {
	cache map[int]map[int][][]int
}

func newCombinator() *Combinator {
	return &Combinator{
		cache: map[int]map[int][][]int{},
	}
}

func (c *Combinator) Generate(numOfOperands, numOfOperators int) [][]int {
	if v1, ok := c.cache[numOfOperands]; ok {
		if v2, ok := v1[numOfOperators]; ok {
			return v2
		}
	}

	total := int(math.Pow(float64(numOfOperators), float64(numOfOperands)))
	combinations := make([][]int, 0, total)

	for i := 0; i < total; i++ {
		combination := make([]int, numOfOperands)
		num := i
		for j := 0; j < numOfOperands; j++ {
			combination[j] = num % numOfOperators
			num /= numOfOperators
		}
		combinations = append(combinations, combination)
	}

	if _, ok := c.cache[numOfOperands]; !ok {
		c.cache[numOfOperands] = map[int][][]int{}
	}

	c.cache[numOfOperands][numOfOperators] = combinations

	return combinations
}

func evaluate(operands []int, combination []int, operators []string) int {
	result := operands[0]

	for i, operatorIdx := range combination {
		operator := operators[operatorIdx]

		switch operator {
		case "+":
			result += operands[i+1]
		case "*":
			result *= operands[i+1]
		case "||":
			joined, _ := strconv.ParseInt(fmt.Sprintf("%v%v", result, operands[i+1]), 10, 64)
			result = int(joined)
		}
	}

	return result
}
