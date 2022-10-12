package main

import (
	"constraints"
	"fmt"
)

type NumberType interface {
	int | float32 | float64 | uint16 | uint64 | uint32 | int32
}

func printAny(thing any) {
	fmt.Println(thing)
}

func addGenerics[T constraints.Ordered](list []T) T {
	var sum T
	for _, item := range list {
		sum += item
	}
	return sum
}

func Intsadd(list []int) int {
	var sum int
	for _, item := range list {
		sum += item
	}
	return sum
}

func Floatsadd(list []float32) float32 {
	var sum float32
	for _, item := range list {
		sum += item
	}
	return sum
}

func main() {

	fmt.Printf("sum of ints: %d\n", Intsadd([]int{2, 3, 4, 5, 6}))
	fmt.Printf("sum of floats: %.2f\n", Floatsadd([]float32{1.2, 2.1, 3.1}))
	fmt.Printf("sum of ints: %d\n", addGenerics([]int{2, 3, 4, 5, 6}))
	fmt.Printf("sum of floats: %.2f\n", addGenerics([]float32{1.2, 2.1, 3.1}))

	printAny("hi")
	printAny(1.2)

}
