package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const (
	Rock     = "A"
	Paper    = "B"
	Scissors = "C"
)

var Points = map[string]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

func CheckPoints(a string, b string) int {
	_points := Points[b]
	if a == b {
		_points = _points + 3
	}

	if (a == Scissors && b == Rock) || (a == Rock && b == Paper) || (a == Paper && b == Scissors) {
		_points = _points + 6
	}

	return _points
}

func CheckStrategyOne(a string, b string) int {
	strategy := map[string]string{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}

	return CheckPoints(a, strategy[b])
}

func CheckStrategyTwo(a string, b string) int {
	if b == "X" {
		switch a {
		case Rock:
			return CheckPoints(a, Scissors)
		case Scissors:
			return CheckPoints(a, Paper)
		case Paper:
			return CheckPoints(a, Rock)
		}
	}

	if b == "Z" {
		switch a {
		case Rock:
			return CheckPoints(a, Paper)
		case Scissors:
			return CheckPoints(a, Rock)
		case Paper:
			return CheckPoints(a, Scissors)
		}
	}

	return CheckPoints(a, a)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	totalStratOne := 0
	totalStratTwo := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		totalStratOne = totalStratOne + CheckStrategyOne(parts[0], parts[1])
		totalStratTwo = totalStratTwo + CheckStrategyTwo(parts[0], parts[1])
	}

	println("Total points are: ")
	println(totalStratOne)
	println(totalStratTwo)
}
