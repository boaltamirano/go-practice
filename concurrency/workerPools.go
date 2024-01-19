package main

import (
	"fmt"
	"time"
)

func Worker(id int, jobs <-chan int, results chan<- int) {
	fmt.Println("Trabajando con el worker con id", id, "started")
	for j := range jobs {
		fmt.Printf("Trabajando con el worker con id %v, numero para obtener fibonacci de %v\n", id, j)
		results <- Fibonacci(j)
		fmt.Printf("Resultado fibonacci de %d es: %d\n", j, Fibonacci(j))
		fmt.Printf("Final del trabajo del worker con id %v, numero para obtener fibonacci trabajado es %v\n\n\n", id, j)
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	tasks := []int{7, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42}
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))
	for w := 1; w <= 10; w++ {
		go Worker(w, jobs, results)
	}
	time.Sleep(3 * time.Second)
	for _, task := range tasks {
		println("Numero enviado para realizar la tarea: ", task)
		jobs <- task
		time.Sleep(1 * time.Second)
	}
	close(jobs)
	for a := 1; a <= len(tasks); a++ {
		fmt.Println(<-results)
	}
}
