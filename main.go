package main

import (
	"flag"
	"fmt"

	day0 "github.com/jibaru/advent-of-code-2024/day_0"
	day1 "github.com/jibaru/advent-of-code-2024/day_1"
	day10 "github.com/jibaru/advent-of-code-2024/day_10"
	day11 "github.com/jibaru/advent-of-code-2024/day_11"
	day12 "github.com/jibaru/advent-of-code-2024/day_12"
	day13 "github.com/jibaru/advent-of-code-2024/day_13"
	day14 "github.com/jibaru/advent-of-code-2024/day_14"
	day15 "github.com/jibaru/advent-of-code-2024/day_15"
	day16 "github.com/jibaru/advent-of-code-2024/day_16"
	day17 "github.com/jibaru/advent-of-code-2024/day_17"
	day18 "github.com/jibaru/advent-of-code-2024/day_18"
	day19 "github.com/jibaru/advent-of-code-2024/day_19"
	day2 "github.com/jibaru/advent-of-code-2024/day_2"
	day20 "github.com/jibaru/advent-of-code-2024/day_20"
	day21 "github.com/jibaru/advent-of-code-2024/day_21"
	day22 "github.com/jibaru/advent-of-code-2024/day_22"
	day3 "github.com/jibaru/advent-of-code-2024/day_3"
	day4 "github.com/jibaru/advent-of-code-2024/day_4"
	day5 "github.com/jibaru/advent-of-code-2024/day_5"
	day6 "github.com/jibaru/advent-of-code-2024/day_6"
	day7 "github.com/jibaru/advent-of-code-2024/day_7"
	day8 "github.com/jibaru/advent-of-code-2024/day_8"
	day9 "github.com/jibaru/advent-of-code-2024/day_9"
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
	case 4:
		answer, err = day4.Solve(*part, *isTest)
	case 5:
		answer, err = day5.Solve(*part, *isTest)
	case 6:
		answer, err = day6.Solve(*part, *isTest)
	case 7:
		answer, err = day7.Solve(*part, *isTest)
	case 8:
		answer, err = day8.Solve(*part, *isTest)
	case 9:
		answer, err = day9.Solve(*part, *isTest)
	case 10:
		answer, err = day10.Solve(*part, *isTest)
	case 11:
		answer, err = day11.Solve(*part, *isTest)
	case 12:
		answer, err = day12.Solve(*part, *isTest)
	case 13:
		answer, err = day13.Solve(*part, *isTest)
	case 14:
		answer, err = day14.Solve(*part, *isTest)
	case 15:
		answer, err = day15.Solve(*part, *isTest)
	case 16:
		answer, err = day16.Solve(*part, *isTest)
	case 17:
		answer, err = day17.Solve(*part, *isTest)
	case 18:
		answer, err = day18.Solve(*part, *isTest)
	case 19:
		answer, err = day19.Solve(*part, *isTest)
	case 20:
		answer, err = day20.Solve(*part, *isTest)
	case 21:
		answer, err = day21.Solve(*part, *isTest)
	case 22:
		answer, err = day22.Solve(*part, *isTest)
	default:
		err = fmt.Errorf("day not allowed")
	}

	if err != nil {
		fmt.Printf("error happened: %v\n", err)
	} else {
		fmt.Printf("answer for day %v part %v: %v\n", *day, *part, answer)
	}
}
