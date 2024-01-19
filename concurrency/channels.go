package main

import "fmt"

func main() {
	c := make(chan int, 3)

	c <- 1
	c <- 4

	fmt.Println(<-c)
	fmt.Println(<-c)
}
