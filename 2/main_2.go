package main

import (
	"fmt"
)

func sliceExample(s []int) []int {
	evenSlice := make([]int, 0, len(s))

	for _, value := range s {
		if value%2 == 0 {
			evenSlice = append(evenSlice, value)
		}
	}
	return evenSlice
}

func addElements(s []int, elment int) []int {
	newSlice := make([]int, len(s)+1)
	copy(newSlice, s)
	newSlice[len(s)] = elment

	return newSlice
}

func copySlice(s []int) []int {
	newSlice := make([]int, len(s))
	copy(newSlice, s)

	return newSlice
}

func removeElement(s []int, ind int) ([]int, error) {
	if ind >= len(s) {
		return nil, fmt.Errorf("index %d out of range %d", ind, len(s))
	}
	newSlice := make([]int, ind, len(s)-1)
	copy(newSlice, s[:ind])
	newSlice = append(newSlice, s[ind+1:]...)
	return newSlice, nil
}

func main() {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Слайс: ")
	fmt.Println(originalSlice)
	fmt.Println()
	fmt.Println("Чётный слайс: ")
	fmt.Println(sliceExample(originalSlice))
	fmt.Println()

	addition := 4
	fmt.Printf("Слайс с добавлением элемента %d: \n", addition)
	fmt.Println(addElements(originalSlice, addition))
	fmt.Println()

	copySlice := copySlice(originalSlice)
	originalSlice[2] = -100

	fmt.Println("Копия слайса:")
	fmt.Println(copySlice)
	fmt.Println()

	anotherSlice, _ := removeElement(copySlice, 0)
	fmt.Println("Копия слайса после удаления 0 элемента:")
	fmt.Println(anotherSlice)
	fmt.Println()
}