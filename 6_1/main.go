package main

import (
	"errors"
	"fmt"
)

// Реализовать тип IntStack , который содержит стэк целых чисел.
// У него должны быть методы Push(i int) и
// Pop() int .

type IntStack struct {
	data []int
}

func (intStack *IntStack) Push(value int) {
	intStack.data = append(intStack.data, value)
}

func (intStack *IntStack) Pop() (int, error) {
	lastIndex := len(intStack.data) - 1

	if lastIndex < 0 {
		return 0, errors.New("Empty")
	}

	lastValue := intStack.data[lastIndex]
	intStack.data = intStack.data[:lastIndex]

	return lastValue, nil
}

func main() {
	intStack := IntStack{}

	for i := 1; i <= 5; i++ {
		intStack.Push(i)
	}

	for i := 0; i < 7; i++ {
		if value, error := intStack.Pop(); error == nil {
			fmt.Println(value)
		} else {
			fmt.Println(error.Error())
		}
	}

	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())

	intStack.Push(2)

	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())
}
