package main

import (
	"flag"
	"fmt"

	day0 "github.com/jibaru/advent-of-code-2024/day_0"
	day1 "github.com/jibaru/advent-of-code-2024/day_1"
	day2 "github.com/jibaru/advent-of-code-2024/day_2"
	day3 "github.com/jibaru/advent-of-code-2024/day_3"
)

func main() {
	day := flag.Int("d", 0, "Specify the day")
	part := flag.Int("p", 1, "Specify part of the day (1 or 2)")
	isTest := flag.Bool("t", false, "Specify is the input is test")

	flag.Parse()

	var answer any
	var err error
	switch *day {
	case 0:
		answer, err = day0.Solve(*part, *isTest)
	case 1:
		answer, err = day1.Solve(*part, *isTest)
	case 2:
		answer, err = day2.Solve(*part, *isTest)
	case 3:
		answer, err = day3.Solve(*part, *isTest)
	default:
		err = fmt.Errorf("day not allowed")
	}

	if err != nil {
		fmt.Printf("error happened: %v\n", err)
	} else {
		fmt.Printf("answer for day %v part %v: %v\n", *day, *part, answer)
	}
}
