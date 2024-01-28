package main

import (
	"log"
	"sync"
)

type FiboService struct {
	mutex sync.RWMutex
	cache map[int]int

	opInProgress  map[int]bool
	opResultsChan map[int][]chan int
}

func NewFiboService() *FiboService {
	return &FiboService{
		cache:         make(map[int]int),
		opInProgress:  make(map[int]bool),
		opResultsChan: make(map[int][]chan int),
	}
}

func (f *FiboService) GetFromCache(key int) (v int, ok bool) {
	f.mutex.RLock()
	v, ok = f.cache[key]
	f.mutex.RUnlock()
	return
}

func (f *FiboService) PutToCache(key, value int) {
	f.mutex.Lock()
	f.cache[key] = value
	f.mutex.Unlock()
}

func (f *FiboService) Fibonacci(n int) int {
	// return easy results
	if n <= 1 {
		return n
	}

	// Check if previously calculated
	if v, ok := f.GetFromCache(n); ok {
		return v
	}

	// Check if already in progress
	f.mutex.Lock()
	if f.opInProgress[n] {
		resultChan := make(chan int)

		f.opResultsChan[n] = append(f.opResultsChan[n], resultChan)
		f.mutex.Unlock()

		log.Println("waiting for result of job: ", n)
		result := <-resultChan
		log.Println("Got result for job: ", n)
		return result
	}

	// calculate new operation
	f.opInProgress[n] = true
	f.mutex.Unlock()

	result := f.Fibonacci(n-1) + f.Fibonacci(n-2)

	f.PutToCache(n, result)

	go f.notifyResultChans(n, result)

	return result
}

func (f *FiboService) notifyResultChans(job, result int) {
	log.Println("Notifying to channels waiting for job: ", job)
	f.mutex.Lock()

	if resultChans := f.opResultsChan[job]; resultChans != nil {
		for _, c := range resultChans {
			c <- result
			close(c)
		}
	}

	delete(f.opInProgress, job)
	delete(f.opResultsChan, job)
	f.mutex.Unlock()

	log.Println("Completed notifying to channels waiting for job: ", job)
}

func main() {
	fiboSvc := NewFiboService()

	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var wg sync.WaitGroup
	wg.Add(len(jobs))

	for _, n := range jobs {
		go func(job int) {
			defer wg.Done()
			result := fiboSvc.Fibonacci(job)

			log.Printf("completed job: %d with result %d", job, result)
		}(n)
	}
	wg.Wait()
}
