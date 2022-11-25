package main

import (
	"example.arainay.com/patterns/behavioral/chainOfResponsibility"
	"example.arainay.com/patterns/behavioral/strategy"
	"fmt"
)

func main() {
	runChainOfResponsibilityExample()
}

func runStrategyExample() {
	strategy.Example()
}

func runChainOfResponsibilityExample() {
	chainOfResponsibility.Example()
}

type Number interface {
	int64 | float64
}

func sum[T int64 | float64](numbers []T) T {
	var sum T

	for _, num := range numbers {
		sum += num
	}

	return sum
}

func contains[T comparable](elements []T, searchEl T) bool {
	for _, el := range elements {
		if searchEl == el {
			return true
		}
	}
	return false
}

func show[T any](value T) {
	fmt.Println(value)
}
