package main

import (
	"bufio"
	"github.com/juliangruber/go-intersect"
	"log"
	"os"
	"sort"
	"strings"
)

type Backpack struct {
	AllItems       []string
	CompartmentOne []string
	CompartmentTwo []string
}

func (b *Backpack) GetType() string {
	_type := intersect.Simple(b.CompartmentOne, b.CompartmentTwo)
	return _type[0].(string)
}

func (b *Backpack) GetPriority() int {
	return GetTypePriority(b.GetType())
}

type Group struct {
	Backpacks []*Backpack
}

func (g *Group) GetBadge() string {
	sort.SliceStable(g.Backpacks, func(i, j int) bool {
		return len(g.Backpacks[i].AllItems) > len(g.Backpacks[j].AllItems)
	})

	intersection := intersect.Simple(g.Backpacks[0].AllItems, g.Backpacks[1].AllItems)
	for _, s := range intersection {
		for _, c := range g.Backpacks[2].AllItems {
			if c == s {
				return c
			}
		}
	}

	log.Fatal("No intersection found??")
	return ""
}

func (g *Group) GetPriority() int {
	return GetTypePriority(g.GetBadge())
}

func GetTypePriority(_type string) int {
	if _type == "" {
		return 0
	}

	char := _type[0]
	asciiCode := int(char) - 96

	if asciiCode < 0 {
		asciiCode += 58
	}

	return asciiCode
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	total := 0
	totalGroup := 0

	currentGroup := Group{}
	for scanner.Scan() {
		line := scanner.Text()

		chars := strings.Split(line, "")
		backpack := Backpack{
			AllItems:       chars,
			CompartmentOne: chars[0 : len(chars)/2],
			CompartmentTwo: chars[len(chars)/2 : len(chars)],
		}
		currentGroup.Backpacks = append(currentGroup.Backpacks, &backpack)
		total += backpack.GetPriority()

		if len(currentGroup.Backpacks) == 3 {
			totalGroup += currentGroup.GetPriority()
			currentGroup = Group{}
		}
	}

	println("Total points are: ")
	println(total)
	println(totalGroup)
}
