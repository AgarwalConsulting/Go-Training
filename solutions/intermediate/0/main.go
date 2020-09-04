package main

import (
	"fmt"
	"strings"
)

func filter(s []string, predicates ...func(string) bool) []string {
	var output []string

	for _, element := range s {
		isSelected := true
		for _, predicate := range predicates {
			isSelected = isSelected && predicate(element)
		}

		if isSelected {
			output = append(output, element)
		}
	}

	return output
}

func main() {
	heroes := []string{"Iron Man", "Batman", "Superman", "Spider-man", "Wonder Woman", "Iron Fist", "Daredevil", "Supergirl", "Flash"}

	heroesWithVowelsAsSecondChar := filter(heroes, func(hero string) bool {
		return strings.Contains("aeiou", string(hero[1]))
	})

	heroesWithoutMan := filter(heroes, func(hero string) bool {
		return !strings.Contains(strings.ToLower(hero), "man")
	})

	fmt.Println(heroesWithVowelsAsSecondChar)
	fmt.Println(heroesWithoutMan)
}
