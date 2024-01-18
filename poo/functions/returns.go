package main

import "fmt"

func sumBasic(a, b int) int {
	return a + b
}

func sumComplete(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}

	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(value int) (int, int, int) {
	return 2 * value, 3 * value, 4 * value
}

func getValuesRefactory(value int) (double int, triple int, quad int) {
	double = 2 * value
	triple = 3 * value
	quad = 4 * value
	return
}

func main() {
	fmt.Println(sumBasic(1, 2))
	fmt.Println(sumComplete(1, 2, 3, 4, 5, 10))

	printNames("Alice", "Bob", "Charlie", "Dave")

	fmt.Println(getValues(2))

	fmt.Println(getValuesRefactory(3))
}
