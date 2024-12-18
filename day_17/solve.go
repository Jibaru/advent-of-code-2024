package day17

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Solve(part int, isTest bool) (any, error) {
	f := "day_17/input.txt"
	if isTest {
		f = "day_17/input-test.txt"
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
	program, err := NewProgram(data)
	if err != nil {
		return nil, err
	}

	if err = program.Run(); err != nil {
		return nil, err
	}
	return program.Output(), nil
}

func partTwo(data string) (any, error) {
	program, err := NewProgram(data)
	if err != nil {
		return nil, err
	}

	length := len(program.instructions)
	validNums := []int{0}

	b := program.regB
	c := program.regC

	for l := 1; l <= length; l++ {
		oldValidNums := validNums
		validNums = []int{}

		for _, num := range oldValidNums {
			for offset := 0; offset < 8; offset++ {
				newNum := 8*num + offset
				program.SetRegisters(newNum, b, c)
				program.ClearOutput()
				program.ResetPointer()

				if err := program.Run(); err != nil {
					return nil, err
				}

				output := program.output

				if len(output) > 0 && ContainsAtEnd(output, program.instructions) {
					validNums = append(validNums, newNum)
				}
			}
		}
	}

	if len(validNums) == 0 {
		return nil, fmt.Errorf("no valid number found")
	}

	return Min(validNums), nil
}

const (
	adv = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

type Program struct {
	regA         int
	regB         int
	regC         int
	instructions []int
	output       []int
	pointer      int
	jumpCalled   bool
}

func NewProgram(input string) (*Program, error) {
	registerRegex := regexp.MustCompile(`Register A:\s*(\d+)\s*Register B:\s*(\d+)\s*Register C:\s*(\d+)`)
	programRegex := regexp.MustCompile(`Program:\s*([\d,]+)`)

	registerMatch := registerRegex.FindStringSubmatch(input)
	if registerMatch == nil {
		return nil, fmt.Errorf("invalid input format for registers")
	}

	regA, _ := strconv.Atoi(registerMatch[1])
	regB, _ := strconv.Atoi(registerMatch[2])
	regC, _ := strconv.Atoi(registerMatch[3])

	programMatch := programRegex.FindStringSubmatch(input)
	if programMatch == nil {
		return nil, fmt.Errorf("invalid input format for program")
	}

	instructionParts := strings.Split(programMatch[1], ",")
	if len(instructionParts)%2 != 0 {
		return nil, fmt.Errorf("instructions must be in pairs")
	}

	var instructions []int
	for i := 0; i < len(instructionParts); i++ {
		v, _ := strconv.Atoi(strings.TrimSpace(instructionParts[i]))
		instructions = append(instructions, v)
	}

	return &Program{
		regA:         regA,
		regB:         regB,
		regC:         regC,
		instructions: instructions,
		output:       make([]int, 0),
		pointer:      0,
		jumpCalled:   false,
	}, nil
}

func (p *Program) Output() string {
	output := make([]string, len(p.output))
	for i, o := range p.output {
		output[i] = fmt.Sprintf("%v", o)
	}

	return strings.Join(output, ",")
}

func (p *Program) HasInstructionsEqualsToOutput() bool {
	if len(p.instructions) != len(p.output) {
		return false
	}
	for i := range p.instructions {
		if p.instructions[i] != p.output[i] {
			return false
		}
	}
	return true
}

func (p *Program) Run() error {
	totalInstructions := len(p.instructions)

	for p.pointer < totalInstructions {
		opcode := p.instructions[p.pointer]
		operand := p.instructions[p.pointer+1]

		err := p.do(opcode, Operand(operand))
		if err != nil {
			return err
		}

		if !p.jumpCalled {
			p.pointer += 2
		}
	}

	return nil
}

func (p *Program) ClearOutput() {
	p.output = make([]int, 0)
	p.pointer = 0
}

func (p *Program) ResetPointer() {
	p.pointer = 0
}

func (p *Program) SetRegisters(a, b, c int) {
	p.regA = a
	p.regB = b
	p.regC = c
}

func (p *Program) do(opcode int, operand Operand) error {
	p.jumpCalled = false

	switch opcode {
	case adv:
		denominator, err := operand.Combo(p)
		if err != nil {
			return err
		}
		numerator := p.regA
		p.regA = numerator / int(math.Pow(2, float64(denominator)))
	case bxl:
		p.regB ^= operand.Literal()
	case bst:
		v, err := operand.Combo(p)
		if err != nil {
			return err
		}
		p.regB = v % 8
	case jnz:
		if p.regA != 0 {
			p.pointer = operand.Literal()
			p.jumpCalled = true
		}
	case bxc:
		p.regB ^= p.regC
	case out:
		v, err := operand.Combo(p)
		if err != nil {
			return err
		}
		p.output = append(p.output, v%8)
	case bdv:
		denominator, err := operand.Combo(p)
		if err != nil {
			return err
		}
		numerator := p.regA
		p.regB = numerator / int(math.Pow(2, float64(denominator)))
	case cdv:
		denominator, err := operand.Combo(p)
		if err != nil {
			return err
		}
		numerator := p.regA
		p.regC = numerator / int(math.Pow(2, float64(denominator)))
	}

	return nil
}

type Operand int

func (o Operand) Combo(p *Program) (int, error) {
	if o <= 3 {
		return o.Literal(), nil
	}

	if o == 4 {
		return p.regA, nil
	}

	if o == 5 {
		return p.regB, nil
	}

	if o == 6 {
		return p.regC, nil
	}

	return 0, fmt.Errorf("invalid combo operand %v", o)
}

func (o Operand) Literal() int {
	return int(o)
}

func Min(nums []int) int {
	mn := nums[0]
	for _, v := range nums {
		if v < mn {
			mn = v
		}
	}

	return mn
}

func ContainsAtEnd(a, b []int) bool {
	if len(a) > len(b) {
		return false
	}

	start := len(b) - len(a)
	for i := 0; i < len(a); i++ {
		if a[i] != b[start+i] {
			return false
		}
	}

	return true
}
