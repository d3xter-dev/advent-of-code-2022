package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	Calories []int
}

func (e *Elf) GetCalories() int {
	calories := 0
	for _, item := range e.Calories {
		calories = calories + item
	}

	return calories
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	var elfList []Elf
	current := Elf{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elfList = append(elfList, current)
			current = Elf{}
			continue
		}

		calories, e := strconv.Atoi(line)
		if e != nil {
			log.Fatal(e)
		}

		current.Calories = append(current.Calories, calories)
	}
	elfList = append(elfList, current)

	// Sort elves by calories
	sort.SliceStable(elfList, func(i, j int) bool {
		return elfList[i].GetCalories() > elfList[j].GetCalories()
	})

	// Echo winner
	println("Top Three is: ")
	println(elfList[0].GetCalories())
	println(elfList[1].GetCalories())
	println(elfList[2].GetCalories())

	println("Top Three Total is: ")
	println(elfList[0].GetCalories() + elfList[1].GetCalories() + elfList[2].GetCalories())
}
