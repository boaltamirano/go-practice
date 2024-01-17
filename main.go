package main

import (
	"fmt"
	"strconv"
)

func main() {
	var testInt int
	testInt = 8
	testSecond := 7

	fmt.Println(testInt)
	fmt.Println(testSecond)

	myValue, err := strconv.ParseInt("7", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}
}
